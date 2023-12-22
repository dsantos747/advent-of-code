package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	tools "github.com/dsantos747/advent-of-code-2023/tools"
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
	size   Pos
	pos    Pos
	valid  bool
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

func dropBricks2(brickSlice [][]int) map[int]Brick {
	brickMap := make(map[int]Brick)
	var init0 Pos
	var init1 Pos
	var pos Pos

	for i, brick := range brickSlice {
		init0 = Pos{brick[0], brick[1], brick[2]}
		init1 = Pos{brick[3], brick[4], brick[5]}
		maxZ := 1
		for _, under := range brickSlice[:i] {
			if overlap2(brick, under) {
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

			if overlap2(brick, brickSlice[j]) {
				heldBy = append(heldBy, j)
			}
		}
		newBrick.heldBy = heldBy
		brickMap[i] = newBrick
		// fmt.Println(brickMap[i])
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

func dropBricks(brickMap map[int]Brick) map[int]Brick {
	var terrain [][]int
	for i := 0; i < 10; i++ {
		terrain = append(terrain, make([]int, 10))
		for j := 0; j < 10; j++ {
			terrain[i][j] = 0
		}
	}
	for i := 0; i < len(brickMap); i++ {
		curr := brickMap[i]
		var newBrick Brick
		var newBrickPos Pos
		var heldBy []int
		xMaxZ := 0
		yMaxZ := 0

		/////////////// Warning loop for testing
		if (curr.size.x > 1 && curr.size.y > 1) || (curr.size.z > 1 && curr.size.y > 1) || (curr.size.x > 1 && curr.size.z > 1) {
			fmt.Println("Big boi block alert!:", curr.size.x, curr.size.y, curr.size.z)
		}

		if curr.size.x > 1 {
			for ix := curr.init0.x; ix <= curr.init1.x; ix++ {
				if terrain[ix][curr.init0.y] > xMaxZ {
					xMaxZ = terrain[ix][curr.init0.y]
				}
			}
			for ix := curr.init0.x; ix <= curr.init1.x; ix++ {
				terrain[ix][curr.init0.y] = xMaxZ + curr.size.z
				newBrickPos = Pos{curr.init0.x, curr.init0.y, xMaxZ + 1}
			}
		} else if curr.size.y > 1 {
			for iy := curr.init0.y; iy <= curr.init1.y; iy++ {
				if terrain[curr.init0.x][iy] > yMaxZ {
					yMaxZ = terrain[curr.init0.x][iy]
				}
			}
			for iy := curr.init0.y; iy <= curr.init1.y; iy++ {
				terrain[curr.init0.x][iy] = yMaxZ + curr.size.z
				newBrickPos = Pos{curr.init0.x, curr.init0.y, yMaxZ + 1}
			}
		} else {
			newBrickPos = Pos{curr.init0.x, curr.init0.y, terrain[curr.init0.x][curr.init0.y] + 1}
			terrain[curr.init0.x][curr.init0.y] += curr.size.z
		}
		// fmt.Println(xMaxZ)
		// fmt.Println(curr.size.z)

		for j := i - 1; j >= 0; j-- {
			if newBrickPos.z-brickMap[j].pos.z != 1 {
				continue
			}

			if overlap(curr, brickMap[j]) {
				heldBy = append(heldBy, j)
			}
		}

		// fmt.Println("Brick", i, "held by", heldByCount)

		newBrick = Brick{id: i, init0: curr.init0, init1: curr.init1, size: curr.size, pos: newBrickPos, valid: false, heldBy: heldBy, holds: []int{}}

		brickMap[i] = newBrick
		// fmt.Println(newBrick)
		// for i := 0; i < 10; i++ {
		// 	fmt.Println(terrain[i])
		// }
		// fmt.Println()
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
	for i := 0; i < len(brickMap); i++ {
		curr := brickMap[i]
		if curr.size.x < 1 || curr.size.y < 1 || curr.size.z < 1 {
			fmt.Println("brick", i, "too small")
		}
		// fmt.Println(brickMap[i])
	}

	// for i := 0; i < len(brickMap); i++ {
	// 	curr := brickMap[i]
	// 	if hasDuplicates(curr.heldBy) {
	// 		fmt.Println("duplicates found!")
	// 	}
	// }

	return brickMap
}

func hasDuplicates(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true // Duplicate found
		}
		seen[num] = true
	}

	return false // No duplicates found
}

func rangeOverlap(start1, end1, start2, end2 int) bool {
	return end1 >= start2 && end2 >= start1
}

func overlap(over, under Brick) bool {
	return max(over.init0.x, under.init0.x) <= min(over.init1.x, under.init1.x) && max(over.init0.y, under.init0.y) <= min(over.init1.y, under.init1.y)
}

func overlap2(over, under []int) bool {
	return max(over[0], under[0]) <= min(over[3], under[3]) && max(over[1], under[1]) <= min(over[4], under[4])
}

func validBricks(brickMap map[int]Brick) int {
	dontRemoveMap := make(map[int]bool)

	for i := 0; i < len(brickMap); i++ {
		// var newBrick Brick
		curr := brickMap[i]
		// curr.valid = true
		// if len(curr.holds) == 0 {
		// 	brickMap[i] = curr
		// 	continue
		// }
		if len(curr.heldBy) == 1 {
			sup := brickMap[curr.heldBy[0]]
			if _, ok := dontRemoveMap[sup.id]; !ok {
				dontRemoveMap[sup.id] = true
			}
		}
		// for j := 0; j < len(curr.holds); j++ {
		// 	jBrick := brickMap[curr.holds[j]]
		// 	if len(jBrick.heldBy) < 2 {
		// 		curr.valid = false
		// 		break
		// 	}
		// }
		// brickMap[i] = curr
	}

	return len(brickMap) - len(dontRemoveMap)
}

func dummycheck(brickMap map[int]Brick) int {
	count := 0
	for i := 0; i < len(brickMap); i++ {
		curr := brickMap[i]
		if curr.size.x == 1 && curr.size.y == 1 && curr.size.z == 1 {
			count++
		}
	}
	return count
}

func chainReact(brickMap map[int]Brick, brick int, cache map[int]int) (int, map[int]int) {

	var count int
	var nextFall int

	if len(brickMap[brick].holds) == 0 {
		return 1, cache
	}

	for _, next := range brickMap[brick].holds {
		if val, ok := cache[next]; ok {
			count += val
			continue
		}

		nextFall, cache = chainReact(brickMap, next, cache)
		cache[next] = count
		count += nextFall

	}

	cache[brick] = count
	return count, cache
}

func chainReactBFS(brickMap map[int]Brick, brick int, cache map[int]int) int {

	queue := []int{brick}

	for j, _ := range brickMap[brick].holds {
		if len(brickMap[j].heldBy) == 1 {
			queue = append(queue, brickMap[j].id)
		}
	}

	toppledCache := make(map[int]bool)
	for _, j := range queue {
		toppledCache[j] = true
	}
	toppledCache[brick] = true

	for len(queue) > 0 {
		b := queue[0]
		queue = queue[1:]
		// curr := brickMap[b]

		topple := true
		for j, _ := range brickMap[b].holds {
			if _, ok := toppledCache[j]; !ok {
				for _, brick := range brickMap[j].heldBy {
					if _, ok := toppledCache[brick]; !ok {
						topple = false
					}
				}
				if topple {
					queue = append(queue, j)
					toppledCache[j] = true
				}

			}
		}

		// continue

		// queue = append(queue, curr.holds...)

		// if len(curr.heldBy) == 0 {
		// 	toppledCache[b] = true
		// 	// queue = append(queue, curr.holds...)
		// 	continue
		// }

		// // if len(curr.holds) == 1 {
		// // 	queue = append(queue, curr.holds[0])
		// // }
		// topple := true
		// for _, brick := range curr.heldBy {
		// 	if _, ok := toppledCache[brick]; !ok {
		// 		topple = false
		// 		break
		// 	}
		// }
		// if topple {
		// 	toppledCache[b] = true
		// 	// queue = append(queue, curr.holds...)
		// }

		// for brick in currentbrick.heldby, if brick IS in toppledCache, add brick to queue (and to toppledCache)

	}
	return len(toppledCache) - 1
}

func part1(input []string) (int, map[int]Brick) {

	brickSlice := makeBrickSlice(input)
	brickMap := dropBricks2(brickSlice)
	return validBricks(brickMap), brickMap

}

func part2(brickMap map[int]Brick) int {
	total := 0
	res := 0
	cache := make(map[int]int)

	for i := 0; i < len(brickMap); i++ {
		// fmt.Println(i, res, "     ", brickMap[i])
		res = chainReactBFS(brickMap, brickMap[i].id, cache)
		total += res
		// fmt.Println(res)
		// if i == 10 {
		// 	break
		// }
	}

	return total
}

func main() {
	data, err := tools.ReadInput("./input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input := strings.Split(data, "\n")

	p1, brickMap := part1(input)
	fmt.Println("The answer to part 1 is", p1)

	p2 := part2(brickMap)
	fmt.Println("The answer to part 2 is", p2)
}
