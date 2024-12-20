package main

import (
	"bufio"
	"fmt"
	"os"
)
var ceres = []string{}
var size = 0

var directions = [][2]int{
	{-1,-1},{0,-1},{1,-1},
	{1,0},{-1,0},
	{-1,1},{0,1},{1,1},
}

func main(){
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file);

	// ceres := []string{}

	for scanner.Scan(){
		ceres = append(ceres, scanner.Text())
	}
	size = len(ceres);

	total_matches := 0
	for x := 0; x < size; x++{
		for y :=0; y < size; y++{
			// Loop over all directions TL, T, TR, L, R, BL, B, BR
			for _, dir := range directions{
				if check_direction(x,y,dir[0],dir[1], 0){
					total_matches++;
				}
			}
		}
	}

	total_x_matches := 0
	for x := 0; x < size; x++{
		for y :=0; y < size; y++{
			if(check_x_mas(x,y)){
				total_x_matches++;
			}
  		}
	}

	
	fmt.Println("Total XMAS: ", total_matches)
	fmt.Println("Total X-MAS: ", total_x_matches)
}

func check_direction(x int, y int, dirX int, dirY int, stage int) bool{
	const xmas = "XMAS"
	if x < 0 || x >= size || y < 0 || y >= size{
		return false
	}

	// Is this the correct letter?
	if get_c(x,y) == string(xmas[stage]){
		// Did we just check the last value?
		if stage >= 3{
			return true
		}
		// If not, check the next letter in that direction, increment X/Y by direction & increase stage
		stage++;
    	return check_direction(x + dirX, y + dirY, dirX, dirY,stage);
	}
	return false
}

func get_c(x int, y int) string{
	return string(ceres[y][x])
}

func compare_cell_safe(x int, y int, CHAR string) bool{
	if x < 0 || x >= size || y < 0 || y >= size {
    	return false;
  	}
	return get_c(x,y) == CHAR
}

func check_x_mas(x int, y int) bool{
	// We must check for diagonal MAS or SAM for the \ direction (either works) AND the / Direction

	// Lets check from the center, center must be A
	if !compare_cell_safe(x,y,"A"){
		return false
	}

	if is_TL_MAS(x,y) && is_TR_MAS(x,y){
		return true
	}


	return false
}

func is_TL_MAS(x int, y int) bool{
  return (
    (compare_cell_safe(x - 1, y - 1, "S") &&
      compare_cell_safe(x + 1, y + 1, "M")) ||
    (compare_cell_safe(x - 1, y - 1, "M") &&
      compare_cell_safe(x + 1, y + 1, "S")))
}
func is_TR_MAS(x int, y int) bool{
  return (
    (compare_cell_safe(x + 1, y - 1, "S") &&
      compare_cell_safe(x - 1, y + 1, "M")) ||
    (compare_cell_safe(x + 1, y - 1, "M") &&
      compare_cell_safe(x - 1, y + 1, "S")))
}
