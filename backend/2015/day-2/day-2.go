package day2

import (
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func part1(input []string) int {
	total := 0

	for _, line := range input {
		var dims [3]int
		var maxi int
		dimStr := strings.Split(line, "x")
		for i, d := range dimStr {
			dims[i], _ = strconv.Atoi(d)
			maxi = max(maxi, dims[i])
		}
		total += (2 * dims[0] * dims[1]) + (2 * dims[1] * dims[2]) + (2 * dims[2] * dims[0]) + ((dims[0] * dims[1] * dims[2]) / maxi)
	}
	return total
}

func part2(input []string) int {
	total := 0

	for _, line := range input {
		var dims [3]int
		var maxi int
		dimStr := strings.Split(line, "x")
		for i, d := range dimStr {
			dims[i], _ = strconv.Atoi(d)
			maxi = max(maxi, dims[i])
		}
		total += (dims[0] * dims[1] * dims[2]) + 2*(dims[0]+dims[1]+dims[2]-maxi)
	}
	return total
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
