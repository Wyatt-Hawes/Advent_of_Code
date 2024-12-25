package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type position struct{
	x int
	y int
}

var level [][]int = make([][]int, 0)
var trail_heads_reached map[position][]position = make(map[position][]position)

var size int;
var total_trails int = 0;
var distinct_trails int = 0;

func main(){
	name := "./input.txt"

	// Getting input into 2d int slice
	file, _ := os.Open(name)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan(){
		line := scanner.Text()
		level = append(level, make([]int, 0))
		for _, char := range line{
			val,_ := strconv.Atoi(string(char))
			if(char == '.'){
				val = -1
			}
			level[y] = append(level[y], val)
		}
		y++
	}
	size = y
	print_lvl()

	// Find all trail starts
	for y, row := range level{
		for x, val := range row{
			if(val == 0){
				num_valid_trails(position{x,y},x,y,0,[]position{})
				fmt.Println("=")
			}
		}
	}
	fmt.Println("Total Trails: ",total_trails)
	fmt.Println("Distinct Trais: ",distinct_trails)
}

func num_valid_trails(start position,x int, y int, value int, path []position){
	if(value == 9 && get(x,y) == 9){
		// We've reached the end of the trail
		set_add(start, position{x,y}, path)
		return
	}
	if get(x,y) != value{
		return
	}

	// Find all values around with value + 1
	
	// Check UP
	num_valid_trails(start,x,y - 1, value + 1,append(path,position{x,y}))

	// Check RIGHT
	num_valid_trails(start,x + 1,y , value + 1,append(path,position{x,y}))

	// Check DOWN
	num_valid_trails(start,x,y + 1, value + 1,append(path,position{x,y}))

	// Check LEFT
	num_valid_trails(start,x - 1,y , value + 1,append(path,position{x,y}))

	return
}

func set_add(start position, end position, path []position){
	
	distinct_trails ++;
	if(!slices.Contains(trail_heads_reached[start],end)){
		total_trails +=1
		trail_heads_reached[start] = append(trail_heads_reached[start], end)
		fmt.Println("PATH FOUND: ", append(path,position{end.x,end.y}))
	}
}

// Safely get to allow us to index out of bounds
func get(x int, y int)int{
	if x < 0 || y < 0 || x >= size || y >= size{
		return -1
	}
	return level[y][x]
}


func print_lvl(){
	for _, row := range level{
		for _, r := range row{
			if(r == -1){
				fmt.Printf(".")
				continue
			}
			fmt.Printf("%d",r)
		}
		fmt.Println("")
	}
	fmt.Println("============")
}