package main

import (
	"fmt"
	"strconv"
	"strings"

	tools "github.com/dsantos747/advent-of-code-2023/tools"
)

type Pos struct {
	x int
	y int
	z int
}

type Line struct {
	x0 int
	y0 int
	z0 int
	dx int
	dy int
	dz int
	m  float64
	c  float64
}

func parseLines(input []string) []Line {
	lines := []Line{}
	var m float64
	var c float64

	for _, line := range input {
		l := []int{}
		s := strings.Split(strings.ReplaceAll(strings.ReplaceAll(line, "@", ","), " ", ""), ",")
		for _, num := range s {
			val, _ := strconv.Atoi(num)
			l = append(l, val)
		}
		m = float64(l[4]) / float64(l[3])
		c = float64(l[1]) - m*float64(l[0])

		lines = append(lines, Line{l[0], l[1], l[2], l[3], l[4], l[5], m, c})
	}
	return lines
}

func checkLines(l1, l2 Line) int {
	if l1.m == l2.m {
		return 0
	}

	var xIn, yIn, boundMin, boundMax, l1x0, l1y0, l2x0, l2y0, l1x1, l1y1, l2x1, l2y1 float64

	boundMin = 200000000000000
	boundMax = 400000000000000
	// boundMin = 7
	// boundMax = 27

	l1x0, l1x1, l1y0, l1y1 = getBounds(l1, boundMin, boundMax)
	l2x0, l2x1, l2y0, l2y1 = getBounds(l2, boundMin, boundMax)

	xIn = (l1.c - l2.c) / (l2.m - l1.m)
	yIn = (xIn * l1.m) + l1.c

	if xIn < l1x0 || xIn > l1x1 || yIn < l1y0 || yIn > l1y1 {
		return 0
	}
	if xIn < l2x0 || xIn > l2x1 || yIn < l2y0 || yIn > l2y1 {
		return 0
	}
	return 1
}

func getBounds(l Line, boundMin, boundMax float64) (float64, float64, float64, float64) {
	var lx0, ly0, lx1, ly1 float64

	if l.dx > 0 {
		if float64(l.x0) < boundMax {
			lx1 = boundMax
			lx0 = max(float64(l.x0), boundMin)
			// if float64(l.x0) < boundMin {
			// 	lx0 = boundMin
			// } else {
			// 	lx0 = float64(l.x0)
			// }
		} else {
			// fmt.Println(l.dx, l.x0)
			// fmt.Println("line out of bounds in x")
		}
	}
	if l.dx <= 0 {
		if float64(l.x0) > boundMin {
			lx0 = boundMin
			lx1 = min(float64(l.x0), boundMax)
			// if float64(l.x0) > boundMax {
			// 	lx1 = boundMax
			// } else {
			// 	lx1 = float64(l.x0)
			// }
		} else {
			// fmt.Println("line out of bounds in x")
		}
		if l.dx == 0 {
			fmt.Println("perfectly horizontal line")
		}
	}
	if l.dy > 0 {
		if float64(l.y0) < boundMax {
			ly1 = boundMax
			ly0 = max(float64(l.y0), boundMin)
			// if float64(l.y0) < boundMin {
			// 	ly0 = boundMin
			// } else {
			// 	ly0 = float64(l.y0)
			// }
		} else {
			// fmt.Println("line out of bounds in y")
		}
	}
	if l.dy <= 0 {
		if float64(l.y0) > boundMin {
			ly0 = boundMin
			ly1 = min(float64(l.y0), boundMax)
			// if float64(l.y0) > boundMax {
			// 	ly1 = boundMax
			// } else {
			// 	ly1 = float64(l.y0)
			// }
		} else {
			// fmt.Println("line out of bounds in y")
		}
		if l.dx == 0 {
			fmt.Println("perfectly vertical line")
		}
	}

	return lx0, lx1, ly0, ly1
}

func part1(input []string) int {
	lines := parseLines(input)
	total := 0

	for i, l1 := range lines {
		for j := i; j < len(lines); j++ {
			l2 := lines[j]
			total += checkLines(l1, l2)
		}
	}

	return total
}

func part2(input []string) int {

	return 0
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
