package main

import (
	"fmt"
	"os"
	"strconv"
)

type gap_info_t struct{
	start int
	size int
	value int
}

func main(){
	part1();
	part2();
}

func part2(){
	file, _ := os.ReadFile("./input.txt")
	disk := string(file)
	compressed := []int{}

	gap_info := []gap_info_t{}
	file_info := []gap_info_t{}

	// Parse incoming data
	index := 0
	for i, char := range disk{
		val,_ := strconv.Atoi(string(char))
		value_to_print := -1
		if i % 2 == 0{
			i := i / 2
			value_to_print = i
			file_info = append(file_info, gap_info_t{index,val, value_to_print})
		}else{
			gap_info = append(gap_info, gap_info_t{index,val, -1})
		}
		for _ = range val{
			index++
			compressed = append(compressed, value_to_print)
		}
	}

	// Loop through all the file blocks backwards to start from highest ID
	for i := len(file_info) - 1; i >= 0; i--{
		current_file := file_info[i]

		// Now find a suitable gap
		for _, gap := range gap_info{
			if current_file.size <= gap.size && gap.start < current_file.start{
				// Found gap
				// Fill in gap
				for i := range(current_file.size){
					compressed[gap.start + i] = current_file.value
				}
				// Replace original spot with spaces
				for i := range(current_file.size){
					compressed[current_file.start + i] = -1
				}
				// Recalculate all the gaps (Brute Force)
				gap_info = recalculate_gaps(compressed)
				break 
			}

		}
	}
	compute_checksum(compressed)
}

func recalculate_gaps(compressed []int)[]gap_info_t{
	gap_info := []gap_info_t{}
	in_range := false
	start := 0

	for i, val := range compressed{
		if val != -1{
			// Is this the end of a gap
			if in_range{
				gap_info = append(gap_info, gap_info_t{start,i - start, -1})
				in_range = false
				continue
			}
			continue
		}
		// Have we started a gap?
		if !in_range{
			start = i
			in_range = true
			continue
		}
	}

	// Check if a gap exists at the end
	if in_range{
		gap_info = append(gap_info, gap_info_t{start,(len(compressed)) - start, -1})
	}

	return gap_info
}

// Print puzzle
func print_comp(compressed []int){
	for _, v := range compressed{
		if(v == -1){
			fmt.Print(".")
			continue
		}
		fmt.Print(v)
	}
	fmt.Println()
}

// Calculate checksum for the puzzle
func compute_checksum(compressed []int){
	checksum := 0
	for i, char := range compressed{
		if char == -1{
			continue
		}
		checksum += (i * char)
	}
	fmt.Println("Checksum: ",checksum)
}


func part1(){
	file, _ := os.ReadFile("./input.txt")
	disk := string(file)
	compressed := []int{}

	// Parse incoming data
	for i, char := range disk{
		val,_ := strconv.Atoi(string(char))
		value_to_print := -1
		if i % 2 == 0{
			id := i / 2
			value_to_print = id
		}

		for _ = range val{
			compressed = append(compressed, value_to_print)
		}
	}

	// 2 pointers to move everything
	left := 0
	right := len(compressed) - 1
	for true {
		if (left >= right){
			break
		}
		if(compressed[left] != -1){
			left++
			continue
		}
		if(compressed[right] == -1){
			right--;
			continue
		}

		// Left is a .  right is a value, we must swap
		compressed[left] = compressed[right]
		compressed[right] = -1
		left++;
		right--;
	}

	compute_checksum(compressed)
}