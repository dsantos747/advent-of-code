package day13

import (
	"strings"
)

func checkRows(lines []string) int {
	for i := 0; i < len(lines)-1; i++ {
		if lines[i] != lines[i+1] {
			continue
		}
		for j := 1; ; j++ {
			if (i-j < 0) || (i+1+j > len(lines)-1) {
				return (i + 1)
			}
			if lines[i-j] != lines[i+1+j] {
				break
			}
		}
	}
	return 0
}

func transpose(lines []string) []string {
	splitLines := [][]string{}
	for _, line := range lines {
		a := strings.Split(line, "")
		splitLines = append(splitLines, a)
	}
	lx := len(splitLines[0])
	ly := len(splitLines)
	result := make([][]string, lx)
	for i := range result {
		result[i] = make([]string, ly)
	}
	for i := 0; i < lx; i++ {
		for j := 0; j < ly; j++ {
			result[i][j] = splitLines[j][i]
		}
	}
	newLines := []string{}
	for i := 0; i < lx; i++ {
		newLines = append(newLines, strings.Join(result[i], ""))
	}
	return newLines
}

func offByOne(l1, l2 string) bool {
	diffCount := 0
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			diffCount++
		}
		if diffCount > 1 {
			diffCount = 0
			return false
		}
	}
	if diffCount == 1 {
		return true
	} else {
		return false
	}
}

func checkRows2(lines []string) int {
	for i := 0; i < len(lines)-1; i++ {
		smudgeFound := false
		if lines[i] != lines[i+1] {
			if !(offByOne(lines[i], lines[i+1])) {
				continue
			}
			smudgeFound = true
		}

		if smudgeFound {
			for j := 1; ; j++ {
				if (i-j < 0) || (i+1+j > len(lines)-1) {
					return (i + 1)
				}
				if lines[i-j] != lines[i+1+j] {
					break
				}
			}
		} else {
			for j := 1; ; j++ {
				if (i-j < 0) || (i+1+j > len(lines)-1) {
					if smudgeFound {
						return (i + 1)
					} else {
						break
					}
				}
				if lines[i-j] != lines[i+1+j] {
					if !(offByOne(lines[i-j], lines[i+1+j])) {
						break
					}
					smudgeFound = true
				}
			}
		}
	}
	return 0
}

func part1(patterns []string) int {
	result := 0
	for _, pattern := range patterns {
		lines := strings.Split(pattern, "\n")
		rowSum := checkRows(lines) * 100
		colSum := 0
		if rowSum == 0 {
			cols := transpose(lines)
			colSum = checkRows(cols)
		}
		result += rowSum + colSum
	}
	return result
}

func part2(patterns []string) int {
	result := 0
	for _, pattern := range patterns {
		lines := strings.Split(pattern, "\n")
		rowSum := checkRows2(lines) * 100
		colSum := 0
		if rowSum == 0 {
			cols := transpose(lines)
			colSum = checkRows2(cols)
		}
		result += rowSum + colSum
	}
	return result
}

func Solve(data string) (int, int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return p1, p2, nil
}
