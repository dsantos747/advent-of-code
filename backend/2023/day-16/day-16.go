package day16

import (
	"fmt"
	"strings"
)

type Pos struct { // ROWS BY COLUMNS
	i int
	j int
}

type State struct {
	pos Pos // e.g. {0 1} = travelling right
	dir Pos
}

func lightPath(input []string, state State, cachePos [][]int, cache map[State]int) ([][]int, map[State]int) {

	if _, ok := cache[state]; ok {
		return cachePos, cache
	} else {
		cache[state] = 1
	}

	i := state.pos.i + state.dir.i
	j := state.pos.j + state.dir.j

	if (i < 0) || (i > len(input)-1) || (j < 0) || (j > len(input[0])-1) {
		return cachePos, cache
	}

	if cachePos[i][j] == 0 {
		cachePos[i][j] = 1
	}

	if input[i][j] == '|' && state.dir.j != 0 {
		cachePos, cache = lightPath(input, State{Pos{i, j}, Pos{-1, 0}}, cachePos, cache)
		cachePos, cache = lightPath(input, State{Pos{i, j}, Pos{1, 0}}, cachePos, cache)
	} else if input[i][j] == '-' && state.dir.i != 0 {
		cachePos, cache = lightPath(input, State{Pos{i, j}, Pos{0, -1}}, cachePos, cache)
		cachePos, cache = lightPath(input, State{Pos{i, j}, Pos{0, 1}}, cachePos, cache)
	} else if input[i][j] == '\\' {
		newDir := Pos{state.dir.j, state.dir.i}
		cachePos, cache = lightPath(input, State{Pos{i, j}, newDir}, cachePos, cache)
	} else if input[i][j] == '/' {
		newDir := Pos{-state.dir.j, -state.dir.i}
		cachePos, cache = lightPath(input, State{Pos{i, j}, newDir}, cachePos, cache)
	} else {
		newDir := Pos{state.dir.i, state.dir.j}
		cachePos, cache = lightPath(input, State{Pos{i, j}, newDir}, cachePos, cache)
	}
	return cachePos, cache
}

func energizer(input []string, init State) int {
	result := 0
	var cachePos [][]int
	for i, line := range input {
		cachePos = append(cachePos, make([]int, len(line)))
		for j := range line {
			cachePos[i][j] = 0
		}
	}
	cache := make(map[State]int)

	cachePos, _ = lightPath(input, init, cachePos, cache)

	for _, line := range cachePos {
		for _, col := range line {
			result += col
		}
	}
	return result
}

func part1(input []string) int {
	result := 0
	init := State{Pos{0, -1}, Pos{0, 1}}

	result += energizer(input, init)

	return result
}

func part2(input []string) int {
	result := 0
	bestResult := 0
	var bestState State
	var init State

	//loop de loop
	for i := range input {
		// Left edge
		init = State{Pos{i, -1}, Pos{0, 1}}
		result = energizer(input, init)
		if result > bestResult {
			bestResult = result
			bestState = init
		}
		// Right edge
		init = State{Pos{i, len(input[i])}, Pos{0, -1}}
		result = energizer(input, init)
		if result > bestResult {
			bestResult = result
			bestState = init
		}
	}
	for j := range input[0] {
		// Top edge
		init = State{Pos{j, -1}, Pos{1, 0}}
		result = energizer(input, init)
		if result > bestResult {
			bestResult = result
			bestState = init
		}
		// Right edge
		init = State{Pos{j, len(input)}, Pos{-1, 0}}
		result = energizer(input, init)
		if result > bestResult {
			bestResult = result
			bestState = init
		}
	}
	fmt.Println(bestState)
	return bestResult
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
