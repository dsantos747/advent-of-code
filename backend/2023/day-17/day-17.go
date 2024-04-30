package day17

import (
	"container/heap"
	"math"
	"slices"
	"strconv"
	"strings"
)

// Big thanks to u/Multipl in the AoC reddit

type Dir struct { // ROWS BY COLUMNS
	i int
	j int
}

var (
	left  = Dir{0, -1}
	right = Dir{0, 1}
	up    = Dir{-1, 0}
	down  = Dir{1, 0}
)

var adj = map[Dir][]Dir{
	left:  {up, down},
	right: {up, down},
	up:    {left, right},
	down:  {left, right},
}

var opp = map[Dir]Dir{
	left:  right,
	right: left,
	up:    down,
	down:  up,
}

type Step struct {
	i    int
	j    int
	dir  Dir
	hist int
}

type Node struct {
	hl    int
	state Step
	index int
}

// ///// PriorityQueue implementation courtesy of Go team (https://pkg.go.dev/container/heap) /////
// A PriorityQueue implements heap.Interface and holds Nodes.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].hl < pq[j].hl
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	Node := x.(*Node)
	Node.index = n
	*pq = append(*pq, Node)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	Node := old[n-1]
	old[n-1] = nil  // avoid memory leak
	Node.index = -1 // for safety
	*pq = old[0 : n-1]
	return Node
}

// update modifies the priority and value of an Node in the queue.
// func (pq *PriorityQueue) update(Node *Node, value string, priority int) {
// 	// Node.value = value
// 	Node.hl = priority
// 	heap.Fix(pq, Node.index)
// }

// /////

func BFSwithPQ(input []string, maxStraight, minTurn int) int {
	bestSum := math.MaxInt // Predefine all paths as inf
	step0 := Step{0, 0, right, 0}
	step1 := Step{0, 0, down, 0}

	minHeatMap := map[Step]int{step0: 0, step1: 0} // Define intial starting points as 0

	pq := PriorityQueue{&Node{0, step0, 0}, &Node{0, step1, 1}} // Add first two steps to queue
	heap.Init(&pq)                                              // Initialise queue

	for len(pq) > 0 { // Until queue is empty
		this := heap.Pop(&pq).(*Node)
		if minHeatMap[this.state] < this.hl {
			continue
		}

		// If you've reached the end and the heatloss is less than previously achieved, save it
		if this.state.i == len(input)-1 && this.state.j == len(input[0])-1 && this.hl < bestSum && this.state.hist >= minTurn {
			bestSum = this.hl
		}

		for _, dir := range []Dir{right, down, left, up} {
			// If (we need to turn AND dir is not in available turn dirs for current step dir) OR dir is reverse of current step dir - continue
			if (this.state.hist == maxStraight && !slices.Contains(adj[this.state.dir], dir)) || dir == opp[this.state.dir] {
				continue
			}

			nextI, nextJ := this.state.i+dir.i, this.state.j+dir.j
			nextHist := this.state.hist

			if this.state.hist < minTurn {
				if dir != this.state.dir { // If we must go straight, but dir isn't straight - continue
					continue
				}
				nextHist += 1
			} else {
				if dir != this.state.dir {
					nextHist = 1 // Change direction, so reset history counter
				} else {
					// nextHist = nextHist%maxStraight + 1
					nextHist += 1
				}
			}

			if nextI < 0 || nextI >= len(input) || nextJ < 0 || nextJ >= len(input[0]) { // If next is out of bounds, continue
				continue
			}

			nextStep := Step{nextI, nextJ, dir, nextHist}
			nextHl, _ := strconv.Atoi(string(input[nextI][nextJ]))

			if _, ok := minHeatMap[nextStep]; ok {
				if minHeatMap[nextStep] <= this.hl+nextHl {
					continue
				}
			}

			minHeatMap[nextStep] = this.hl + nextHl

			heap.Push(&pq, &Node{hl: this.hl + nextHl, state: nextStep})

		}
	}

	return bestSum
}

func part1(input []string) int {
	bestSum := BFSwithPQ(input, 3, 0)
	return bestSum
}

func part2(input []string) int {
	bestSum := BFSwithPQ(input, 10, 4)
	return bestSum
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
