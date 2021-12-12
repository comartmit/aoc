package main

import (
	"fmt"
	"strconv"
	"strings"
)

type octopus struct {
	val int
}

var input [][]*octopus // pointers for determining "nil"


func main(){
	Init()
	fmt.Printf("part 1: %v\n", partOne())

	// starts at 100
	fmt.Printf("part 2: %v\n", partTwo())
}

func playRound(os [][]*octopus) (count int){
	// prepare next round and apply step
	oshined := make([][]bool, len(os))
	for i, row := range os {
		oshined[i] = make([]bool, len(row))
		for _, o := range row {
			o.val++
		}
	}

	// evaluate chain reaction
	shined := true
	for shined {
		shined = false // reset round
		for i, row := range os {
			for j, o := range row {
				// only shine once per round
				if o.val > 9 && !oshined[i][j] {
					oshined[i][j] = true
					incrementNeighbors(os, i, j)
					shined = true
					count++
				}
			}
		}
	}

	// copy unshined octopuses to next grid
	for i, row := range os {
		for j, o := range row {
			// only shine once per round
			if oshined[i][j] {
				o.val = 0
			}
		}
	}

	return count
}

func partOne() (count int){
	for k := 0; k < 100; k++ {
		count += playRound(input)
	}

	return count
}

func partTwo() int {
	var count, step = 0, 100

	for count < 100 {
		count = playRound(input)
		step++
	}

	return step
}

func printGrid(os[][]*octopus) {
	for _, v := range os {
		for _, o := range v {
			fmt.Print(o.val)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func inGrid(os[][]*octopus, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(os) && y < len(os[0])
}

func incrementNeighbors(os [][]*octopus, x, y int) {
	for _x := x-1; _x <= x+1; _x++ {
		for _y := y-1; _y <= y+1; _y++ {
			if !(_x == x && _y == y) && inGrid(os, _x, _y){
				os[_x][_y].val += 1
			}
		}
	}
}

func Init() {
	ins := strings.Split(_input, "\n")
	input = make([][]*octopus, len(ins))
	for i, v := range ins {
		row := make([]*octopus, len(v))
		for j, v := range v {
			o, _ := strconv.Atoi(string(v))
			row[j] = &octopus{o}
		}
		input[i] = row
	}
}

var _input = `3322874652
5636588857
7755117548
5854121833
2856682477
3124873812
1541372254
8634383236
2424323348
2265635842`

var _input_test = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`