package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var adapters []int

func main() {
	fmt.Printf("Part 1: %v\n", part1(adapters))
	fmt.Printf("Part 2: %v\n", part2(adapters))
}

func part1(inputs []int) (t int) {
	diffs := append([]int{}, 0, 0, 0, 1) // seed with device built-in adapter
	var joltage int

	for _, v := range inputs {
		var diff int
		diff, joltage = v-joltage, v
		if diff > 3 {
			fmt.Printf("oops: %v, %v", diff, joltage)
		}
		diffs[diff] += 1
	}

	return diffs[1] * diffs[3]
}

func part2(inputs []int) (t int) {
	// for every joltage, that number of combinations = sum(number of remaining combinations given a choice)
	seen := map[int]int{}

	var iterate func(j int, ins []int) (t int)
	iterate = func(j int, ins []int) (t int) {
		for i, v := range ins {
			if v-j > 3 {
				break
			}
			if _, ok := seen[v]; !ok {
				if i == len(ins)-1 {
					seen[v] = 1
				} else {
					seen[v] = iterate(v, ins[i+1:])
				}
			}
			t += seen[v]
		}
		return t
	}

	return iterate(0, inputs)

}

func init() {
	inputs := strings.Split(input, "\n")
	adapters = make([]int, len(inputs))
	for i, v := range inputs {
		adapters[i], _ = strconv.Atoi(v)
	}
	sort.IntSlice(adapters).Sort()
}

//var input = `28
//33
//18
//42
//31
//14
//46
//20
//48
//47
//24
//23
//49
//45
//19
//38
//39
//11
//1
//32
//25
//35
//8
//17
//7
//9
//4
//2
//34
//10
//3`

var input = `107
13
116
132
24
44
56
69
28
135
152
109
42
112
10
43
122
87
49
155
175
71
39
173
50
156
120
145
176
45
149
148
15
1
68
9
168
131
150
59
83
167
3
169
6
123
174
81
138
72
157
144
65
75
33
19
140
160
16
57
93
90
8
58
98
130
141
114
84
29
22
94
113
129
108
36
14
115
102
151
78
139
170
82
2
70
126
101
25
62
95
104
23
163
32
103
121
119
48
166
7
53`
