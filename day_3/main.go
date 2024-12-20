package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var r1 = []int{}
var r2 = []int{}

var safe_reports_1 = 0
var safe_reports_2 = 0

func main(){
	// Purposefully throw out error values
	byte_slice, _ := os.ReadFile("./input.txt")

	input_string := string(byte_slice)

	// Match the following:
	// mul(
	// 1+ of digits 0-9
	// a literal ,
	// 1+ of digits 0-9
	// )
	// Stores the 2 numbers in match[1] and match[2]
	regex_part_1, _ := regexp.Compile("mul[(]([0-9]+),([0-9]+)[)]")

	// Same as above but OR "do()" and "don't()" literal matches
	regex_part_2, _ := regexp.Compile("mul[(]([0-9]+),([0-9]+)[)]|do[(][)]|don't[(][)]")

	total_part_1 := 0
	for _, match := range regex_part_1.FindAllStringSubmatch(input_string,-1){
		v1, _ := strconv.Atoi(match[1])
		v2, _ := strconv.Atoi(match[2])
		total_part_1 += (v1 * v2)
	}

	total_part_2 := 0
	enabled := true
	for _, match := range regex_part_2.FindAllStringSubmatch(input_string, -1){
		if match[0] == "do()"{
			enabled = true;
			continue
		}
		if(match[0] == "don't()"){
			enabled = false;
			continue
		}

		if (enabled){
			v1, _ := strconv.Atoi(match[1])
			v2, _ := strconv.Atoi(match[2])
			total_part_2 += (v1 * v2)
		}
	}

	fmt.Println("Total: ", total_part_1)
	fmt.Println("Total w/ do() & don't(): ", total_part_2)
}