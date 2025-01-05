package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct{
	x int
	y int
}

var level [][]rune
var pos position
var size int

func main(){
	// Setup
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	y := 0
	for scanner.Scan(){
		if len(scanner.Text()) == 0{
			break
		}
		level = append(level, make([]rune, 0))
		for x, char := range scanner.Text(){
			switch char{
			case '@':
				pos = position{x*2,y}
				level[y] = append(level[y], char)
				level[y] = append(level[y], '.')
			case  'O':
				level[y] = append(level[y], '[')
				level[y] = append(level[y], ']')
			case '#':
				level[y] = append(level[y], '#')
				level[y] = append(level[y], '#')
			default:
				level[y] = append(level[y], '.')
				level[y] = append(level[y], '.')
			}
			
		
		}
		y++
	}
	size = y

	// Now to process moves
	for scanner.Scan(){
		for _, char := range(scanner.Text()){
			switch char{
			case '<':
				move(-1,0);
			
			case '^':
				move(0,-1);
			case '>':
				move(1,0);
			case 'v':
				move(0,1);
			}
		}
	}
	fmt.Println("Total: ", sum_lvl())
}

func move(dir_x int, dir_y int){
	// If space is empty, move
	if get(pos.x+dir_x,pos.y+dir_y) == '.'{
		level[pos.y][pos.x] = '.'
		level[pos.y + dir_y][pos.x + dir_x] = '@'
		pos.x = pos.x + dir_x
		pos.y = pos.y + dir_y
		return
	}

	// If space has box, push against it, then try to move
	if get(pos.x+dir_x,pos.y+dir_y) == '[' || get(pos.x+dir_x,pos.y+dir_y) == ']'{
		// If push successful, we can move
		if push_box(pos.x + dir_x, pos.y + dir_y, dir_x, dir_y){
			level[pos.y][pos.x] = '.'
			level[pos.y + dir_y][pos.x + dir_x] = '@'
			pos.x = pos.x + dir_x
			pos.y = pos.y + dir_y
		}
		return
	}
}

func push_box(x int, y int, dir_x int, dir_y int)bool{
	// Attempt to push a box
	// If There is no box, we can push something on this spot
	if get(x,y) == '.'{
		return true
	}
	// We cannot push walls
	if get(x,y) == '#'{
		return false
	}
	// First lets get all our left and right side positions
	var left position
	var right position;

	if get(x,y) == '['{
		left = position{x,y}
		right = position{x+1,y}
	} else if get(x,y) == ']'{
		left = position{x-1,y}
		right = position{x,y}		
	}

	// Different cases for each direction
	// LEFT
	if dir_x == -1 && dir_y == 0{
		// If space is empty, push
		if get(left.x - 1, left.y) == '.'{
			level[left.y][left.x - 1] = '['
			level[left.y][left.x] = ']'
			level[left.y][left.x + 1] = '.'
			return true
		}

		// Try and push the other box
		if get(left.x - 1, left.y) == ']' && push_box(left.x - 1,left.y,-1,0){
			level[left.y][left.x - 1] = '['
			level[left.y][left.x] = ']'
			level[left.y][left.x + 1] = '.'
			return true			
		}
	}

	// RIGHT
	if dir_x == 1 && dir_y == 0{
		if get(right.x + 1, right.y) == '.'{
			level[right.y][right.x + 1] = ']'
			level[right.y][right.x] = '['
			level[right.y][right.x - 1] = '.'
			return true
		}
		if get(right.x + 1, right.y) == '[' && push_box(right.x + 1, right.y, 1, 0){
			level[right.y][right.x + 1] = ']'
			level[right.y][right.x] = '['
			level[right.y][right.x - 1] = '.'
			return true
		}
	}

	copy := deep_level_copy()
	// UP & DOWN
	if dir_y == -1 || dir_y == 1{
		// Can we push into both spaces?
		if push_box(left.x,left.y + dir_y,0,dir_y) && push_box(right.x,right.y + dir_y,0,dir_y){
			level[left.y + dir_y][left.x] = '['
			level[right.y + dir_y][right.x] = ']'
			level[left.y][left.x] = '.'
			level[right.y][right.x] = '.'
			return true		
		}else{
			// We may have pushed 1 box but not the other, revert
			rollback_level(copy)
			return false
		}
	}
	return false
}

func get(x int, y int)rune{
	if x < 0 || y < 0 || x >= size * 2 || y >= size{
		return '#'
	}
	return level[y][x]
}

func print_lvl(){
	for _, row := range level{
		for _, char := range row{
			fmt.Printf("%c",char)
		}
		fmt.Println()
	}
	fmt.Println("")
}

func sum_lvl()(int){
	total := 0
	for y, row := range level{
		for x, char := range row{
			if char == '['{
				total +=x
				total +=(y *100)
			}
		}
	}
	return total
}

func deep_level_copy()[][]rune{
	y := 0
	copy := make([][]rune,0)
	for _,row := range level{
		copy = append(copy, make([]rune, 0))
		for _, char := range row{
			copy[y] = append(copy[y], char)
		}
		y++;
	}
	return copy
}

func rollback_level(copy [][]rune){
	for y := range size{
		for x := range size*2{
			level[y][x] = copy[y][x]
		}
	}
}
