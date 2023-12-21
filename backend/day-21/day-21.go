package main

import (
	"fmt"
	"strings"

	tools "github.com/dsantos747/advent-of-code-2023/tools"
)

type Pos struct {
	i int
	j int
}

type State struct {
	i         int
	j         int
	stepCount int
}

// Big thanks to u/rumkuhgel and the general r/aoc crowd for part 2

func BFS(input []string, init State, targetSteps int) int {
	pathMap := make(map[State]bool)
	posMap := make(map[Pos]int)
	queue := []State{init}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if pathMap[curr] {
			continue
		}

		char := input[curr.i][curr.j]

		if char == '#' { // Is this necessary?
			continue
		}

		pathMap[curr] = true

		if curr.stepCount == targetSteps {
			posMap[Pos{curr.i, curr.j}] = curr.stepCount
			continue
		}

		for _, dir := range [4]Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := Pos{curr.i + dir.i, curr.j + dir.j}
			if isInField(input, next) && input[next.i][next.j] != '#' {
				queue = append(queue, State{next.i, next.j, curr.stepCount + 1})
			}

		}
	}

	return len(posMap)
}

func solveLagrange(x int, values []int) int {
	b0 := values[0]
	b1 := values[1] - values[0]
	b2 := values[2] - values[1]

	return b0 + (b1 * x) + (x*(x-1)/2)*(b2-b1)
}

func isInField(grid []string, pos Pos) bool {
	return pos.j >= 0 && pos.j < len(grid[0]) && pos.i >= 0 && pos.i < len(grid)
}

func stepBFS(input []string, visitable map[Pos]bool, part2 bool) map[Pos]bool {
	newPosMap := map[Pos]bool{}

	for position := range visitable {
		for _, dir := range [4]Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			nextLocation := Pos{position.i + dir.i, position.j + dir.j}
			if (part2 && input[tools.mod(nextLocation.i, len(input))][tools.mod(nextLocation.j, len(input[0]))] != '#') ||
				(!part2 && isInField(input, nextLocation) && input[nextLocation.i][nextLocation.j] != '#') {
				newPosMap[nextLocation] = true
			}
		}
	}
	return newPosMap
}

func oneStep(input []string, visitable map[Pos]bool, part2 bool) map[Pos]bool {
	newPosMap := map[Pos]bool{}

	for position := range visitable {
		for _, dir := range [4]Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			nextLocation := Pos{position.i + dir.i, position.j + dir.j}
			if (part2 && input[tools.mod(nextLocation.i, len(input))][tools.mod(nextLocation.j, len(input[0]))] != '#') ||
				(!part2 && isInField(input, nextLocation) && input[nextLocation.i][nextLocation.j] != '#') {
				newPosMap[nextLocation] = true
			}
		}
	}
	return newPosMap
}

func part1(input []string) int {
	i0, j0 := tools.findSingleSubstring(input, "S")
	// i0, j0 := findSingleSubstring(input, "S")
	initialState := State{i0, j0, 0} // Find initial i and j which correspond to "S"
	result := BFS(input, initialState, 64)
	return result
}

func part2(input []string) int {
	i0, j0 := tools.findSingleSubstring(input, "S")

	initialPos := Pos{i0, j0}

	posMap := map[Pos]bool{initialPos: true}
	yValues := [3]int{}
	for i := 1; i <= (len(input)/2)+2*len(input); i++ {
		posMap = stepBFS(input, posMap, true)
		if i == (len(input) / 2) {
			fmt.Println(i, len(posMap))
			yValues[0] = len(posMap)
		} else if i == (len(input)/2)+len(input) {
			fmt.Println(i, len(posMap))
			yValues[1] = len(posMap)
		} else if i == (len(input)/2)+2*len(input) {
			fmt.Println(i, len(posMap))
			yValues[2] = len(posMap)
		}
	}

	result2 := solveLagrange(26501365/131, yValues[:])

	return result2
}

func main() {
	data, err := tools.ReadInput("./input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input := strings.Split(data, "\n")

	p1 := part1(input)
	fmt.Println("The answer to part 1 is", p1)

	p2 := part2(input)
	fmt.Println("The answer to part 2 is", p2)
}
