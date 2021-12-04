package main

import (
	"fmt"
	"strings"
)

type instruction struct {}

func parseInstruction(in string) instruction {
	return instruction{}
}

var input []instruction


func main(){
	fmt.Printf("part 1: %v\n", partOne())
	fmt.Printf("part 2: %v\n", partTwo())
}

func partOne() int64{
	return 0
}

func partTwo() int64 {
	return 0
}

func init() {
	ins := strings.Split(_input, "\n")
	input := make([]instruction, len(ins))
	for i, v := range ins {
		input[i] = parseInstruction(v)
	}
}

var _input = ``