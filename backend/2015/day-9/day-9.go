package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/dsantos747/advent-of-code/tools"
)

// ... how the heck...
// maybe look at https://github.com/coussej/adventofcode-solutions/blob/master/2015/day09/main.go
// Thanks again to reddit user u/coussej

type distances map[string]map[string]int

func makeDistanceMap(input []string) distances {
	distances := distances{}

	for _, line := range input {
		s := strings.Split(line, " ")
		dist, _ := strconv.Atoi(s[4])

		if _, ok := distances[s[0]]; !ok {
			distances[s[0]] = map[string]int{}
		}
		distances[s[0]][s[2]] = dist
	}

	return distances
}

func (d distances) findShortestRoute(visited map[string]bool, currPlace string, totalDist int) int {
	allVisited := true

	for _, isVisited := range visited {
		if !isVisited {
			allVisited = false
			break
		}
	}

	if allVisited {
		return totalDist
	}

	best := math.MaxInt

	for place, isVisited := range visited {
		if !isVisited {
			visited[place] = true

			distance := 0
			if currPlace == "" {
				distance = d.findShortestRoute(visited, place, 0)
			} else {
				distance = d.findShortestRoute(visited, place, totalDist+d[currPlace][place])
			}
			best = max(best, distance)
			visited[place] = false
		}
	}

	return best
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func part1(data []string) int {
	dists := makeDistanceMap(data)
	visited := map[string]bool{}
	for place := range dists {
		visited[place] = false
	}
	return dists.findShortestRoute(visited, "", 0)
}

func part2(data []string) int {
	result := 0

	return result
}

// func Solve(data string) (*int, *int, error) {
// 	input := strings.Split(data, "\n")

// 	p1 := part1(input)
// 	p2 := part2(input)

// 	return &p1, &p2, nil
// }

func main() {
	data, _ := tools.ReadInput("./input.txt")
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	fmt.Println(p1)
	fmt.Println(p2)
}
