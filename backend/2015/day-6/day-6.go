package day6

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	action     string
	start, end [2]int
}

func parseRange(input string) [2]int {
	split := strings.Split(input, ",")
	i, _ := strconv.Atoi(split[0])
	j, _ := strconv.Atoi(split[1])

	return [2]int{i, j}
}

func parseInstructions(input []string) []instruction {
	var instructions []instruction

	for _, line := range input {
		split := strings.Split(line, " ")
		if len(split) == 5 {
			split = split[1:]
		}

		start := parseRange(split[1])
		end := parseRange(split[3])

		i := instruction{split[0], start, end}
		instructions = append(instructions, i)
	}
	return instructions
}

func part1(input []string) int {
	instructions := parseInstructions(input)
	lit := 0

	grid := [1000][1000]bool{}

	for _, step := range instructions {
		for i := step.start[0]; i <= step.end[0]; i++ {
			for j := step.start[1]; j <= step.end[1]; j++ {
				if step.action == "toggle" {
					grid[i][j] = !grid[i][j]
				} else if step.action == "on" {
					grid[i][j] = true
				} else if step.action == "off" {
					grid[i][j] = false
				} else {
					fmt.Println("uh oh")
				}
			}
		}
	}

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] {
				lit++
			}
		}
	}

	return lit
}

func part2(input []string) int {
	instructions := parseInstructions(input)
	brightness := 0

	grid := [1000][1000]int{}

	for _, step := range instructions {
		for i := step.start[0]; i <= step.end[0]; i++ {
			for j := step.start[1]; j <= step.end[1]; j++ {
				if step.action == "toggle" {
					grid[i][j] += 2
				} else if step.action == "on" {
					grid[i][j]++
				} else if step.action == "off" && grid[i][j] > 0 {
					grid[i][j]--
				}
			}
		}
	}

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			brightness += grid[i][j]
		}
	}

	return brightness
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
