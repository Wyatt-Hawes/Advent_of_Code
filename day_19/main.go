package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var towels_list []string;

var existing_towels = make(map[string]bool)

var cache = make(map[string]bool)
var ways_cache = make(map[string]int)

var currently_working_on string;

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	// Get list of towels
	towels_list = strings.Split(scanner.Text(),", ")
	for _, val := range towels_list{
		existing_towels[val] = false
	}

	scanner.Scan() // For the empty line

	total :=0
	total_ways := 0
	for scanner.Scan(){
		currently_working_on = scanner.Text()

		possible := is_valid_pattern(scanner.Text())
		ways := count_different_ways(scanner.Text())

		if possible {total++;total_ways += ways}
	}
	fmt.Println("Total Valid:",total, total_ways)
}

func count_different_ways(pattern string)int{
	if len(pattern) == 0{
		return 1
	}

	value, exists := ways_cache[pattern]
	if exists{
		return value
	}
	total := 0
	// Loop through all patterns
	for towel, _ := range existing_towels{
		if towel_fits_pattern(towel, pattern){
			// Recurse
			s := count_different_ways(pattern[len(towel):])
			total += s
		}
	}

	ways_cache[pattern] = total
	return total
}

func is_valid_pattern(pattern string)bool{
	if len(pattern) == 0{
		return true
	}

	value, exists := cache[pattern]
	if exists{
		return value
	}

	// Loop through all patterns
	for towel, _ := range existing_towels{
		if towel_fits_pattern(towel, pattern){
			// Recurse
			s := is_valid_pattern(pattern[len(towel):])
			cache[pattern[len(towel):]] = s
	
			if (s){
				cache[pattern] = true
				return true
			}
		}
	}
	cache[pattern] = false
	return false
}

func towel_fits_pattern(towel string, pattern string) bool{
	for i, _ := range towel{
		if len(towel) > len(pattern){
			return false
		}
		if pattern[i] != towel[i]{
			return false
		}
	}
	return true
}

func tab(num int)string{
	s := ""
	for range num{
		s+= " "
	}
	return s
}