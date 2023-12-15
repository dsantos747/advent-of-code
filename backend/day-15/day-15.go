package main

import (
	"fmt"
	"strconv"
	"strings"

	tools "github.com/dsantos747/advent-of-code-2023/backend/tools"
)

type Lens struct {
	label string
	fl int
}

func HASH(input string) int {
	curr := 0
	for _,c := range input {
		curr += int(c)
		curr *= 17
		curr = curr % 256
	}
	return curr
}

func part1(input []string) int {
	result := 0
	for _,str := range input {
		result += HASH(str)
	}
	return result
}

func part2(input []string) int {
	result := 0
	var key int
	var lensMap = make(map[int][]Lens)
	for _,str := range input {
		if (strings.Contains(str,"=")) {
			label := strings.Split(str,"=")[0]
			key = HASH(label)
			fl,_ := strconv.Atoi(strings.Split(str,"=")[1])
			lensFound := false
				for i,lens := range lensMap[key] {
					if lens.label == label {
						lensMap[key][i].fl = fl
						lensFound = true
					}
				}
			if !lensFound {
				lensMap[key] = append(lensMap[key], Lens{label, fl})
			}
		} else {
			label := str[:len(str)-1]
			key = HASH(label)
			for i,lens := range lensMap[key] {
				if lens.label == label {
					lensMap[key] = append(lensMap[key][:i], lensMap[key][i+1:]...)
				}
			}
		}
	}
	for i,box := range lensMap {
		for j,lens := range box {
			result += (i+1)*(j+1)*(lens.fl)
		}
	}
	return result
}

func main() {
	data,err := tools.ReadInput("./input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input := strings.Split(data, ",")

	p1 := part1(input)
	fmt.Println("The answer to part 1 is",p1)

	p2 := part2(input)
	fmt.Println("The answer to part 2 is",p2)
}