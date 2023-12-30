package day23

import (
	"strings"

	tools "github.com/dsantos747/advent-of-code-2023/tools"
)

type ij struct { // ROWS BY COLUMNS
	i int
	j int
}

var (
	left  = ij{0, -1}
	right = ij{0, 1}
	up    = ij{-1, 0}
	down  = ij{1, 0}
)

var adj = map[ij][]ij{
	left:  {up, down},
	right: {up, down},
	up:    {left, right},
	down:  {left, right},
}

type Path struct {
	id     int
	to     ij
	length int
}

func DFS(input []string, pos, dir ij, visitCache map[ij]int, dry bool) {
	curr := pos
	currStepCount := visitCache[curr]

	for _, dir := range []ij{dir, adj[dir][0], adj[dir][1]} {
		next := ij{curr.i + dir.i, curr.j + dir.j}

		if next.i < 0 || next.i >= len(input) || next.j < 0 || next.j >= len(input[0]) {
			continue
		}

		if input[next.i][next.j] == '#' {
			continue
		}

		if !dry {
			if input[next.i][next.j] == '>' && dir == left {
				continue
			} else if input[next.i][next.j] == 'v' && dir == up {
				continue
			} else if input[next.i][next.j] == '<' && dir == right {
				continue
			} else if input[next.i][next.j] == '^' && dir == down {
				continue
			}
		}

		if val, ok := visitCache[next]; !ok || val < currStepCount+1 {
			visitCache[next] = currStepCount + 1
			DFS(input, next, dir, visitCache, dry)
		}

	}
}

func getIntersects(input []string) map[ij]bool {
	junctions := map[ij]bool{}
	for i, line := range input {
		for j, char := range line {
			if char == '#' {
				continue
			}
			neighbours := 0
			for _, dir := range []ij{left, up, right, down} {
				nI, nJ := i+dir.i, j+dir.j
				if tools.InField(input, nI, nJ) && input[nI][nJ] != '#' {
					neighbours++
				}
			}
			if neighbours > 2 {
				junctions[ij{i, j}] = true
			}
		}
	}
	return junctions
}

func getPaths(input []string, junctions map[ij]bool, start, goal ij) map[ij][]Path {
	paths := map[ij][]Path{}
	juncID := 0

	for junc := range junctions {
		blocked := -1
		for d, dir := range []ij{left, up, right, down} {
			nI, nJ := junc.i+dir.i, junc.j+dir.j
			if tools.InField(input, nI, nJ) && input[nI][nJ] != '#' {
				path := getPath(input, junc, ij{nI, nJ}, dir, 1, junctions)
				path.id = juncID
				paths[junc] = append(paths[junc], path)
			} else {
				blocked = d
			}
		}
		if blocked != -1 && junc != start && junc != goal {
			x := 0
			if blocked == 2 {
				x = 1
			}
			paths[junc] = append(paths[junc][:x], paths[junc][x+1:]...)
		}
		juncID++
	}

	return paths
}

func getPath(input []string, start, currPos, currDir ij, length int, junctions map[ij]bool) Path {
	for _, dir := range []ij{currDir, adj[currDir][0], adj[currDir][1]} {
		nI, nJ := currPos.i+dir.i, currPos.j+dir.j
		if input[nI][nJ] != '#' {
			if _, ok := junctions[ij{nI, nJ}]; ok {
				return Path{0, ij{nI, nJ}, length + 1}
			}
			return getPath(input, start, ij{nI, nJ}, dir, length+1, junctions)
		}
	}
	return Path{0, ij{-1, -1}, 0}
}

func longestPath(input []string, paths map[ij][]Path, start, goal ij, length int, visitCache []bool) int {
	maxLength := 0
	for _, path := range paths[start] {
		if path.to == goal {
			return length + path.length
		}
		id := paths[path.to][0].id
		if !visitCache[id] {
			visitCache[id] = true
			maxLength = max(maxLength, longestPath(input, paths, path.to, goal, length+path.length, visitCache))
			visitCache[id] = false
		}
	}
	return maxLength
}

func part1(input []string) int {
	start := ij{0, 1}
	goal := ij{len(input) - 1, len(input[0]) - 2}
	visitCache := map[ij]int{start: 0}

	DFS(input, start, down, visitCache, false)

	return visitCache[goal]
}

func part2(input []string) int {
	start := ij{0, 1}
	goal := ij{len(input) - 1, len(input[0]) - 2}

	junctions := getIntersects(input)
	junctions[start] = true
	junctions[goal] = true

	paths := getPaths(input, junctions, start, goal)

	visitCache := make([]bool, len(junctions))
	visitCache[paths[start][0].id] = true

	return longestPath(input, paths, start, goal, 0, visitCache)
}

func Solve(data string) (int, int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return p1, p2, nil
}
