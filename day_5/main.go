package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rules map[int][]int = make(map[int][]int)
var updates [][]int = make([][]int,0)

var valid_orderings [][]int = make([][]int,0)
var orderings_to_fix [][]int = make([][]int, 0)

func main(){
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// First scan all the rules
	for scanner.Scan(){
		if(scanner.Text() == ""){
			break
		}
		rule := strings.Split(scanner.Text(), "|")
		int0, _ := strconv.Atoi(rule[0])
		int1, _ := strconv.Atoi(rule[1])

		rules[int1] = append(rules[int1], int0)
	}

	// Then scan all the updates
	i := 0
	for scanner.Scan(){
		update := strings.Split(scanner.Text(),",")
		updates = append(updates, make([]int,0))
		for _, v_s := range update{
			v, _ := strconv.Atoi(v_s)
			updates[i] = append(updates[i], v)
		}
		i++
	}

	// Lets now check all of the orderings
	
	for _, update := range updates{
		is_valid := true;
		seen := calculate_seen(update)

		update_loop:
		for _, val := range update{
			number_rules := rules[val]
			for _, rule := range number_rules{
				// Is the index of a dependency greater than our current value?
				if seen[rule] > seen[val]{
					is_valid = false
					break update_loop
				}
			}
		}
		if is_valid{
			valid_orderings = append(valid_orderings, update)
			continue
		}
		orderings_to_fix = append(orderings_to_fix, update)
	}

	// Now lets fix the invalid orderings

	for x, update := range orderings_to_fix{
		seen := calculate_seen(update)
		for j := 0; j < len(update); j++{
			num := update[j]
			number_rules := rules[num]
			for _, rule := range number_rules{
				if seen[rule] > seen[num]{
					// Swap the values
					orderings_to_fix[x][seen[rule]] = num
					orderings_to_fix[x][seen[num]] = rule

					// Recalculate Seen
					seen = calculate_seen(update)
					j = -1;
					break
				}
			}
		}
	}

	valid_score := 0
	fixed_score := 0

	for _, order := range valid_orderings{
		len := len(order)
		len = len / 2
		valid_score += order[len]
	}
	for _, order := range orderings_to_fix{
		len := len(order)
		len = len / 2
		fixed_score += order[len]
	}

	fmt.Println("=======")
	fmt.Println(valid_score)
	fmt.Println(fixed_score)
}


func calculate_seen(update []int)map[int]int{
	seen := make(map[int]int)
	for i, value := range update{
		seen[value] = i
	}
	return seen
}