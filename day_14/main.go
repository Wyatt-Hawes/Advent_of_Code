package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type robot struct{
	x int
	y int
	v_x int
	v_y int
}

// var WIDTH = 11
// var HEIGHT = 7

var WIDTH = 101
var HEIGHT = 103

func main(){
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	robots := make([]robot,0)

	for scanner.Scan(){
		x, y, v_x, v_y := 0,0,0,0
		fmt.Sscanf(scanner.Text(),"p=%d,%d v=%d,%d",&x,&y,&v_x,&v_y)
		robots = append(robots, robot{x,y,v_x,v_y})
	}
	// Part 1
	p1_robots := tick(robots,100)
	TL, TR, BL, BR := calculate_robots_each_quadrant(p1_robots)
	fmt.Println(TL,TR,BL,BR)
	fmt.Println("Safety Factor:",TL * TR * BL * BR)

	// Part 2
	p2_robots := robots
	for i := range(10000){
		p2_robots = tick(p2_robots,1)
		if(print_game(p2_robots, true)){
			continue
		}
		fmt.Println("======",i+1,"======")
		break
	}
}

func tick(robots []robot, time int)[]robot{
	new_positions := make([]robot,0)

	for _, bot := range robots{
		x := ((bot.x + (bot.v_x * time)) % WIDTH)
		x = (x + WIDTH) % WIDTH
		y := (bot.y + (bot.v_y * time) + HEIGHT) % HEIGHT
		y = (y + HEIGHT) % HEIGHT
		new_positions = append(new_positions, robot{x,y,bot.v_x,bot.v_y})
	}
	return new_positions
}

func calculate_robots_each_quadrant(bots []robot)(int, int, int, int){
	TL, TR, BL, BR := 0,0,0,0
	for _, bot := range bots{
		x,y := bot.x, bot.y
		if x < WIDTH / 2 && y < HEIGHT / 2{
			TL++
		}
		if x > WIDTH / 2 && y < HEIGHT / 2{
			TR++
		}
		if x < WIDTH / 2 && y > HEIGHT / 2{
			BL++
		}
		if x > WIDTH / 2 && y > HEIGHT / 2{
			BR++
		}
	}
	return TL, TR, BL, BR
}

func print_game(bots []robot, skip_early bool)bool{
	game := make([][]int,HEIGHT)

	// Setup
	for i, _ := range game{
		game[i] = make([]int,WIDTH)
		for x := range WIDTH{
			game[i][x] = 0
		}
	}

	// Count, dont print if any index is > 1
	for _, bot := range bots{
		game[bot.y][bot.x]++;
		if(game[bot.y][bot.x] > 1 && skip_early){
			return true
		}
	}

	// Print
	s := ""
	fmt.Println("==============")
	for y := range HEIGHT{
		for x := range WIDTH{
			v := game[y][x]
			if(v == 0){
				s += " "
			}else{
				s += strconv.Itoa(v)
			}
		}
		s += string('\n')
	}
	fmt.Println(s)
	return false
}