package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	tools "github.com/dsantos747/advent-of-code-2023/tools"
)

type Pos struct { // ROWS BY COLUMNS
	i int
	j int
}

func part1(input []string) int {
	result := 0
	perim := 0
	area := 0

	// Use area of polygon from coordinates - https://www.themathdoctors.org/polygon-coordinates-and-areas/
	var coords []Pos
	c := Pos{0, 0}
	coords = append(coords, c)

	for _, line := range input {
		dir := line[0]
		dist, _ := strconv.Atoi(strings.Fields(line)[1])

		if dir == 'R' {
			c = Pos{c.i, c.j + dist}
		} else if dir == 'L' {
			c = Pos{c.i, c.j - dist}
		} else if dir == 'D' {
			c = Pos{c.i + dist, c.j}
		} else if dir == 'U' {
			c = Pos{c.i - dist, c.j}
		}
		coords = append(coords, c)

		perim += dist
	}

	perim = (perim / 2) + 1

	for i := 0; i < len(coords)-1; i++ {
		area += ((coords[i].j * coords[i+1].i) - (coords[i+1].j * coords[i].i))
	}

	area = int(math.Abs(float64(area)) / 2)

	result = perim + area

	return result
}

func part2(input []string) int {
	result := 0
	perim := 0
	area := 0

	// Use area of polygon from coordinates - https://www.themathdoctors.org/polygon-coordinates-and-areas/
	var coords []Pos
	c := Pos{0, 0}
	coords = append(coords, c)

	for _, line := range input {
		code := strings.ReplaceAll(strings.ReplaceAll(strings.Fields(line)[2], "(#", ""), ")", "")
		dir := code[len(code)-1]
		dist := tools.HexToInt(code[:len(code)-1])

		if dir == '0' {
			c = Pos{c.i, c.j + dist}
		} else if dir == '2' {
			c = Pos{c.i, c.j - dist}
		} else if dir == '1' {
			c = Pos{c.i + dist, c.j}
		} else if dir == '3' {
			c = Pos{c.i - dist, c.j}
		}
		coords = append(coords, c)

		perim += dist
	}

	perim = (perim / 2) + 1

	for i := 0; i < len(coords)-1; i++ {
		area += ((coords[i].j * coords[i+1].i) - (coords[i+1].j * coords[i].i))
	}

	area = int(math.Abs(float64(area)) / 2)

	result = perim + area

	return result
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
