package main

import (
	// "AOC23/tools"
	"fmt"
	"strings"

	tools "github.com/dsantos747/advent-of-code-2023/tools"
)

type Position struct {
	x int
	y int
}

func findStart(input []string) (int,int) {
	for i,line := range input {
		for j,char := range line {
			if char == 'S' {
				return i,j
			}
		}
	}
	return 0,0
}

func calcNext(x, y int, input []string, prevx, prevy int ) (int, int) {
	if (input[x][y] == '-') {
		if (prevy == y-1) {
			y++
		} else {
			y--
		}
		return x, y
	}
	if (input[x][y] == '|') {
		if (prevx == x-1) {
			x++
		} else {
			x--
		}
		return x, y
	}
	if (input[x][y] == 'J') {
		if (prevy == y-1) {
			x--
		}
		if (prevx == x-1) {
			y--
		}
		return x, y
	}
	if (input[x][y] == 'L') {
		if (prevy == y+1) {
			x--
		}
		if (prevx == x-1) {
			y++
		}
		return x, y
	}
	if (input[x][y] == 'F') {
		if (prevy == y+1) {
			x++
		}
		if (prevx == x+1) {
			y++
		}
		return x, y
	}
	if (input[x][y] == '7') {
		if (prevy == y-1) {
			x++
		}
		if (prevx == x+1) {
			y--
		}
		return x, y
	}
	return x,y
}

func findNextAfterStart(x int, y int, input []string) (int, int) {
	if (y < len(input[x])-1) && (input[x][y+1] == '-' || input[x][y+1] == 'J' ||input[x][y+1] == '7') {
		y++
		return x,y
	}
	if (y > 0) && (input[x][y-1] == '-' || input[x][y-1] == 'F' ||input[x][y-1] == 'L') {
		y--
		return x,y
	}
	if (x < len(input)-1) && (input[x+1][y] == '|' || input[x+1][y] == 'J' ||input[x+1][y] == 'L') {
		x++
		return x,y
	}
	if (x > 0) && (input[x-1][y] == '|' || input[x-1][y] == 'F' ||input[x-1][y] == '7') {
		x--
		return x,y
	}
	return x,y
}

func part1(data []string) (int,[]Position){
	pipeRoute := []Position{}

	x,y := findStart(data)
	pipeRoute = append(pipeRoute, Position{
		x:x,
		y:y,
	})

	x,y = findNextAfterStart(x,y,data)

	for data[x][y] != 'S' {
		prev := pipeRoute[len(pipeRoute)-1]
		pipeRoute = append(pipeRoute, Position{
			x: x,
			y: y,
		})
		x1,y1 := calcNext(x,y,data,prev.x,prev.y)
		x,y = x1,y1
	}
	return len(pipeRoute)/2 , pipeRoute
}

func convertS(s, first, lastChar Position, input []string) byte {
	for _, char := range "-|JL7F" {
		duplicateInput := input
		a := strings.Split(duplicateInput[s.x], "")
		a[s.y] = string(char)
		duplicateInput[s.x] = strings.Join(a, "")
		s1x, s1y := calcNext(s.x, s.y, duplicateInput, lastChar.x, lastChar.y)
		if (s1x == first.x) && (s1y == first.y) {
			fmt.Println(string(char))
			return byte(char)
		}
	}
	return 'S'
}

type NestCount struct {
	line string
	count int
}

func part2(input []string, path []Position) (int, map[int]NestCount) {
	subs := map[string]string{"J": "┘","L": "└","7":"┐","F":"┌","|":"│","-":"─"}

	routeMap := make(map[int][]int)
	resultMap := make(map[int]NestCount)
	sum := 0

	for _, p := range path {
		routeMap[p.x] = append(routeMap[p.x], p.y)
	}

	for k, v := range routeMap { // Using routeMap lets us skip lines which don't contain any path
		a := strings.Split(input[k],"")
		for _, j := range v {
			if a[j] != "S" {
				a[j] = subs[a[j]]
			} else {
				a[j] = subs[string(convertS(path[0], path[1], path[len(path)-1], input))]
			}
		}
		left,right := false, false
		for i,j := range a {
			if (i == len(a)-1-i) {
				break
			}
			r := a[len(a)-i-1]
			if j == "┘" || j == "└" || j == "┐" || j == "┌" || j == "│" || j == "─" {
				left = true
			}
			if r == "┘" || r == "└" || r == "┐" || r == "┌" || r == "│" || r == "─" {
				right = true
			}
			if !left {
				a[i] = " "
			}
			if !right {
				a[len(a)-1-i] = " "
			}
		}

		count := 0
		inside := 0
		lastChar := "-"
		for i,char := range a {
					if (char == "─") {
						continue
					}
					if (char == "│") {
						inside++
						continue
					}
					if lastChar == "-" && (char == "┘" || char == "└" || char == "┐" || char == "┌") {
							lastChar = char
							continue
					} else if lastChar != "-" && (char == "┘" || char == "└" || char == "┐" || char == "┌") {
						if lastChar == "└" && char == "┐" {
							inside++
						}
						if (lastChar == "┌") && (char == "┘") {
							inside++
						}
						lastChar = "-"
						continue
					}
					if (inside % 2 == 0) {
						a[i] = " "
					}
					if (inside % 2 != 0) {
						a[i] = "\033[0;32m█\033[0m"
						count++
					}
		}
		resultMap[k] = NestCount{
			line: strings.Join(a,""),
			count: count,
		}
		sum += count
	}
	return sum, resultMap
}

func main() {
	input,err := tools.ReadInput("./input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	data := strings.Split(input, "\n")

	p1,pipeRoute := part1(data)
	fmt.Println("The answer to part 1 is",p1)

	p2, pipeMap := part2(data,pipeRoute)
	for i := 0; i < len(input); i++ {
		if _, ok := pipeMap[i]; ok {
			fmt.Println(pipeMap[i].line, pipeMap[i].count)
		}
	}
	fmt.Println("The answer to part 2 is",p2)
}