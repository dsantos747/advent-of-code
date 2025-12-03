package day3

import (
	"math"
	"strconv"
	"strings"
)

func part1(input []string) int {
	totalJolts := 0
	for _, bank := range input {
		maxJoltage := getMaxJoltage(bank, 2)
		totalJolts += maxJoltage
	}

	return totalJolts
}

func part2(input []string) int {
	totalJolts := 0
	for _, bank := range input {
		maxJoltage := getMaxJoltage(bank, 12)
		totalJolts += maxJoltage
	}

	return totalJolts
}

func getMaxJoltage(bank string, n int) int {
	joltVals := []string{"9", "8", "7", "6", "5", "4", "3", "2", "1"}

	for _, joltVal1 := range joltVals {
		ind := strings.Index(bank, joltVal1)
		if ind == -1 || ind+n > len(bank) {
			continue
		}

		jolt, err := strconv.Atoi(string(bank[ind]))
		if err != nil {
			panic("parsing int")
		}

		return jolt*int(math.Pow10(n-1)) + getMaxJoltage(bank[ind+1:], n-1)
	}
	return 0
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
