package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseHand(hand string) bool {
	handMap := map[string]int{"red": 0, "blue": 0, "green": 0}

	pattern := `(\d+) (red|blue|green)`
	reg := regexp.MustCompile(pattern)
	matches := reg.FindAllStringSubmatch(hand,-1)

	for _,match := range matches {
		count,_ := strconv.Atoi(match[1])
		color := match[2]
		handMap[color] += count
	}
	if (handMap["red"]>12 || handMap["green"]>13 || handMap["blue"]>14) {
		return false
	}

	return true
}

func parseGame(game string) int {
	reg := regexp.MustCompile(`\d+`)
	result,_ := strconv.Atoi(string(reg.Find([]byte(game))))

	hands := strings.Split(game,";")

	for _,hand := range hands {
		handResult := parseHand(hand)
		if (!handResult) {
			return 0
		}
	}

	return result
}

func main() {
	input,err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error loading input", err)
		return
	}
	defer input.Close()

	var gameSum int = 0


	sc := bufio.NewScanner(input)
	for sc.Scan() {
		gameSum += parseGame(sc.Text())
	}

	fmt.Println("The sum of all valid game numbers is",gameSum)
}