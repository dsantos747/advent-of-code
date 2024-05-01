package day5

import (
	"strings"
)

func part1(input []string) int {
	total := 0
	var repeat bool
	var vowels int

	for _, line := range input {
		repeat = false
		vowels = 0
		var prev rune
		if strings.Contains(line, "ab") || strings.Contains(line, "cd") || strings.Contains(line, "pq") || strings.Contains(line, "xy") {
			continue
		}
		for _, c := range line {
			if c == prev {
				repeat = true
			}
			prev = c

			if c == rune('a') || c == rune('e') || c == rune('i') || c == rune('o') || c == rune('u') {
				vowels++
			}
		}

		if repeat && vowels >= 3 {
			total++
		}

	}

	return total
}

func part2(input []string) int {
	total := 0
	var triplet bool
	var pairs bool

	for _, line := range input {
		triplet = false
		pairs = false
		pairMap := map[string]int{}

		for i := 1; i < len(line); i++ {
			if i > 1 {
				if line[i-2] == line[i] {
					triplet = true
				}
			}
			sub := line[i-1 : i+1]

			if val, ok := pairMap[sub]; !ok {
				pairMap[sub] = i
			} else {
				if i-val >= 2 {
					pairs = true
				}
			}

		}

		if triplet && pairs {
			total++
		}

	}

	return total
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
