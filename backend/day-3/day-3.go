package day3

import (
	"regexp"
	"strconv"
	"strings"
)

type NumStruct struct {
	Num      string
	Line     int
	IndexArr []int
}

func checkAdjacent(num NumStruct, symInfo [][]int) int {
	partNum := 0
	currLine := num.Line
	if currLine > 0 {
		for _, val := range symInfo[currLine-1] {
			if val >= (num.IndexArr[0])-1 && val <= (num.IndexArr[1]) {
				partNum, _ = strconv.Atoi(num.Num)
				return partNum
			}
			if val > (num.IndexArr[1]) {
				break
			}
		}
	}
	for _, val := range symInfo[currLine] {
		if val >= (num.IndexArr[0])-1 && val <= (num.IndexArr[1]) {
			partNum, _ = strconv.Atoi(num.Num)
			return partNum
		}
		if val > (num.IndexArr[1]) {
			break
		}
	}
	if currLine < len(symInfo)-1 {
		for _, val := range symInfo[currLine+1] {
			if val >= (num.IndexArr[0])-1 && val <= (num.IndexArr[1]) {
				partNum, _ = strconv.Atoi(num.Num)
				return partNum
			}
			if val > (num.IndexArr[1]) {
				break
			}
		}
	}

	return partNum
}

func checkAdjacentNums(lineNum int, asterixIndex int, numInfo [][]NumStruct) int {
	partNum := 0
	var adj []int
	if lineNum > 0 {
		for _, numEntry := range numInfo[lineNum-1] {
			if asterixIndex >= (numEntry.IndexArr[0])-1 && asterixIndex <= (numEntry.IndexArr[1]) {
				partNum, _ = strconv.Atoi(numEntry.Num)
				adj = append(adj, partNum)
			}
		}
	}
	for _, numEntry := range numInfo[lineNum] {
		if asterixIndex >= (numEntry.IndexArr[0])-1 && asterixIndex <= (numEntry.IndexArr[1]) {
			partNum, _ = strconv.Atoi(numEntry.Num)
			adj = append(adj, partNum)
		}
	}
	if lineNum < len(numInfo)-1 {
		for _, numEntry := range numInfo[lineNum+1] {
			if asterixIndex >= (numEntry.IndexArr[0])-1 && asterixIndex <= (numEntry.IndexArr[1]) {
				partNum, _ = strconv.Atoi(numEntry.Num)
				adj = append(adj, partNum)
			}
		}
	}

	if len(adj) > 1 {
		ratio := adj[0] * adj[1] // From insection, we know len is never greater than 2
		return ratio
	}
	return 0
}

func parse(input []string, part2 bool) ([][]NumStruct, [][]int) {
	var numInfo [][]NumStruct
	var symInfo [][]int

	regNum := regexp.MustCompile(`\d+`)
	regSym := regexp.MustCompile(`[^0-9.]`)

	for i, line := range input {
		numMatches := (regNum.FindAll([]byte(line), -1))
		indMatches := (regNum.FindAllIndex([]byte(line), -1))
		var lineNums []NumStruct
		for j, match := range numMatches {
			lineNums = append(lineNums, NumStruct{string(match), i, indMatches[j]})
		}
		numInfo = append(numInfo, lineNums)

		var lineSyms []int
		if part2 {
			for j, char := range line {
				if char == '*' {
					lineSyms = append(lineSyms, j)
				}
			}
		} else {
			indSyms := (regSym.FindAllIndex([]byte(line), -1))
			for _, match := range indSyms {
				lineSyms = append(lineSyms, match[0])
			}
		}
		symInfo = append(symInfo, lineSyms)

	}
	return numInfo, symInfo
}

func part1(input []string) int {
	var partSum int = 0
	var numInfo [][]NumStruct
	var symInfo [][]int

	numInfo, symInfo = parse(input, false)

	for _, line := range numInfo {
		for _, match := range line {
			partSum += checkAdjacent(match, symInfo)
		}
	}

	return partSum
}

func part2(input []string) int {
	var partSum int = 0
	var numInfo [][]NumStruct
	var asterixInfo [][]int

	numInfo, asterixInfo = parse(input, true)

	for i, line := range asterixInfo {
		for _, match := range line {
			partSum += checkAdjacentNums(i, match, numInfo)
		}
	}

	return partSum
}

func Solve(data string) (int, int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return p1, p2, nil
}
