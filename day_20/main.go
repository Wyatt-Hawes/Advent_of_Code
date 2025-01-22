package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type position struct{
	x int
	y int
}
var start position;
var current_pos position;
var goal position;

var distances [][]int
var level [][]rune = make([][]rune, 0)
var available_cheats = make(map[int]int)

var level_size = 0

var CHEAT_DISTANCE = 2; // FOR PART 1
// var CHEAT_DISTANCE = 20; // FOR PART 2


func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan(){
		line := scanner.Text()
		level = append(level, make([]rune, 0))
		for x, char := range line{
			if char == 'S'{
				char = '.'
				start = position{x,y}
			}else if char == 'E'{
				char = '.'
				goal = position{x,y}
			}

			level[y] = append(level[y], char)
		}
		y++
	}
	level_size = y
	
	distances = make_empty_distances()
	get_points_to_goal(start.x,start.y,0)

	// Time to find cheats
	for y := range level_size{
		for x := range level_size{
			get_available_cheats_at(x,y,CHEAT_DISTANCE)
		}
	}

	// Now count all the cheats
	total := 0
	for cheat_dist, num_cheats := range available_cheats{
		if cheat_dist >=100{
			total += num_cheats
		}
	}
	fmt.Println(total,"cheats save at least 100 picoseconds")
}

func get_available_cheats_at(x int, y int, noclip int){
	// Get our start 
	start_val := get_dist(x,y)

	// We are a wall/unreachable, return
	if start_val == -1{
		return 
	}
	// Get the cheats in all directions

	// Loop through all potential cheat squares, has some extra that we need to pass over
	for j := y - noclip; j <= y + noclip; j++{
		for i := x - noclip; i <= x + noclip; i++{
			// Check if this square is valid
			if j < 0 || i < 0 || j >= level_size || i >= level_size{continue} // We are out of bounds
			are_within, man_dist := within_manhatten(x,y,i,j,noclip)
			if !are_within{continue} // Too far from starting position
			if distances[y][x] == -1{continue} // Not a valid stopping location

			// How much distance have we saved?
			cheat_dist := start_val + man_dist;
			orig_dist := distances[j][i]

			// Add cheat if we have saved distance
			if cheat_dist < orig_dist{
				add_cheat(orig_dist - cheat_dist)
			}
		}
	} 
}

// Are we within a given manhatten distance, returns TRUE/FALSE, manhatten_distance
func within_manhatten(start_x int, start_y int, end_x int, end_y int, max_distance int)(bool, int){
	x_dist := int(math.Abs(float64(start_x) - float64(end_x)))
	y_dist := int(math.Abs(float64(start_y)-float64(end_y)))
	return (x_dist + y_dist) <= max_distance, (x_dist + y_dist)
}

func add_cheat(time_save int){
	val, exists := available_cheats[time_save]
	if !exists{
		available_cheats[time_save] = 0
	}
	available_cheats[time_save] = val + 1
}

func get_points_to_goal(x int, y int, steps int){
	// Out of bounds
	if x < 0 || x >= level_size || y < 0 || y >= level_size{
		return
	}
	// Wall, so invalid
	if get(x,y) == '#'{
		return
	}
	// Shorter distance already exists
	if distances[y][x] != -1 && distances[y][x] <= steps{
		return
	}
	// Mark that we currently are the shortest distance
	if distances[y][x] == -1 || distances[y][x] > steps{
		distances[y][x] = steps
	}
	// Dont keep searching if we reached the goal
	if x == goal.x && y == goal.y{
		return
	}
	
	// Search in all 4 directions
	// Right
	get_points_to_goal(x + 1, y, steps + 1)
	// Down
	get_points_to_goal(x , y + 1, steps + 1)
	// Up
	get_points_to_goal(x , y - 1, steps + 1)
	// Left
	get_points_to_goal(x - 1, y, steps + 1)
}

// Safely get a part of the level
func get(x int, y int)rune{
	if x < 0 || y < 0 || x >= level_size || y >= level_size{
		return '#'
	}
	return level[y][x]
}

// Safely get a caclulated distance
func get_dist(x int, y int)int{
	if x < 0 || y < 0 || x >= level_size || y >= level_size{
		return -1
	}
	return distances[y][x]
}

// Reset distance array
func make_empty_distances()[][]int{
	arr := make([][]int, level_size)

	for i := range level_size {
		arr[i] = make([]int, level_size)
		for j := range level_size {
			arr[i][j] = -1
		}
	}
	return arr
}

// Print full level
func print_lvl() {
	for _, row := range level {
		for _, char := range row {
			fmt.Printf("%c", char)
		}
		fmt.Println()
	}
	fmt.Println(start, goal)
}

// Print calculated distances
func print_distances(){
	for _, row := range distances {
		for _, char := range row {
			if char == -1{
				fmt.Print("   ")
				continue
			}
			fmt.Printf("%02d ", char)
		}
		fmt.Println()
	}
	fmt.Println("")
}