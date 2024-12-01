package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/dsantos747/advent-of-code/tools"
)

func part1(input []string) int {
	var l1, l2 []int
	for _, line := range input {
		nums := strings.Fields(line)
		if len(nums) != 2 {
			fmt.Println("invalid input")
			return -1
		}
		n1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("failed to parse string to int")
			return -1
		}
		n2, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println("failed to parse string to int")
			return -1
		}
		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}

	slices.Sort(l1)
	slices.Sort(l2)

	res := 0

	for i, n1 := range l1 {
		res += int(math.Abs(float64(n1 - l2[i])))
	}
	return res
}

func part2(input []string) int {
	var l1, l2 []int
	for _, line := range input {
		nums := strings.Fields(line)
		if len(nums) != 2 {
			fmt.Println("invalid input")
			return -1
		}
		n1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("failed to parse string to int")
			return -1
		}
		n2, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println("failed to parse string to int")
			return -1
		}
		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}

	counts := make(map[int]int)

	for _, n1 := range l1 {
		if _, ok := counts[n1]; ok {
			continue
		}
		total := 0
		for _, n2 := range l2 {
			if n2 == n1 {
				total++
			}
		}
		counts[n1] = total
	}

	res := 0
	for _, n1 := range l1 {
		res += counts[n1] * n1
	}

	return res
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}

func main() {
	data, err := tools.ReadInput("./2024/day-1/input.txt")
	if err != nil {
		panic(err)
	}
	p1, p2, err := Solve(data)
	fmt.Println("Part 1:", *p1)
	fmt.Println("Part 2:", *p2)
}
