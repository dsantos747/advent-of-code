package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "strings"
)

func checkAdjacent(num NumStruct, symInfo [][]int) int {
	partNum := 0
	currLine := num.Line -1
	if currLine > 0 {
		for _, val := range symInfo[currLine-1] {
			if val >= (num.IndexArr[0])-1 && val <= (num.IndexArr[1]) {
				partNum,_ = strconv.Atoi(num.Num)
				return partNum
			}
			if val > (num.IndexArr[1]) {
				break
			}
		}
	}
	for _, val := range symInfo[currLine] {
		if val >= (num.IndexArr[0])-1 && val <= (num.IndexArr[1]) {
			partNum,_ = strconv.Atoi(num.Num)
			return partNum
		}
		if val > (num.IndexArr[1]) {
			break
		}
	}
	if currLine < len(symInfo) -1 {
		for _, val := range symInfo[currLine+1] {
			if  val >= (num.IndexArr[0])-1 && val <= (num.IndexArr[1]) {
				partNum,_ = strconv.Atoi(num.Num)
				return partNum
			}
			if val > (num.IndexArr[1]) {
				break
			}
		}
	}

	return partNum
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

func parseSyms(line string, lineNum int, symInfo []int ) []int{
	regSym := regexp.MustCompile(`[^0-9.]`)
	indSyms := (regSym.FindAllIndex([]byte(line),-1))

	for _,match := range indSyms {
		symInfo = append(symInfo, match[0])
	}

	return symInfo
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
	var symLineInfo []int
	var numInfo [][]NumStruct
	var symInfo [][]int

	sc := bufio.NewScanner(input)
	for sc.Scan() {
		lineNum ++
		parseNums(sc.Text(), lineNum, numLineInfo)
		parseSyms(sc.Text(), lineNum, symLineInfo)
		numInfo = append(numInfo,parseNums(sc.Text(), lineNum, numLineInfo))
		symInfo = append(symInfo,parseSyms(sc.Text(), lineNum, symLineInfo))
	}
	
	for _,line := range numInfo {
		for _,match := range line {
			partSum += checkAdjacent(match,symInfo)


		}
	}

	fmt.Println("The sum of valid engine parts is",partSum)
}