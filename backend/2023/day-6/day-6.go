package day6

import (
	"math"
	"strconv"
	"strings"
)

func part1(input []string) int {
	times := strings.Fields(strings.Split(input[0], ":")[1])
	dists := strings.Fields(strings.Split(input[1], ":")[1])

	result := 1

	for i, time := range times {
		t, _ := strconv.ParseFloat(time, 32)
		d, _ := strconv.ParseFloat(dists[i], 32)
		d += 0.0000000001

		minButton := math.Ceil((t - math.Sqrt(math.Pow(t, 2)-(4*d))) / 2)
		maxButton := math.Floor((t + math.Sqrt(math.Pow(t, 2)-(4*d))) / 2)

		result *= (int(maxButton) - int(minButton) + 1)
	}

	return result
}

func part2(input []string) int {
	time := strings.Join(strings.Fields(strings.Split(input[0], ":")[1]), "")
	dist := strings.Join(strings.Fields(strings.Split(input[1], ":")[1]), "")

	t, _ := strconv.ParseFloat(time, 64)
	d, _ := strconv.ParseFloat(dist, 64)
	d += 0.0000000001

	minButton := math.Ceil((t - math.Sqrt(math.Pow(t, 2)-(4*d))) / 2)
	maxButton := math.Floor((t + math.Sqrt(math.Pow(t, 2)-(4*d))) / 2)

	result := (int(maxButton) - int(minButton) + 1)

	return result
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
