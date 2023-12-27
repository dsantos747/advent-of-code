package day9

import (
	"strconv"
	"strings"
)

func subSeq(seq []int, end *bool) ([]int, bool) {
	var slice []int
	var val int
	*end = true
	for i := 1; i < len(seq); i++ {
		val = seq[i] - seq[i-1]
		slice = append(slice, val)
		if val != 0 {
			*end = false
		}
	}
	return slice, *end
}

func extrapolateFwd(seq string) [][]int {
	var end bool
	seqSlice := make([][]int, 1)
	strSlice := strings.Fields(seq)
	for _, val := range strSlice {
		v, _ := strconv.Atoi(string(val))
		seqSlice[0] = append(seqSlice[0], v)
	}

	i := 0
	for !end {
		var newSlice []int
		newSlice, end = subSeq(seqSlice[i], &end)
		seqSlice = append(seqSlice, newSlice)
		i++
	}
	for i := len(seqSlice) - 1; i > 0; i-- {
		seqSlice[i-1] = append(seqSlice[i-1], seqSlice[i-1][len(seqSlice[i-1])-1]+seqSlice[i][len(seqSlice[i])-1])
	}

	return seqSlice
}

func part1(data []string) int {
	result := 0
	for _, seq := range data {
		seqSlice := extrapolateFwd(seq)
		result += seqSlice[0][len(seqSlice[0])-1]
	}
	return result
}

func part2(data []string) int {
	result := 0
	for _, seq := range data {
		seqSlice := extrapolateFwd(seq)
		for i := len(seqSlice) - 1; i > 0; i-- {
			seqSlice[i-1] = append([]int{seqSlice[i-1][0] - seqSlice[i][0]}, seqSlice[i-1]...)
		}
		result += seqSlice[0][0]
	}
	return result
}

func Solve(data string) (int, int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return p1, p2, nil
}
