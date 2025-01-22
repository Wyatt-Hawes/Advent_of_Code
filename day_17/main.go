package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var A = 0
var B = 0
var C = 0

var program []int = make([]int, 0)
var ip = 0
var output []int = make([]int, 0)

var value = 01000000000000000
var solution = -1;

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	fmt.Sscanf(scanner.Text(),"Register A: %d",&A)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(),"Register B: %d",&B)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(),"Register C: %d",&C)

	// This is a newline
	scanner.Scan()
	scanner.Scan()

	temp_str := ""
	fmt.Sscanf(scanner.Text(),"Program: %s",&temp_str)
	temp_arr := strings.Split(temp_str, ",")
	for _, val_str := range temp_arr{
		num, _ := strconv.Atoi(val_str)
		program = append(program, num)
	}

	complete_program()
	fmt.Println("Output:",strings.Trim(strings.Join(strings.Fields(fmt.Sprint(output)), ","), "[]"))

	// Now for part 2
	find_digit(15)

	fmt.Println("SMALLEST SOLUTION: ",solution)
}

func find_digit(digit int){
	if digit == -1{
		if solution == -1{
			solution = value;
		}
		solution = int(math.Min(float64(solution),float64(value)))
		return
	}
	for i := 0; i < 8; i++{
		reset(value)
		complete_program()
		if is_still_valid(digit){
			// Go deeper
			find_digit(digit -1)
		}
		// Add 1 to this digit and try again
		value += int(math.Pow(8,float64(digit)))
	}
	// After attempting all digits, undo changes
	value -= int(math.Pow(8,float64(digit))) * 8
}

func is_still_valid(digit int)bool{
	for i := digit; i < 16; i++{
		if program[i] != output[i]{
			return false
		}
	}
	return true
}

func complete_program(){
	for ; ip < len(program);{
		do_instruction(program[ip],program[ip+1])
		ip+=2
	}
}

func reset(new_A int){
	A = new_A
	B = 0
	C = 0
	ip = 0
	output = make([]int, 0)
}

func do_instruction(instruction int, combo int)bool{
	// First ensure combo is the real value we want
	switch instruction{
		// ADV
		case 0:
			combo = interpret_combo(combo)
			A = A / int((math.Pow(2,float64(combo))))
			return false
		case 1: // BXL
			B = B^combo
			return false
		case 2: // BST
			combo = interpret_combo(combo)
			B = combo % 8
			return false
		case 3: // JNZ
			if A != 0{
				ip = combo - 2
			}
			return false
		case 4: // BXC
			B = B ^ C
			return false
		case 5: // OUT
			combo = interpret_combo(combo) % 8
			output = append(output, combo)
			return false
		case 6: // BDV
			combo = interpret_combo(combo)
			B = A / int((math.Pow(2,float64(combo))))
			return false
		case 7: // CDV
			combo = interpret_combo(combo)
			C = A / int((math.Pow(2,float64(combo))))
			return false

	}
	return false
}


func interpret_combo(combo int)(int){
	switch combo{
		case 0,1,2,3:
			return combo
		case 4:
			return A
		case 5:
			return B
		case 6: 
			return C
	}
	fmt.Println("COMBO ERROR: ",combo)
	return combo
}

func print_state(){
	fmt.Println(A,B,C)
	fmt.Println(program)
	
	fmt.Print(" ")
	for range ip{
		fmt.Print("  ")
	}
	fmt.Print("^")
	fmt.Println("\n=============")
}