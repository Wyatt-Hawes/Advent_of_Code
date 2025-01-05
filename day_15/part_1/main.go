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
			level[y] = append(level[y], char)
			if char == '@'{
				pos = position{x,y}
			}
		}
		y++
	}
	size = y
	// fmt.Println("Initial State")
	// print_lvl()
	// Now to process moves
	for scanner.Scan(){
		for _, char := range(scanner.Text()){
			switch char{
			case '<':
				//fmt.Println("Move <:")
				move(-1,0);
				//print_lvl()
			
			case '^':
				//fmt.Println("Move ^:")
				move(0,-1);
				//print_lvl()
			case '>':
				//fmt.Println("Move >:")
				move(1,0);
				//print_lvl()
			case 'v':
				//fmt.Println("Move v:")
				move(0,1);
				//print_lvl()
			}
		}
	}
	fmt.Println("Total: ", sum_lvl())
}

func move(dir_x int, dir_y int){
	// If space is empty, move
	//fmt.Printf("Ahead is %c\n",get(pos.x+dir_x,pos.y+dir_y))
	if get(pos.x+dir_x,pos.y+dir_y) == '.'{
		level[pos.y][pos.x] = '.'
		level[pos.y + dir_y][pos.x + dir_x] = '@'
		pos.x = pos.x + dir_x
		pos.y = pos.y + dir_y
		return
	}

	// If space has box, push against it, then try to move
	if get(pos.x+dir_x,pos.y+dir_y) == 'O'{
		// If push successful, we can move
		if push(pos.x + dir_x, pos.y + dir_y, dir_x, dir_y){
			level[pos.y][pos.x] = '.'
			level[pos.y + dir_y][pos.x + dir_x] = '@'
			pos.x = pos.x + dir_x
			pos.y = pos.y + dir_y
		}
		return
	}
}

func push(x int, y int, dir_x int, dir_y int)bool{
	// Attempt to push a box
	// If There is no box, we can push something on this spot
	if get(x,y) == '.'{
		return true
	}
	// We cannot push walls
	if get(x,y) == '#'{
		return false
	}
	// If there is a box in the next space and we can push it, we can push this box
	if get(x+dir_x,y+dir_y) == 'O' && push(x + dir_x, y + dir_y, dir_x, dir_y){
		level[y][x] = '.'
		level[y+dir_y][x+dir_x] = 'O'
		return true
	}
	// if the next space is empty, we can push into it
	if get(x+dir_x,y+dir_y) == '.'{
		level[y][x] = '.'
		level[y+dir_y][x+dir_x] = 'O'
		return true
	}

	return false
}

func get(x int, y int)rune{
	if x < 0 || y < 0 || x >= size || y >= size{
		return '#'
	}
	return level[y][x]
}

func print_lvl(){
	//fmt.Println(pos.x,pos.y)
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
			if char == 'O'{
				total +=x
				total +=(y *100)
			}
		}
	}
	return total
}