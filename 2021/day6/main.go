package main

import (
	"fmt"
	"strconv"
	"strings"
)

type lf int

func parseLF(in string) lf {
	val, _ := strconv.Atoi(in)
	return lf(val)
}

var input []lf


func main(){
	fmt.Printf("part 1: %v\n", partOne())
	fmt.Printf("part 2: %v\n", partTwo())
}

func partOne() int{
	fish := input

	for i := 1; i<=80; i++ {
		var _fish []lf
		for _, f := range fish {
			if f == 0 {
				_fish = append(_fish, 6, 8 )
			} else {
				f--
				_fish = append(_fish, f)
			}
		}
		fish = _fish
	}
	return len(fish)
}

var fishBecomes = map[string]int{}

func partTwo() (count int) {
	for _, f := range input {
		count += getLineage(f, 256)
	}
	return count
}

func getLineage(f lf, days int) int {
	key := fmt.Sprintf("%d,%d", f, days)
	if _, ok := fishBecomes[key]; !ok {
		var val int
		_days := days - 1
		if days == 0 {
			val = 1
		} else if f == 0 {
			val = getLineage(6, _days) + getLineage(8, _days)
		} else {
			f--
			val = getLineage(f, _days)
		}

		fishBecomes[key] = val
	}

	return fishBecomes[key]
}



func init() {
	ins := strings.Split(_input, ",")
	//ins := strings.Split(_inputSimple, ",")
	input = make([]lf, len(ins))
	for i, v := range ins {
		input[i] = parseLF(v)
	}
}

var _inputSimple = `3,4,3,1,2`
var _input = `1,1,1,2,1,5,1,1,2,1,4,1,4,1,1,1,1,1,1,4,1,1,1,1,4,1,1,5,1,3,1,2,1,1,1,2,1,1,1,4,1,1,3,1,5,1,1,1,1,3,5,5,2,1,1,1,2,1,1,1,1,1,1,1,1,5,4,1,1,1,1,1,3,1,1,2,4,4,1,1,1,1,1,1,3,1,1,1,1,5,1,3,1,5,1,2,1,1,5,1,1,1,5,3,3,1,4,1,3,1,3,1,1,1,1,3,1,4,1,1,1,1,1,2,1,1,1,4,2,1,1,5,1,1,1,2,1,1,1,1,1,1,1,1,2,1,1,1,1,1,5,1,1,1,1,3,1,1,1,1,1,3,4,1,2,1,3,2,1,1,2,1,1,1,1,4,1,1,1,1,4,1,1,1,1,1,2,1,1,4,1,1,1,5,3,2,2,1,1,3,1,5,1,5,1,1,1,1,1,5,1,4,1,2,1,1,1,1,2,1,3,1,1,1,1,1,1,2,1,1,1,3,1,4,3,1,4,1,3,2,1,1,1,1,1,3,1,1,1,1,1,1,1,1,1,1,2,1,5,1,1,1,1,2,1,1,1,3,5,1,1,1,1,5,1,1,2,1,2,4,2,2,1,1,1,5,2,1,1,5,1,1,1,1,5,1,1,1,2,1`