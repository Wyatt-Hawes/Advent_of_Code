package main

import (
	"bufio"
	"fmt"
	"os"
)

var size = 71
var scan_bytes = 1024 // Part 1

var level [][]rune = make([][]rune, size)

var distances [][]int

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	for i := range size {
		level[i] = make([]rune, size)
		for j := range size {
			level[i][j] = '.'
		}
	}

	

	// Now process input
	num_bytes := 0

	// Part 1
	for scanner.Scan(){
		
		if num_bytes >= scan_bytes{
			break
		}
		num_bytes++
		x, y := 0, 0
		fmt.Sscanf(scanner.Text(),"%d,%d",&x,&y)
		level[y][x] = '#'
	}

	make_empty_distances()
	get_points_to_goal(0,0,0)
	fmt.Println("Shortest Distance: ", distances[size-1][size-1])
	
	// Part 2
	
	// Applying the next blocker then computing distance
	// It would probably be much faster to work backwards and find the first NON blocked one
	// But this is faster to code since we process in order of the input
	for scanner.Scan(){
		num_bytes++
		x, y := 0, 0
		fmt.Sscanf(scanner.Text(),"%d,%d",&x,&y)
		level[y][x] = '#'

		make_empty_distances()
		get_points_to_goal(0,0,0)
		if num_bytes % 200 == 0{
			fmt.Println("At:",num_bytes)
		}
		if distances[size-1][size-1] == -1{
			fmt.Print("Exit Blocked! Byte:",num_bytes," ")
			fmt.Printf("%d,%d\n",x,y)
			break
		}
	}
}

func get_points_to_goal(x int, y int, steps int){
	if x < 0 || x >= size || y < 0 || y >= size{
		return
	}
	if get(x,y) == '#'{
		return
	}

	if distances[y][x] != -1 && distances[y][x] <= steps{
		return
	}

	if distances[y][x] == -1 || distances[y][x] > steps{
		distances[y][x] = steps
	}

	if x == size - 1 && y == size - 1{
		return
	}
	
	// Right
	get_points_to_goal(x + 1, y, steps + 1)
	
	// Down
	get_points_to_goal(x , y + 1, steps + 1)
	
	
	
	// Up
	get_points_to_goal(x , y - 1, steps + 1)
	// Left
	get_points_to_goal(x - 1, y, steps + 1)
}

func get(x int, y int)rune{
	if x < 0 || y < 0 || x >= size || y >= size{
		return '#'
	}
	return level[y][x]
}

func make_empty_distances(){
	distances = make([][]int, size)

	for i := range size {
		distances[i] = make([]int, size)
		for j := range size {
			distances[i][j] = -1
		}
	}
}

func print_lvl() {
	for _, row := range level {
		for _, char := range row {
			fmt.Printf("%c", char)
		}
		fmt.Println()
	}
	fmt.Println("")
}

func print_distances(){
	for _, row := range distances {
		for _, char := range row {
			if char == -1{
				fmt.Print("     ")
				continue
			}
			fmt.Printf("%04d ", char)
		}
		fmt.Println()
	}
	fmt.Println("")
}