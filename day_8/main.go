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

var level [][]rune = make([][]rune, 0)
var seen map[rune][]position = make(map[rune][]position)
var size int;
var unique_locations = 0

var placed_locations map[position]bool = make(map[position]bool)

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan(){
		line := scanner.Text()
		level = append(level, make([]rune, 0))
		for x, char := range line{
			level[y] = append(level[y], char)
			if(char != '.'){
				seen[char] = append(seen[char], position{x,y})
			}
		}
		y++
	}
	size = y
	print_lvl()
	//fmt.Println(seen)

	for _,values := range seen{
		for _, value := range values{
			others := removeElement(values,value)
			
			for _, other := range others{
				distance := get_distance(value, other)
				curr_pos := position{value.x - distance.x, value.y - distance.y}
				place_interference(curr_pos.x, curr_pos.y)
				
				// FOR PART 2
				place_interference(curr_pos.x + distance.x, curr_pos.y + distance.y)
				for ok := true; ok; ok = is_valid(curr_pos.x,curr_pos.y){
					place_interference(curr_pos.x, curr_pos.y)
					curr_pos = position{curr_pos.x - distance.x, curr_pos.y - distance.y}
				}
				// END OF PART 2
			}
		}
	}

	print_lvl()
	//fmt.Println(placed_locations)
	fmt.Println("Unique Locations: ", unique_locations)
}

func is_valid(x int, y int)bool{
	if x < 0 || y < 0 || x >= size || y >= size{
		return false
	}
	return true
}

func place_interference(x int, y int){
	if !is_valid(x,y){
		return
	}
	
	_, exists := placed_locations[position{x,y}]
	if(!exists){
		unique_locations++
	}
	placed_locations[position{x,y}] = true
	if level[y][x] == '.'{
		level[y][x] = '#'
	}
}

// MISUNDERSTOOD THE PROBLEM
// func has_two_in_line(x int, y int)bool{
// 	duplicates := make(map[rune]int)
// 	// Check all x values
// 	for x1 := range size{
// 		if level[y][x1] != '.' && level[y][x1] != '#'{
// 			_, exists := duplicates[level[y][x1]]
// 			if !exists{
// 				duplicates[level[y][x1]] = 0
// 			}
// 			duplicates[level[y][x1]]++
// 		}
// 	}
// 	// Check all y values
// 	for y1 := range size{
// 		if level[y1][x] != '.' && level[y1][x] != '#'{
// 			_, exists := duplicates[level[y1][x]]
// 			if (!exists){
// 				duplicates[level[y1][x]] = 0
// 			}
// 			duplicates[level[y1][x]]++
// 		}
// 	}
// 	for _, value := range duplicates{
// 		if value >=2{
// 			return true
// 		}
// 	}
// 	return false
// }

func print_lvl(){
	for _, row := range level{
		for _, r := range row{
			fmt.Printf("%c",r)
		}
		fmt.Println("")
	}
	fmt.Println("============")
}

func get_distance(p1 position, p2 position)position{
	return position{p2.x - p1.x, p2.y - p1.y}
}


func removeElement(slice []position, value position) []position {
    newSlice := []position{}
    for _, v := range slice {
        if v.x != value.x && v.y != value.y {
            newSlice = append(newSlice, v)
        }
    }
    return newSlice
}