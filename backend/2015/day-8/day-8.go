package day8

import (
	"fmt"
	"strconv"
	"strings"
)

func part1(data []string) int {
	total := 0

	for _, line := range data {
		parsed, _ := strconv.Unquote(line)
		total += len(line) - len(parsed)
	}

	return total
}

func part2(data []string) int {
	total := 0

	for _, line := range data {
		escaped := fmt.Sprintf("%q", line)
		total += len(escaped) - len(line)
	}

	return total
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
