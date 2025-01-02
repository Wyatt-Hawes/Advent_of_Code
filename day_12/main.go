package main

import (
	"bufio"
	"fmt"
	"os"
)

type plot struct {
	plant rune
	counted bool
}

type region struct {
	plant rune
	area int
	perimeter int
	sides int
}

var level [][]plot = make([][]plot, 0)
var size int

var regions []region = make([]region, 0)

func main(){
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	y:= 0
	for scanner.Scan(){
		line := scanner.Text()
		level = append(level, make([]plot, 0))
		for _, r := range line{
			level[y] = append(level[y], plot{r,false})
		}
		y++
	}
	size = y
	//print_lvl()

	// Iterate through all cells to find regions
	for y := range size{
		for x := range size{
			new_region := find_region(x,y)
			if new_region.area == 0{
				continue
			}
			regions = append(regions, new_region)
			//print_region(new_region)
			//print_lvl()
		}
	}
	
	//fmt.Println(get_side_contribution(1,0))
	find_price(regions)
}


func find_region(x int, y int) region{
	new_region := region{get(x,y).plant,0,0,0}
	if new_region.plant == -1 || get(x,y).counted{
		return new_region
	}

	// Cell isnt counted, its the first in our region
	add_to_region(&new_region,x,y)
	
	return new_region
}

func add_to_region(region *region, x int, y int){
	plot := get(x,y)
	if plot.counted || region.plant != plot.plant{
		return
	}
	count(x,y)
	region.perimeter += get_perimeter(x,y)
	region.sides += get_side_contribution(x,y)
	region.area +=1

	// Now check all directions
	add_to_region(region, x - 1, y)
	add_to_region(region, x + 1, y)
	add_to_region(region, x, y - 1)
	add_to_region(region, x, y + 1)
}

func get_perimeter(x int, y int) int{
	plot := get(x,y)
	perimeter := 0

	if get(x-1,y).plant != plot.plant{
		perimeter++
	}
	if get(x+1,y).plant != plot.plant{
		perimeter++
	} 
	if get(x,y-1).plant != plot.plant{
		perimeter++
	} 
	if get(x,y+1).plant != plot.plant{
		perimeter++
	} 
	return perimeter
}

func get_side_contribution(x int, y int)int{
	plot := get(x,y)
	sides := 0
	p := plot.plant

	// An L shape of different values Means a corner
	if get(x - 1, y).plant != p && get(x,y-1).plant != p{
		sides++
	}
	if get(x + 1, y).plant != p && get(x,y-1).plant != p{
		sides++
	}
	if get(x + 1, y).plant != p && get(x,y+1).plant != p{
		sides++
	}
	if get(x - 1, y).plant != p && get(x,y+1).plant != p{
		sides++
	}

	// A "square" w/ 3 same values and 1 non-value means a corner
	if get(x - 1, y).plant == p && get(x,y-1).plant == p && get(x - 1, y-1).plant != p{
		sides++
	}
	if get(x + 1, y).plant == p && get(x,y-1).plant == p && get(x + 1, y-1).plant != p{
		sides++
	}
	if get(x + 1, y).plant == p && get(x,y+1).plant == p && get(x + 1, y+1).plant != p{
		sides++
	}
	if get(x - 1, y).plant == p && get(x,y+1).plant == p && get(x - 1, y+1).plant != p{
		sides++
	}
	return sides
}

func count(x int, y int){
	level[y][x].counted = true
}

// Safely get to allow us to index out of bounds
func get(x int, y int)plot{
	if x < 0 || y < 0 || x >= size || y >= size{
		return plot{-1,true}
	}
	return level[y][x]
}

func print_lvl(){
	for y := range size{
		row := level[y]
		for _, p := range row{
			fmt.Printf("%c",p.plant)
		}
		fmt.Print(" | ")
		for _, p := range row{
			if p.counted{
				fmt.Print(".")
			}else{
				fmt.Print("X")
			}
		}
		fmt.Println()
	}
}

func find_price(region_list []region){
	total_area := 0
	total_perimeter := 0
	total_price := 0
	discount := 0
	for _, region := range region_list{
		total_area += region.area
		total_perimeter += region.perimeter
		total_price += region.area * region.perimeter
		discount += region.area * region.sides
	}

	fmt.Printf("Total Area: %d\nTotal Perimeter: %d\nTotal Price: %d\nDiscounted Price: %d\n",total_area,total_perimeter,total_price,discount)
}

func print_region(r region){
	fmt.Printf("%c: \nPerimeter: %d\nArea: %d\nSides: %d\n===========\n", r.plant, r.perimeter, r.area,r.sides)
}