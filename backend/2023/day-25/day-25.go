package day25

import (
	"strings"
)

// Credit to u/rumkuhgel for this one

type Edge struct {
	from, to string
}

func createNodeMap(input []string) map[string][]string {
	nodeMap := map[string][]string{}

	for _, line := range input {
		split := strings.Split(line, ": ")
		node := split[0]
		if _, ok := nodeMap[node]; !ok {
			nodeMap[node] = []string{}
		}
		children := strings.Fields(split[1])
		for _, child := range children {
			nodeMap[node] = append(nodeMap[node], child)
			if _, ok := nodeMap[child]; !ok {
				nodeMap[child] = []string{}
			}
			nodeMap[child] = append(nodeMap[child], node)
		}
	}
	return nodeMap
}

func countEdges(nodes map[string][]string) map[Edge]int {
	found := map[Edge]int{}
	i := 0
	for from := range nodes {
		BFS(nodes, from, found)
		i++
		if i > 50 {
			break
		}
	}
	return found
}

func countNodes(nodes map[string][]string, start string) int {
	visitCache := map[string]bool{}
	q := []string{start}

	for len(q) > 0 {
		from := q[0]
		q = q[1:]

		for _, to := range nodes[from] {
			if _, ok := visitCache[to]; ok {
				continue
			}
			q = append(q, to)
			visitCache[to] = true
		}
	}

	return len(visitCache)
}

func BFS(nodes map[string][]string, start string, found map[Edge]int) {
	visited := map[string]bool{}
	q := []string{start}

	for len(q) > 0 {
		from := q[0]
		q = q[1:]

		for _, to := range nodes[from] {
			if _, ok := visited[to]; ok {
				continue
			}
			q = append(q, to)
			visited[to] = true
			var edge Edge
			if from < to {
				edge = Edge{from, to}
			} else {
				edge = Edge{to, from}
			}
			found[edge]++
		}
	}
}

func removeMax(edges map[Edge]int) Edge {
	max := 0
	var maxEdge Edge
	for edge, val := range edges {
		if val > max {
			max = val
			maxEdge = edge
		}
	}
	delete(edges, maxEdge)
	return maxEdge
}

func removeEdgeFromNodes(nodes map[string][]string, edge Edge) {
	new := []string{}
	for _, to := range nodes[edge.from] {
		if to != edge.to {
			new = append(new, to)
		}
	}
	nodes[edge.from] = new
	new = []string{}
	for _, from := range nodes[edge.to] {
		if from != edge.from {
			new = append(new, from)
		}
	}
	nodes[edge.to] = new
}

func part1(input []string) int {
	nodeMap := createNodeMap(input)

	edges := countEdges(nodeMap)
	first := removeMax(edges)
	removeEdgeFromNodes(nodeMap, first)

	edges = countEdges(nodeMap)
	second := removeMax(edges)
	removeEdgeFromNodes(nodeMap, second)

	edges = countEdges(nodeMap)
	third := removeMax(edges)
	removeEdgeFromNodes(nodeMap, third)

	group1 := countNodes(nodeMap, first.from)
	group2 := len(nodeMap) - group1

	return group1 * group2
}

func part2(input []string) int {
	return 0
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
