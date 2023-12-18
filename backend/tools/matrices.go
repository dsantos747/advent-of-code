package tools

import (
	"strings"
)

func Transpose(lines []string) []string {
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

func Zeros(m, n int) [][]int {
	var zeros [][]int
	for i := 0; i < m; i++ {
		zeros = append(zeros, make([]int, n))
		for j := 0; j < n; j++ {
			zeros[i][j] = 0
		}
	}
	return zeros
}
