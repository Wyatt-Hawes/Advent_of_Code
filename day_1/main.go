package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var r1 = []int{}
var r2 = []int{}



func main(){
	// Purposefully throw out error values

	file, _ := os.Open("./input.txt");
	defer file.Close()

	scanner := bufio.NewScanner(file);

	// Get 2 columns in arrays
	for scanner.Scan(){
		slice := strings.Split(scanner.Text(),"   ");

		v1, _ := strconv.Atoi(slice[0])
		v2, _ := strconv.Atoi(slice[1])

		r1 = append(r1, v1)
		r2 = append(r2, v2)
	}

	// sort them
	sort.Ints(r1)
	sort.Ints(r2)

	total_distance := 0
	for i := 0; i < len(r1); i++{
		// Golang apparently doesnt have Int ABS? Convert to Float64 then bakc to Int
		total_distance += int(math.Abs(float64(r1[i] - r2[i])))
	}
	fmt.Println("Total Distance: ",total_distance)

	// Now for part 2
	// Count occurrences of values in R2
	r2_count := map[int]int{}
	for i := 0; i < len(r2); i++{

		// If value doesnt exist, initialize to 0
		_, exists := r2_count[r2[i]]
		if !exists{
			r2_count[r2[i]] = 0;
		}

		r2_count[r2[i]]++;
	}

	similarity_score := 0;
	for i := 0; i < len(r1); i++{
		occurrences, exists := r2_count[r1[i]]
		if(!exists){
			occurrences = 0;
		} 
		similarity_score += (r1[i] * occurrences)
	}
	fmt.Println("Simularity Score: ",similarity_score)
}