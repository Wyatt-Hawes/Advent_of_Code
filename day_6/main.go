package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type position struct{
	x int
	y int
}

var level [][]rune = make([][]rune, 0)
var directions_traveled [][][]rune
var curr_pos position = position{0,0}
var original_pos position = position{0,0}
var size int;
var unique_positions = 0

func main(){
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Lets read the input and grab the current position
	y := 0
	for scanner.Scan(){
		line := scanner.Text()
		level = append(level, make([]rune, 0))
		for x, char := range line{
			level[y] = append(level[y], char)
			if char == '^'{
				curr_pos.x = x
				curr_pos.y = y
				original_pos.x = x
				original_pos.y = y
			}
		}
		y++
	}

	// Now that we have the size, lets make our slices
	size = y
	directions_traveled = make([][][]rune, size)
	for i := range(directions_traveled){
		directions_traveled[i] = make([][]rune,size)
	}

	// Make a copy so we can restore later
	duplicate := make_copy()
	
	// Now lets play the game
	for is_valid(curr_pos.x,curr_pos.y){
		play_move()
	}
	final_pos := unique_positions
	
	// Restore
	restore(duplicate)

	// Now lets see how many blockers we can put up to cause infinite loops
	infinite_loops := 0
	for y := range size{
		for x := range size{
			// Skip
			if(get(x,y) != '.'){continue}

			// Set blocker
			set(x,y,'O');

			// Play game
			for is_valid(curr_pos.x,curr_pos.y){
				if(!play_move()){
					infinite_loops++
					break // We are in a infinite loop
				}
			}
			//Restore
			restore(duplicate)
		}
	}


	fmt.Println("Unique Positions: ", final_pos)
	fmt.Println("Blocker Positions: ", infinite_loops)
}

func play_move() bool{ 
	// First get direction
	
	direction := get(curr_pos.x, curr_pos.y)
	front := get_front();

	// We can "Walk" forward
	if front != '#' && front != 'O'{
		if(front == '.'){
			unique_positions++
		}

		// Have we left the square at the current direction?
		exit_dirs := directions_traveled[curr_pos.y][curr_pos.x]
		if slices.Contains(exit_dirs,direction){
			// We are in a loop
			return false;
		}

		// We are exiting current square at the direction
		directions_traveled[curr_pos.y][curr_pos.x] = append(directions_traveled[curr_pos.y][curr_pos.x], direction)
		walk()
		return true;
	}
	// We must turn
	turn()
	return true;
}

func is_valid(x int, y int)bool{
	if x < 0 || x >= size || y < 0 || y >= size{
		return false
	}
	return true
}

func get_front()rune{
	direction := get(curr_pos.x, curr_pos.y)
	var front rune;
	switch direction{
	case '^':
		// UP
		front = get(curr_pos.x, curr_pos.y -1)
	case '>':
		// RIGHT
		front = get(curr_pos.x+1, curr_pos.y)
	case 'v':
		// DOWN
		front = get(curr_pos.x, curr_pos.y + 1)
	case '<':
		// LEFT
		front = get(curr_pos.x-1, curr_pos.y)
	
	default:
		fmt.Println("ERROR: PLAY ERROR")
	}
	return front
}

func turn(){
	direction := get(curr_pos.x, curr_pos.y)
	var char rune;
	switch direction{
	case '^':
		// UP
		char = '>'
	case '>':
		// RIGHT
		char = 'v'
	case 'v':
		// DOWN
		char = '<'
	case '<':
		// LEFT
		char = '^'
	default:
		fmt.Println("ERROR: TURN ERROR")
	}
	set(curr_pos.x,curr_pos.y,char)
}

func walk(){
	direction := get(curr_pos.x, curr_pos.y)
	set(curr_pos.x,curr_pos.y,'X')
	var char rune;

	switch direction{
	case '^':
		// UP
		curr_pos.y-=1
		char = '^'
	case '>':
		// RIGHT
		curr_pos.x +=1
		char = '>'
	case 'v':
		// DOWN
		curr_pos.y +=1
		char = 'v'
	case '<':
		// LEFT
		curr_pos.x -=1
		char = '<'
	
	default:
		fmt.Println("ERROR: WALK ERROR")
	}
	set(curr_pos.x,curr_pos.y,char)
}

func get(x int, y int)rune{
	if(!is_valid(x,y)){
		return '.'
	}
	return level[y][x]
}

func set(x int, y int, value rune){
	if(!is_valid(x,y)){
		return
	}
	level[y][x] = value
}

func print_lvl(){
	for _, row := range level{
		for _, r := range row{
			fmt.Printf("%c",r)
		}
		fmt.Println("")
	}
	fmt.Println("==========",curr_pos, unique_positions)
	fmt.Println("")
}

func make_copy() [][]rune{
	cpy := make([][]rune,size)
	for i := range cpy{
		cpy[i] = make([]rune, size)
		copy(cpy[i],level[i])
	}
	return cpy
}

func restore(cpy [][]rune){
	directions_traveled = make([][][]rune, size)
	for i := range(directions_traveled){
		directions_traveled[i] = make([][]rune,size)
	}
	curr_pos = original_pos
	unique_positions = 0
	for i := range cpy{
		copy(level[i],cpy[i])
	}
}