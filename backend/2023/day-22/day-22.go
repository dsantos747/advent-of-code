package day22

import (
	"sort"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
	z int
}

type Brick struct {
	id     int
	init0  Pos
	init1  Pos
	pos    Pos
	heldBy []int
	holds  []int
}

func makeBrickSlice(input []string) [][]int {
	var brickSlice [][]int

	for _, line := range input {
		ls := append(strings.Split(strings.Split(line, "~")[0], ","), strings.Split(strings.Split(line, "~")[1], ",")...)
		var lineSlice []int
		for _, char := range ls {
			val, _ := strconv.Atoi(char)
			lineSlice = append(lineSlice, val)
		}
		brickSlice = append(brickSlice, lineSlice)
	}

	sort.Slice(brickSlice, func(i, j int) bool {
		return brickSlice[i][2] < brickSlice[j][2]
	})

	return brickSlice
}

func dropBricks(brickSlice [][]int) map[int]Brick {
	brickMap := make(map[int]Brick)
	var init0 Pos
	var init1 Pos
	var pos Pos

	for i, brick := range brickSlice {
		init0 = Pos{brick[0], brick[1], brick[2]}
		init1 = Pos{brick[3], brick[4], brick[5]}
		maxZ := 1
		for _, under := range brickSlice[:i] {
			if overlap(brick, under) {
				maxZ = max(maxZ, under[5]+1)
			}
		}
		brick[5] -= brick[2] - maxZ
		brick[2] = maxZ

		pos = Pos{brick[0], brick[1], brick[2]}
		brickMap[i] = Brick{id: i, init0: init0, init1: init1, pos: pos}
	}
	sort.Slice(brickSlice, func(i, j int) bool {
		return brickSlice[i][2] < brickSlice[j][2]
	})

	for i, brick := range brickSlice {
		var heldBy []int

		newBrick := brickMap[i]
		for j := i - 1; j >= 0; j-- {
			if brick[2]-brickSlice[j][5] != 1 {
				continue
			}

			if overlap(brick, brickSlice[j]) {
				heldBy = append(heldBy, j)
			}
		}
		newBrick.heldBy = heldBy
		brickMap[i] = newBrick
	}

	for i := 0; i < len(brickMap); i++ {
		var newBrick Brick
		curr := brickMap[i]
		if len(curr.heldBy) == 0 {
			continue
		}
		for _, j := range curr.heldBy {
			newBrick = brickMap[j]
			newBrick.holds = append(newBrick.holds, i)
			brickMap[j] = newBrick
		}
	}

	return brickMap
}

func overlap(over, under []int) bool {
	return max(over[0], under[0]) <= min(over[3], under[3]) && max(over[1], under[1]) <= min(over[4], under[4])
}

func validBricks(brickMap map[int]Brick) int {
	dontRemoveMap := make(map[int]bool)
	for i := 0; i < len(brickMap); i++ {
		curr := brickMap[i]
		if len(curr.heldBy) == 1 {
			sup := brickMap[curr.heldBy[0]]
			if _, ok := dontRemoveMap[sup.id]; !ok {
				dontRemoveMap[sup.id] = true
			}
		}
	}
	return len(brickMap) - len(dontRemoveMap)
}

func chainReactBFS(brickMap map[int]Brick, brick int, cache map[int]int) int {
	queue := []int{brick}
	toppledCache := make(map[int]bool)
	toppledCache[brick] = true

	for _, j := range brickMap[brick].holds {
		if len(brickMap[j].heldBy) == 1 {
			queue = append(queue, j)
		}
	}
	for _, j := range queue {
		toppledCache[j] = true
	}

	for len(queue) > 0 {
		b := queue[0]
		queue = queue[1:]
		for _, j := range brickMap[b].holds {
			if _, ok := toppledCache[j]; !ok {
				if isSublist(brickMap[j].heldBy, toppledCache) {
					queue = append(queue, j)
					toppledCache[j] = true
				}
			}
		}
	}
	return len(toppledCache) - 1
}

func isSublist(a []int, b map[int]bool) bool {
	for _, i := range a {
		if _, ok := b[i]; !ok {
			return false
		}
	}
	return true
}

func part1(input []string) (int, map[int]Brick) {
	brickSlice := makeBrickSlice(input)
	brickMap := dropBricks(brickSlice)
	return validBricks(brickMap), brickMap
}

func part2(brickMap map[int]Brick) int {
	total := 0
	res := 0
	cache := make(map[int]int)

	for i := range brickMap {
		res = chainReactBFS(brickMap, brickMap[i].id, cache)
		total += res
	}
	return total
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1, brickMap := part1(input)
	p2 := part2(brickMap)

	return &p1, &p2, nil
}
