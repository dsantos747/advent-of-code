package day21

import (
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

func part1(input []string) int {
	i0, j0 := tools.FindSingleSubstring(input, "S")
	init := State{i0, j0, 0}

	return BFS(input, init, 64)
}

func part2(input []string) int {
	l := len(input)
	i0, j0 := tools.FindSingleSubstring(input, "S")
	init := Pos{i0, j0}
	var q = []Pos{init}

	yVals := make([]int, ((l/2)+2*l)+1)
	posMapOdd := map[Pos]bool{}
	posMapEven := map[Pos]bool{init: true}

	for i := 1; i < len(yVals); i++ {
		var posMap *map[Pos]bool
		if i%2 == 0 {
			posMap = &posMapEven
		} else {
			posMap = &posMapOdd
		}

		newQ := []Pos{}
		for i := 0; i < len(q); i++ {
			curr := q[i]
			for _, dir := range [4]Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				next := Pos{curr.i + dir.i, curr.j + dir.j}
				if input[tools.Mod(next.i, len(input))][tools.Mod(next.j, len(input[0]))] != '#' {
					if _, ok := (*posMap)[next]; !ok {
						(*posMap)[next] = true
						newQ = append(newQ, next)
					}
				}
			}
		}
		q = newQ
		yVals[i] = len(*posMap)
	}
	return solveLagrange(26501365/l, []int{yVals[l/2], yVals[(l/2)+l], yVals[(l/2)+2*l]})
}

func Solve(data string) (int, int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return p1, p2, nil
}
