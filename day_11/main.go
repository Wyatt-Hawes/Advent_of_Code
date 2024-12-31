package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stone_step struct {
	next_stones []int
}

var curr_step map[int]int = make(map[int]int)
var calculations map[int]stone_step = make(map[int]stone_step)

func main(){
	// Get input
	file, _ := os.ReadFile("./input.txt")
	stones_str := strings.Split(string(file), " ")
	for _, stone := range stones_str{
		v, _ := strconv.Atoi(stone)
		add_stone(curr_step,v,1)
	}

	// Compute the blinks
	for i := range(75){
		blink()

		if(i == 24 || i == 74){
			fmt.Println(i+1,":", count_stones(curr_step))
		}
	}
}

// Calculate all of the same number stones at the same time
func blink(){
	next_step := make(map[int]int)
	for stone, num_stones := range curr_step{
		calcs := get_calculation(stone)
		for _, new_stone := range calcs.next_stones{
			// We add the number of stones that we have to the next step
			add_stone(next_step, new_stone, num_stones)
		}
	}
	curr_step = next_step
}

// Count all the stones
func count_stones(step map[int]int)int{
	total := 0
	for _, num_stones := range step{
		total+= num_stones
	}
	return total
}

// Add a stone to an existing entry, creating the entry if it doesn't exist
func add_stone(step map[int]int,stone int, amount int){
	current_num, exists := step[stone]
	if !exists{
		step[stone] = 0
	}
	step[stone] = current_num + amount
}

// Get the calculation if we have it
func get_calculation(stone int) stone_step{
	calc, exists := calculations[stone]
	if(!exists){
		// Calculate it if we dont have it
		calc = calculate(stone)
		calculations[stone] = calc
	}
	return calc
}

func calculate(stone int)stone_step{
	if(stone == 0){
		return stone_step{[]int{1}}
	}
	if(len(strconv.Itoa(stone)) % 2 == 0){
		stone_string := strconv.Itoa(stone)
		length := len(stone_string)
		stone_1, _ := strconv.Atoi(stone_string[0:length / 2])
		stone_2, _ := strconv.Atoi(stone_string[length / 2:])
		return stone_step{[]int{stone_1,stone_2}}
	}
	return stone_step{[]int{stone * 2024}}
}

// func blink_brute_force(stones []int)[]int{
// 	new_stones := make([]int,0)
// 	for _, stone := range stones{
// 		// If stone engraved w/ 0, replaced with 1
// 		if(stone == 0){
// 			new_stones = append(new_stones, 1)
// 			continue
// 		}
// 		// If even digits, split into 2 stones
// 		if(len(strconv.Itoa(stone)) % 2 == 0){
// 			stone_string := strconv.Itoa(stone)
// 			length := len(stone_string)
// 			stone_1, _ := strconv.Atoi(stone_string[0:length / 2])
// 			stone_2, _ := strconv.Atoi(stone_string[length / 2:])
// 			new_stones = append(new_stones, stone_1)
// 			new_stones = append(new_stones, stone_2)
// 			continue
// 		}
// 		new_stones = append(new_stones, stone * 2024)
// 	}
// 	return new_stones
// }