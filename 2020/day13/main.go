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

func parse(input string) (int, []string) {
	parts := strings.Split(input, "\n")
	depart, _ := strconv.Atoi(parts[0])
	buses := strings.Split(parts[1], ",")

	return depart, buses
}

var input = `1005595
41,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,557,x,29,x,x,x,x,x,x,x,x,x,x,13,x,x,x,17,x,x,x,x,x,23,x,x,x,x,x,x,x,419,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,19`
