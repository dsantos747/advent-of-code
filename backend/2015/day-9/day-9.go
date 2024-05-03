package day9

import (
	"math"
	"strconv"
	"strings"
)

// Thanks to github.com/teivah - needed some help from him to get me started

type Routes struct {
	distances map[string]map[string]int
}

func makeDistanceMap(input []string) Routes {
	r := Routes{distances: make(map[string]map[string]int)}

	for _, line := range input {
		s := strings.Split(line, " ")
		from, to := s[0], s[2]
		dist, _ := strconv.Atoi(s[4])

		if _, ok := r.distances[from]; !ok {
			r.distances[from] = map[string]int{}
		}
		r.distances[from][to] = dist

		if _, ok := r.distances[to]; !ok {
			r.distances[to] = map[string]int{}
		}
		r.distances[to][from] = dist
	}

	return r
}

func (r *Routes) findShortestRoute(visited map[string]bool, currPlace string, totalDist int) int {
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
				distance = r.findShortestRoute(visited, place, 0)
			} else {
				distance = r.findShortestRoute(visited, place, totalDist+r.distances[currPlace][place])
			}
			best = min(best, distance)
			visited[place] = false
		}
	}

	return best
}

func (r *Routes) findLongestRoute(visited map[string]bool, currPlace string, totalDist int) int {
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

	best := -1

	for place, isVisited := range visited {
		if !isVisited {
			visited[place] = true

			distance := 0
			if currPlace == "" {
				distance = r.findLongestRoute(visited, place, 0)
			} else {
				distance = r.findLongestRoute(visited, place, totalDist+r.distances[currPlace][place])
			}
			best = max(best, distance)
			visited[place] = false
		}
	}

	return best
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func part1(data []string) int {
	r := makeDistanceMap(data)
	visited := map[string]bool{}
	for place := range r.distances {
		visited[place] = false
	}

	return r.findShortestRoute(visited, "", 0)
}

func part2(data []string) int {
	r := makeDistanceMap(data)
	visited := map[string]bool{}
	for place := range r.distances {
		visited[place] = false
	}

	return r.findLongestRoute(visited, "", 0)
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
