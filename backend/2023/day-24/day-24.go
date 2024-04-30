package day24

import (
	"slices"
	"strconv"
	"strings"
)

type Hailstone struct {
	x0 float64
	y0 float64
	z0 float64
	dx float64
	dy float64
	dz float64
	m  float64
	c  float64
}

func parseLines(input []string) []Hailstone {
	lines := []Hailstone{}
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

		lines = append(lines, Hailstone{float64(l[0]), float64(l[1]), float64(l[2]), float64(l[3]), float64(l[4]), float64(l[5]), m, c})
	}
	return lines
}

func checkLines(l1, l2 Hailstone) int {
	if l1.m == l2.m {
		return 0
	}

	var xIn, yIn, boundMin, boundMax, l1x0, l1y0, l2x0, l2y0, l1x1, l1y1, l2x1, l2y1 float64

	boundMin, boundMax = 200000000000000, 400000000000000
	// boundMin,boundMax = 7,27

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

func getBounds(l Hailstone, boundMin, boundMax float64) (float64, float64, float64, float64) {
	var lx0, ly0, lx1, ly1 float64

	if l.dx > 0 && float64(l.x0) < boundMax {
		lx1 = boundMax
		lx0 = max(float64(l.x0), boundMin)

	}
	if l.dx <= 0 && float64(l.x0) > boundMin {
		lx0 = boundMin
		lx1 = min(float64(l.x0), boundMax)

	}
	if l.dy > 0 && float64(l.y0) < boundMax {
		ly1 = boundMax
		ly0 = max(float64(l.y0), boundMin)

	}
	if l.dy <= 0 && float64(l.y0) > boundMin {
		ly0 = boundMin
		ly1 = min(float64(l.y0), boundMax)
	}

	return lx0, lx1, ly0, ly1
}

func getIntersect3D(x, y []int) []int {
	res := []int{}
	for _, val := range x {
		if slices.Contains(y, val) {
			res = append(res, val)
		}
	}
	return res
}

func matchVelocity(dv, pv int) []int {
	res := []int{}
	for v := -1000; v <= 1000; v++ {
		if v != pv && dv%(v-pv) == 0 {
			res = append(res, v)
		}
	}
	return res
}

func part1(input []string) (int, []Hailstone) {
	lines := parseLines(input)
	total := 0

	for i, l1 := range lines {
		for j := i; j < len(lines); j++ {
			l2 := lines[j]
			total += checkLines(l1, l2)
		}
	}

	return total, lines
}

func part2(hailstones []Hailstone) int {
	maybeX, maybeY, maybeZ := []int{}, []int{}, []int{}

	for i := 0; i < len(hailstones)-1; i++ {
		for j := i + 1; j < len(hailstones); j++ {
			a, b := hailstones[i], hailstones[j]
			if a.dx == b.dx {
				next := matchVelocity(int(b.x0-a.x0), int(a.dx))
				if len(maybeX) == 0 {
					maybeX = next
				} else {
					maybeX = getIntersect3D(maybeX, next)
				}
			}
			if a.dy == b.dy {
				next := matchVelocity(int(b.y0-a.y0), int(a.dy))
				if len(maybeY) == 0 {
					maybeY = next
				} else {
					maybeY = getIntersect3D(maybeY, next)
				}
			}
			if a.dz == b.dz {
				next := matchVelocity(int(b.z0-a.z0), int(a.dz))
				if len(maybeZ) == 0 {
					maybeZ = next
				} else {
					maybeZ = getIntersect3D(maybeZ, next)
				}
			}
		}
	}

	result := 0
	if len(maybeX) == 1 && len(maybeY) == 1 && len(maybeZ) == 1 {
		rdx, rdy, rdz := float64(maybeX[0]), float64(maybeY[0]), float64(maybeZ[0])
		hs1, hs2 := hailstones[0], hailstones[1]
		m1 := (hs1.dy - rdy) / (hs1.dx - rdx)
		m2 := (hs2.dy - rdy) / (hs2.dx - rdx)
		c1 := hs1.y0 - (m1 * hs1.x0)
		c2 := hs2.y0 - (m2 * hs2.x0)
		rx := (c2 - c1) / (m1 - m2)
		ry := m1*rx + c1
		t := (rx - hs1.x0) / (hs1.dx - rdx)
		rz := hs1.z0 + (hs1.dz-rdz)*t
		result = int(rx + ry + rz)
	}

	return result
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1, hailstones := part1(input)
	p2 := part2(hailstones)

	return &p1, &p2, nil
}
