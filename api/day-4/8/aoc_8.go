package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseGame(game string) int{
	count:= 0
	
	game = game[9:]
	splitGame := strings.Split(game,"|")
	gameNums := strings.Fields(splitGame[0])
	myNums := strings.Fields(splitGame[1])
	for _,g := range gameNums {
		gameNum,_ := strconv.Atoi(g)
		for _,m := range myNums {
			myNum,_ := strconv.Atoi(m)
			if (gameNum == myNum) {
				count++
			}
		}
	}
	return count
}

func main() {
	input,err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error loading input", err)
		return
	}
	defer input.Close()

	var gameSum int = 0
	var gameLen int = 187 // CHANGE THIS - Used because was struggling reset the scanner
	gameArr := make([]int,0)

	for i:=0; i<gameLen; i++ {
		gameArr = append(gameArr, 1)
	}
	
	var gameInd int = 0
	sc := bufio.NewScanner(input)
	
	for sc.Scan() {
		gameScore := parseGame(sc.Text())
		for i := 0; i<gameScore; i++ {
			gameArr[gameInd + i +1] += gameArr[gameInd]
		}
		gameSum += gameArr[gameInd]
		gameInd++
	}
	

	fmt.Println("The sum of valid points is",gameSum)
}