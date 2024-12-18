package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var r1 = []int{}
var r2 = []int{}

var safe_reports_1 = 0
var safe_reports_2 = 0

func main(){
	// Purposefully throw out error values

	file, _ := os.Open("./input.txt");
	defer file.Close()

	scanner := bufio.NewScanner(file);

	// Get 2 columns in arrays
	//line_iterate:
	for scanner.Scan(){
		report_str := strings.Split(scanner.Text()," ");
		
		// Convert to int
		report := []int{}
		for i:=0; i < len(report_str); i++{
			val, _ := strconv.Atoi(report_str[i])
			report = append(report, val)
		}

		// Is the report safe?
		if is_safe(report){
			safe_reports_1++
			safe_reports_2++
			continue;
		}

		// Make a copy of the slice
		original_slice := make([]int,len(report));
		copy(original_slice,report)

		// Remove an index and still safe?
		for i:=0; i < len(report_str); i++{
			temp := append(report[:i],report[i+1:]...)

			// Append may modify the original report so we need to restore it
			if is_safe(temp){
				safe_reports_2++
				copy(report, original_slice)
				break
			}
			copy(report, original_slice)
		}
	}
	fmt.Println("Safe Reports: ", safe_reports_1)
	fmt.Println("Safe Reports w/ Removed: ", safe_reports_2)
}

func is_safe(report []int) bool{
	is_decreasing := false

	if report[1] < report[0]{
		is_decreasing = true
	} else if report[1] > report[0]{
		is_decreasing = false
	} else{
		return false;
	}

	for i := 0; i < len(report) - 1; i++{
		var diff int;

		if is_decreasing{
			diff = report[i] - report[i+1]
		}else{
			diff = report[i + 1] - report[i]
		}

		if (diff >= 1 && diff <= 3) {
      		continue;
    	}
		return false;
	}
	return true
}