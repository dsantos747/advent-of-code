package day2

import (
	"strconv"
	"strings"
)

func parseHand(hand string) bool {
	handMap := map[string]int{"red": 0, "blue": 0, "green": 0}
	cubes := strings.Split(hand, ", ")

	var c []string
	var count int
	var colour string

	for _, cube := range cubes {
		c = strings.Fields(cube)
		count, _ = strconv.Atoi(c[0])
		colour = c[1]
		handMap[colour] += count
	}

	if handMap["red"] > 12 || handMap["green"] > 13 || handMap["blue"] > 14 {
		return false
	}

	return true
}

func parseHand2(hand string, handMap map[string]int) map[string]int {
	cubes := strings.Split(hand, ", ")

	var c []string
	var count int
	var colour string

	for _, cube := range cubes {
		c = strings.Fields(cube)
		count, _ = strconv.Atoi(c[0])
		colour = c[1]
		if count > handMap[colour] {
			handMap[colour] = count
		}
	}

	return handMap
}

func parseGame(game string) int {
	gameSplit := strings.Split(game, ": ")
	gameNo, _ := strconv.Atoi(gameSplit[0][5:])

	hands := strings.Split(gameSplit[1], "; ")

	for _, hand := range hands {
		handResult := parseHand(hand)
		if !handResult {
			return 0
		}
	}

	return gameNo
}

func parseGame2(game string) int {
	gameSplit := strings.Split(game, ": ")

	hands := strings.Split(gameSplit[1], "; ")
	handMap := map[string]int{"red": 0, "blue": 0, "green": 0}

	for _, hand := range hands {
		handMap = parseHand2(hand, handMap)
	}
	result := handMap["red"] * handMap["blue"] * handMap["green"]

	return result
}

func part1(input []string) int {
	gameSum := 0
	for _, line := range input {
		gameSum += parseGame(line)
	}
	return gameSum
}

func part2(input []string) int {
	gameSum := 0
	for _, line := range input {
		gameSum += parseGame2(line)
	}
	return gameSum
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
