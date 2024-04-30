package day1

import (
	"strings"
)

func part1(input string) int {
	opens := strings.Count(input, "(")
	closes := len(input) - opens
	return opens - closes
}

func part2(input string) int {
	var delta int

	for i, c := range input {
		if c == '(' {
			delta++
		} else {
			delta--
		}
		if delta < 0 {
			return i + 1
		}
	}

	return 0
}

func Solve(data string) (*int, *int, error) {
	p1 := part1(data)
	p2 := part2(data)

	return &p1, &p2, nil
}
