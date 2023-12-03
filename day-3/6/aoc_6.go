package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "strings"
)

func checkAdjacentNums(lineNum int, asterixIndex int, numInfo [][]NumStruct) int {
	partNum := 0
	// ratio := 1
	var adj []int
	// fmt.Println("on line number", lineNum, asterixIndex)
	if lineNum > 0 {
		for _, numEntry := range numInfo[lineNum-1] {
			if asterixIndex >= (numEntry.IndexArr[0])-1 && asterixIndex <= (numEntry.IndexArr[1]) {
				partNum,_ = strconv.Atoi(numEntry.Num)
				adj = append(adj, partNum)
			}
		}
	}
	for _, numEntry := range numInfo[lineNum] {
		if asterixIndex >= (numEntry.IndexArr[0])-1 && asterixIndex <= (numEntry.IndexArr[1]) {
			partNum,_ = strconv.Atoi(numEntry.Num)
			adj = append(adj, partNum)
		}
	}
	if lineNum < len(numInfo) -1 {
		for _, numEntry := range numInfo[lineNum+1] {
			if  asterixIndex >= (numEntry.IndexArr[0])-1 && asterixIndex <= (numEntry.IndexArr[1]) {
				partNum,_ = strconv.Atoi(numEntry.Num)
				adj = append(adj, partNum)
			}
		}
	}
	
	if (len(adj) > 1) {
		// fmt.Println(adj)
		ratio := adj[0] * adj[1] // From insection, we know len is never greater than 2
		return ratio
	} else {
		return 0
	}
}

func parseNums(line string, lineNum int, numInfo []NumStruct ) []NumStruct{
	regNum := regexp.MustCompile(`\d+`)
	numMatches := (regNum.FindAll([]byte(line),-1))
	indMatches := (regNum.FindAllIndex([]byte(line),-1))

	for i,match := range numMatches {
		numInfo = append(numInfo, NumStruct{string(match),lineNum,indMatches[i]})
	}
	return numInfo
}

func parseAsterix(line string, lineNum int, asterixInfo []int ) []int{
	regAsterix := regexp.MustCompile(`\*`)
	indAsterix := (regAsterix.FindAllIndex([]byte(line),-1))

	for _,match := range indAsterix {
		asterixInfo = append(asterixInfo, match[0])
	}

	return asterixInfo
}

type NumStruct struct {
	Num   string
	Line   int
	IndexArr []int
}

func main() {
	input,err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error loading input", err)
		return
	}
	defer input.Close()


	var partSum int = 0
	var lineNum int = 0

	var numLineInfo []NumStruct
	var asterixLineInfo []int
	var numInfo [][]NumStruct
	var asterixInfo [][]int

	sc := bufio.NewScanner(input)
	for sc.Scan() {
		lineNum ++
		parseNums(sc.Text(), lineNum, numLineInfo)
		parseAsterix(sc.Text(), lineNum, asterixLineInfo)
		numInfo = append(numInfo,parseNums(sc.Text(), lineNum, numLineInfo))
		asterixInfo = append(asterixInfo,parseAsterix(sc.Text(), lineNum, asterixLineInfo))
	}
	
	for i,line := range asterixInfo {
		for _,match := range line {
			partSum += checkAdjacentNums(i,match,numInfo)
		}
	}

	fmt.Println("The total gear ratio is",partSum)
}