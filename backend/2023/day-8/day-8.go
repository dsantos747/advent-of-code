package day8

import (
	"strconv"
	"strings"

	tools "github.com/dsantos747/advent-of-code/tools"
)

func part1(data []string) int {
	dirs := strings.ReplaceAll(strings.ReplaceAll(data[0], "L", "0"), "R", "1")
	routes := data[2:]

	mapMap := make(map[string][]string)
	for _, route := range routes {
		r := strings.Split(strings.ReplaceAll(strings.ReplaceAll(route, "(", ""), ")", ""), "=")
		key := strings.TrimSpace(r[0])
		branches := strings.Split(r[1], ",")
		brLeft := strings.TrimSpace(branches[0])
		brRight := strings.TrimSpace(branches[1])
		mapMap[key] = []string{brLeft, brRight}
	}

	nextKey := "AAA"
	stepCount := 0
	for {
		for _, d := range dirs {
			dir, _ := strconv.Atoi(string(d))
			nextKey = mapMap[nextKey][dir]
			stepCount++
			if nextKey == "ZZZ" {
				return stepCount
			}

		}
	}
}

func part2(data []string) int {
	dirs := strings.ReplaceAll(strings.ReplaceAll(data[0], "L", "0"), "R", "1")
	routes := data[2:]

	mapMap := make(map[string][]string)
	var keyList []string
	for _, route := range routes {
		r := strings.Split(strings.ReplaceAll(strings.ReplaceAll(route, "(", ""), ")", ""), "=")
		key := strings.TrimSpace(r[0])
		if string(key[2]) == "A" {
			keyList = append(keyList, string(key))
		}
		branches := strings.Split(r[1], ",")
		brLeft := strings.TrimSpace(branches[0])
		brRight := strings.TrimSpace(branches[1])
		mapMap[key] = []string{brLeft, brRight}
	}

	countSlice := make([]int, len(keyList))
	stepCount := 0
	for {
		for _, d := range dirs {
			dir, _ := strconv.Atoi(string(d))
			stepCount++

			for i, key := range keyList {
				keyList[i] = mapMap[key][dir]
				if string(keyList[i][2]) == "Z" {
					countSlice[i] = stepCount
				}

			}
			lcm := tools.LCM(countSlice)
			if lcm > 0 {
				return lcm
			}

		}
	}
}

// Note - comment this out if using p2 test input
// NEED TO FIX THIS - ENSURE IT WORKS FOR BOTH TEST AND NON TEST

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
