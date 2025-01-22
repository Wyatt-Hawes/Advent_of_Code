package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var level [][]rune
type position struct{
	x int
	y int
	direction int
}

type cord struct{
	x int
	y int
}

var start_pos position

var UP = 0
var RIGHT = 1
var DOWN = 2
var LEFT = 3

var goal position
var size_x int
var size_y int

var total_lowest = 9999999;

var valid_path_list [][]cord = make([][]cord, 0)

// var shortest_path = -1;
var seen = make(map[position]int)

var input = bufio.NewScanner(os.Stdin)

func main() {
	file, _ := os.Open("./test2.txt")
	scanner := bufio.NewScanner(file)

	y := 0
	for scanner.Scan(){
		if len(scanner.Text()) == 0{
			break
		}
		level = append(level, make([]rune, 0))
		for x, char := range scanner.Text(){
			level[y] = append(level[y], char)
			if char == 'S'{
				start_pos = position{x,y,1}
			}
			if char == 'E'{
				goal = position{x,y,1}
				seen[goal] = 99999999999999
			}
			if y == 0{
				size_x ++
			}
		}
		y++
		
	}
	traversed := make([]cord,0)
	size_y = y
	get_points_to_goal(start_pos, 0,traversed)
	fmt.Println("Points: ",total_lowest)
	print_paths()
}

func get_points_to_goal(pos position, curr_points int, traversed []cord){
	// END CONDITION at goal
	val, exists := seen[pos]

	if exists && val < curr_points{
		return
	}

	if !exists || val >= curr_points{
		seen[pos] = curr_points
		if pos.x == goal.x && pos.y == goal.y{
			if(total_lowest > curr_points){
				valid_path_list = make([][]cord, 0);
				total_lowest = curr_points;
			}
			if curr_points == total_lowest{
				valid_path_list = append(valid_path_list, traversed)
			}
		}
	}

	// TRY UP, dont backtrack && make sure space is available
	if get(pos.x, pos.y - 1) != '#'{
		get_points_to_goal(position{pos.x,pos.y - 1, UP},curr_points + 1 + (num_rotations(pos.direction, UP) * 1000), append([]cord{{pos.x,pos.y}} , traversed...))
	}

	// TRY RIGHT
	if get(pos.x + 1, pos.y) != '#'{
		get_points_to_goal(position{pos.x + 1,pos.y, RIGHT},curr_points + 1 + (num_rotations(pos.direction, RIGHT) * 1000), append([]cord{{pos.x,pos.y}} , traversed...))
	}

	// TRY DOWN
	if get(pos.x , pos.y + 1) != '#'{
		get_points_to_goal(position{pos.x,pos.y + 1, DOWN},curr_points + 1 + (num_rotations(pos.direction, DOWN) * 1000), append([]cord{{pos.x,pos.y}} , traversed...))
	}

	// TRY LEFT
	if get(pos.x - 1, pos.y) != '#'{
		get_points_to_goal(position{pos.x - 1,pos.y, LEFT},curr_points + 1 + (num_rotations(pos.direction, LEFT) * 1000), append([]cord{{pos.x,pos.y}} , traversed...))
	}
}

func get(x int, y int)rune{
	if x < 0 || y < 0 || x >= size_x || y >= size_y{
		return '#'
	}
	return level[y][x]
}

func print_lvl(pos position){
	for y, row := range level{
		for x, char := range row{
			if x == pos.x && y == pos.y{
				fmt.Print("S")
				continue
			}
			if char == 'S'{
				fmt.Print(".")
				continue
			}
			fmt.Printf("%c",char)
		}
		fmt.Println()
	}
	fmt.Println("P:",start_pos,"G:",goal,"S:",size_x,size_y)
	fmt.Println("")
}

func abs(num int)int{
	return int(math.Abs(float64(num)))
}

func num_rotations(direction1 int, direction2 int)int{
	rotations := abs(direction1 - direction2)
	if rotations == 3{
		return 1
	}
	return rotations
}

func copy_map(orig map[cord]bool)map[cord]bool{
	new_map := make(map[cord]bool)

	for key, _ := range orig{
		new_map[key] = true
	}

	return new_map;
}

func restore_map(r *map[cord]int,data map[cord]int){
	*r = make(map[cord]int)
	for key, val := range data{
		(*r)[key] = val
	}
}

func print_paths(){
	len := 1
	for _, path := range valid_path_list{
		for _, cord := range path{
			if level[cord.y][cord.x] != 'O'{
				len++
			}
			level[cord.y][cord.x] = 'O'
			
		}
	}
	fmt.Println("Tiles: ",len)
}
