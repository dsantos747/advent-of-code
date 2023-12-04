package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	// "regexp"
	"strconv"
	"strings"
)

func parseGame(game string) int{
	scoreExp:= -1
	win:= false

	game = game[9:]
	splitGame := strings.Split(game,"|")
	gameNums := strings.Fields(splitGame[0])
	myNums := strings.Fields(splitGame[1])

	for _,g := range gameNums {
		gameNum,_ := strconv.Atoi(g)
		for _,m := range myNums {
			myNum,_ := strconv.Atoi(m)
			if (gameNum == myNum) {
				scoreExp++
				win = true
			}
		}
	}

	if (win) {
		return int(math.Pow(2, float64(scoreExp)))
	}
	return 0
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
	

	fmt.Println("The sum of valid points is",gameSum)
}