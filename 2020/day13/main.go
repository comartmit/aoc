package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	start, buses := parse(input)
	fmt.Printf("part 1: %v \n", part1(start, buses))
	fmt.Printf("part 2: %v \n", part2(buses))
}

func part1(start int, buses []string) (t int) {
	bs := map[int]int{}
	for _, b := range buses {
		if b == "x" {
			continue
		}
		bid, _ := strconv.Atoi(b)
		bs[bid] = 0
	}

	bestbid, best := 0, int(math.MaxInt64)
	for bid, _ := range bs {
		depart := ((start / bid) + 1) * bid
		if depart < best {
			bestbid = bid
			best = depart
		}
	}
	return bestbid * (best - start)
}

func part2(bs []string) (t int) {
	increment := 1

	for i, b := range bs {
		if b == "x" {
			continue
		}
		bus, _ := strconv.Atoi(b)
		for ((t+i)%bus) != 0 || t == 0 {
			t += increment
		}
		increment *= bus
	}

	return t
}

func parse(input string) (int, []string) {
	parts := strings.Split(input, "\n")
	depart, _ := strconv.Atoi(parts[0])
	buses := strings.Split(parts[1], ",")

	return depart, buses
}

var input = `1005595
41,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,557,x,29,x,x,x,x,x,x,x,x,x,x,13,x,x,x,17,x,x,x,x,x,23,x,x,x,x,x,x,x,419,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,19`

var inputSimple = `11
7,13,x,x,59,x,31,19`
