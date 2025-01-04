package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type moves struct{
	A int
	B int
}

var a_x, a_y, b_x, b_y int
var prize_x, prize_y int
var scanner *bufio.Scanner
var wait = bufio.NewScanner(os.Stdin)

var rx = regexp.MustCompile("X.([0-9]+)|Y.([0-9]+)")

func main() {
	file, _ := os.Open("./input.txt")
	scanner = bufio.NewScanner(file)
	total := 0
	part2_total := 0
	for load_next_game(){
		a,b := play_game()
		//fmt.Println(a,b)
		if a != -1 && b != -1{
			//fmt.Println(a,b)
			total += (a*3) + (b)
		}

		a,b = play_game_2()
		fmt.Println(a,b)
		if a != -1 && b != -1{
			//fmt.Println(a,b)
			part2_total += (a*3) + (b)
		}
	}
	fmt.Println("Total: ", total)
	fmt.Println("Total 2: ", part2_total)
}	

func play_game()(int,int){

	num_a_presses := -1
	
	// End of part 2
	for true{
		num_a_presses++
		targ_x, targ_y := a_x * num_a_presses, a_y * num_a_presses
		if(targ_x > prize_x || targ_y > prize_y){
			// Not possible
			//fmt.Println("Not possible, reached ", num_a_presses)
			return -1, -1
		}
		// Can we now reach the prize with B presses?
		dif_x, dif_y := prize_x - targ_x, prize_y - targ_y
		if dif_x % b_x == 0 && dif_y % b_y == 0{
			// We know its possible to reach
			num_b_presses := dif_x / b_x

			if(num_b_presses > 100 || (dif_x / b_x != dif_y / b_y)){
				// Try to find another solution with less presses
				continue
			}
			return num_a_presses, num_b_presses
		}
		if(num_a_presses > 100){
			//fmt.Println("Not possible, reached1 ", num_a_presses)
			return -1, -1
		}
	}
	return -1, -1
}

func play_game_2()(int,int){
	prize_x+=10000000000000
	prize_y+=10000000000000
	
	v1 := (b_x * prize_y) - (b_y * prize_x)
	v2 := (b_x * a_y) - (b_y * a_x)

	if v1 % v2 != 0{
		return -1,-1
	}

	a_presses := v1 / v2
	v3 := prize_x - (a_x * a_presses)

	if v3 % b_x != 0{
		return -1, -1
	}
	
	b_presses := v3 / b_x
	return a_presses, b_presses
}

func load_next_game()bool{
	var _ error; 
	
	// Find A
	more := scanner.Scan()
	if!more{
		return false
	}
	line := scanner.Text()
	match := rx.FindAllStringSubmatch(line,-1)
	a_x,_ = strconv.Atoi(match[0][1])
	a_y,_ = strconv.Atoi(match[1][2])

	// Find B
	scanner.Scan()
	line = scanner.Text()
	match = rx.FindAllStringSubmatch(line,-1)
	b_x,_ = strconv.Atoi(match[0][1])
	b_y,_ = strconv.Atoi(match[1][2])

	// Find Prize
	scanner.Scan()
	line = scanner.Text()
	match = rx.FindAllStringSubmatch(line,-1)
	prize_x,_ = strconv.Atoi(match[0][1])
	prize_y,_ = strconv.Atoi(match[1][2])

	// Scan next empty line
	scanner.Scan()
	//print_game()
	return true
}

func print_game(){
	fmt.Println("==============")
	fmt.Println("A: ", a_x, a_y)
	fmt.Println("B: ", b_x, b_y)
	fmt.Println("P: ", prize_x, prize_y)
	fmt.Println("==============")
}