package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operations struct {
	added      int
	multiplied int
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	solution := 0
	for scanner.Scan(){
		line := scanner.Text()

		split := strings.Split(line, ": ")
		val_str := strings.Split(split[1]," ")
		values := make([]int,0)
	 	total,_ := strconv.Atoi(split[0])

		for _, v_s := range val_str{
			v,_ := strconv.Atoi(v_s)
			values = append(values, v)
		}

		is_possible := calculate(values, total)

		if(is_possible){
			solution += total
		}
	}
	fmt.Println("Result: ", solution)
}

func calculate(values []int, total int) bool {
	
	if len(values) == 1{
		return values[0] == total
	}
	
	add_list := append([]int{values[0] + values[1]},values[2:]...)
	mul_list := append([]int{values[0] * values[1]},values[2:]...)
	combined_val, _ := strconv.Atoi(strconv.Itoa((values[0])) + strconv.Itoa((values[1])))
	con_list := append([]int{combined_val},values[2:]...)

	return calculate(add_list, total) || calculate(mul_list, total) || calculate(con_list, total)
}