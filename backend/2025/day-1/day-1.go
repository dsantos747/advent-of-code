package day1

import (
	"math"
	"strconv"
	"strings"

	"github.com/dsantos747/advent-of-code/tools"
)

func part1(input []string) int {
	start := 50
	counter := 0

	for _, line := range input {
		if len(line) == 0 {
			continue
		}

		dir := string(line[0])
		num, err := strconv.Atoi(strings.TrimPrefix(line, dir))
		if err != nil {
			panic(err)
		}

		switch dir {
		case "L":
			start -= num
		case "R":
			start += num
		default:
			panic("unknown direction")
		}

		if int(math.Abs(float64(start)))%100 == 0 {
			counter++
		}

	}

	return counter
}

func part2(input []string) int {
	start := 50
	counter := 0

	for _, line := range input {
		if len(line) == 0 {
			continue
		}

		parsedLine := strings.ReplaceAll(strings.ReplaceAll(line, "R", ""), "L", "-")

		num, err := strconv.Atoi(parsedLine)
		if err != nil {
			panic(err)
		}

		counter += int(math.Floor(float64(tools.Abs(num)) / 100))

		num -= (num / 100 * 100)

		next := start + num

		switch {
		case next == 0 && start != 0:
			counter++
		case next > 99:
			counter++
		case next < 0 && start > 0:
			counter++
		}

		start = tools.Mod(next, 100)
	}

	return counter
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
