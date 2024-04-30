package day4

import (
	"math"
	"strconv"
	"strings"
)

func parseGameP2(game string) int {
	count := 0

	game = game[9:]
	splitGame := strings.Split(game, "|")
	gameNums := strings.Fields(splitGame[0])
	myNums := strings.Fields(splitGame[1])
	for _, g := range gameNums {
		gameNum, _ := strconv.Atoi(g)
		for _, m := range myNums {
			myNum, _ := strconv.Atoi(m)
			if gameNum == myNum {
				count++
			}
		}
	}
	return count
}

func part1(input []string) int {
	var gameSum int = 0

	for _, game := range input {
		scoreExp := -1
		win := false

		game = game[9:]
		splitGame := strings.Split(game, "|")
		gameNums := strings.Fields(splitGame[0])
		myNums := strings.Fields(splitGame[1])

		for _, g := range gameNums {
			gameNum, _ := strconv.Atoi(g)
			for _, m := range myNums {
				myNum, _ := strconv.Atoi(m)
				if gameNum == myNum {
					scoreExp++
					win = true
				}
			}
		}

		if win {
			gameSum += int(math.Pow(2, float64(scoreExp)))
		}
	}

	return gameSum
}

func part2(input []string) int {
	var gameSum int = 0

	gameArr := make([]int, 0)

	for i := 0; i < len(input); i++ {
		gameArr = append(gameArr, 1)
	}

	for i, game := range input {
		gameScore := parseGameP2(game)
		for j := 0; j < gameScore; j++ {
			gameArr[i+j+1] += gameArr[i]
		}
		gameSum += gameArr[i]
	}

	return gameSum
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
