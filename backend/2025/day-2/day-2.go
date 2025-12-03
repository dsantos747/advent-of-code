package day2

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func part1(input []string) int {
	total := 0

	// brute force here we go
	for _, line := range input {
		split := strings.Split(line, "-")
		if len(split) != 2 {
			panic("weird input")
		}

		start, err1 := strconv.Atoi(split[0])
		end, err2 := strconv.Atoi(split[1])
		if err1 != nil || err2 != nil {
			panic(errors.Join(err1, err2))
		}

		for i := start; i <= end; i++ {
			if hasPattern(i) {
				total += i
			}
		}
	}

	return total
}

func part2(input []string) int {
	total := 0

	// brute force here we go
	for _, line := range input {
		split := strings.Split(line, "-")
		if len(split) != 2 {
			panic("weird input")
		}

		start, err1 := strconv.Atoi(split[0])
		end, err2 := strconv.Atoi(split[1])
		if err1 != nil || err2 != nil {
			panic(errors.Join(err1, err2))
		}

		for i := start; i <= end; i++ {
			if hasPatternV2(i) {
				total += i
			}
		}
	}

	return total
}

func hasPattern(num int) bool {
	numStr := strconv.Itoa(num)

	strLen := len(numStr)

	if strLen%2 != 0 {
		return false
	}

	if numStr[:strLen/2] != numStr[strLen/2:] {
		return false
	}

	return true
}

func hasPatternV2(num int) bool {
	if num < 11 { // smallest repeating pattern is 11
		return false
	}

	numStr := strconv.Itoa(num)
	strLen := len(numStr)
	divisors := getDivisors(strLen)

	for _, d := range divisors {
		segment := ""
		found := true

		for j := 0; j+d <= strLen; j += d {
			if segment == "" {
				segment = numStr[j : j+d]
				continue
			}
			if segment != numStr[j:j+d] {
				found = false
			}
		}
		if found {
			return true
		}
	}

	return false
}

func getDivisors(num int) []int {
	res := []int{1}

	for i := 2; i < math.MaxInt; i++ {
		if num/i <= 1 {
			break
		}
		if num%i == 0 {
			res = append(res, i)
			continue
		}

	}

	return res
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, ",")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
