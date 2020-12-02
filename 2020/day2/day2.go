package main

import (
	"regexp"
	"testing"
)

func parse(s string) []string {
	// 1-4 s: lssss
	re := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) == 0 {
		return matches
	}
	return matches[1:]
}

func Test(t *testing.T) {
	// type test struct {
	// 	input  int
	// 	expect int
	// }

	// tests := []test{
	// 	test{
	// 		// ...
	// 	},
	// }

	// for i, test := range tests {
	// 	t.Run(fmt.Sprint(i), func(t *testing.T) {
	// 		require.Equal(t, test.expect, test.input)
	// 	})
	// }
}
