package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseHand(hand string, handMap map[string]int) map[string]int {
	

	pattern := `(\d+) (red|blue|green)`
	reg := regexp.MustCompile(pattern)
	matches := reg.FindAllStringSubmatch(hand,-1)

	for _,match := range matches {
		count,_ := strconv.Atoi(match[1])
		color := match[2]
		if (count > handMap[color]) {
			handMap[color] = count
		}
	}
	return handMap
}

func parseGame(game string) int {
	hands := strings.Split(game,";")
	handMap := map[string]int{"red": 0, "blue": 0, "green": 0}

	for _,hand := range hands {
		handMap = parseHand(hand,handMap)
	}
	result := handMap["red"] * handMap["blue"] * handMap["green"]

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

	fmt.Println("The sum of the power of the minimum set of cubes for each game is",gameSum)
}