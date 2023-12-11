package main

import (
	// "AOC23/tools"
	"fmt"
	"strconv"
	"strings"

	tools "github.com/dsantos747/advent-of-code-2023/api/tools"
)

type Galaxy struct {
	num int
	x int
	y int
}

type Space struct {
	ind int
	size int
}

func mapRows(input []string, factor int) []Space {
	var rowMap []Space
	for i,line := range input {
		if (!strings.Contains(line,"#")) {
			rowMap = append(rowMap, Space{i,factor})
		} else {
			rowMap = append(rowMap, Space{i,1})
		}
	}
	return rowMap
}

func mapCols(input []string, factor int) []Space {
	var colMap []Space
	for i:=0; i<len(input[0]); i++ {
		colMap = append(colMap, Space{i,factor})
	}
	for _,line := range input {
		for i,char := range line {
			if string(char) == "#" {
				colMap[i].size = 1
			}
		}
	}
	return colMap
}

func mapGals(input []string) ([]Galaxy){
	galaxyMap := []Galaxy{}
	galaxyCount := 1
	for i,line := range input {
		a := strings.Split(line,"")
		for j,char := range a {
			if string(char) == "#" {
				a[j] = strconv.Itoa(galaxyCount)
				galaxyMap = append(galaxyMap, Galaxy{galaxyCount,j,i})
				galaxyCount++
			}
		}
	}
	return galaxyMap
}

func calcDistance(g1, g2 Galaxy, rowMap, colMap []Space) int {
	xSum := 0
	ySum := 0
	if (g2.x >= g1.x) {
		for i:=g1.x; i<g2.x; i++ {
			xSum += colMap[i].size
		}
	} else {
		for i:=g2.x; i<g1.x; i++ {
			xSum += colMap[i].size
		}
	}
	for i:=g1.y; i<g2.y; i++ {
		ySum += rowMap[i].size
	}
	return xSum + ySum
}

func part1(input []string) (int){
	result := 0
	expansionFactor := 2 // Replace one empty row with 2 empty rows
	rowSlice := mapRows(input,expansionFactor)
	colSlice := mapCols(input,expansionFactor)
	galaxyMap := mapGals(input)
	for i,g1 := range galaxyMap {
		for j:=i+1; j<len(galaxyMap); j++ {
			g2 := galaxyMap[j]
			result += calcDistance(g1,g2,rowSlice,colSlice)
		}
	}
	return result
}

func part2(input []string) (int) {
	result := 0
	expansionFactor := 1000000 // Replace one empty row with 1000000 empty rows
	rowSlice := mapRows(input,expansionFactor)
	colSlice := mapCols(input,expansionFactor)
	galaxyMap := mapGals(input)
	for i,g1 := range galaxyMap {
		for j:=i+1; j<len(galaxyMap); j++ {
			g2 := galaxyMap[j]
			result += calcDistance(g1,g2,rowSlice,colSlice)
		}
	}
	return result
}

func main() {
	data,err := tools.ReadInput("./input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input := strings.Split(data, "\n")

	p1 := part1(input)
	fmt.Println("The answer to part 1 is",p1)

	p2 := part2(input)
	fmt.Println("The answer to part 2 is",p2)
}