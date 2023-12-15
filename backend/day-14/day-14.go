package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	tools "github.com/dsantos747/advent-of-code-2023/backend/tools"
)

func rollRocks(line string, left bool) string {
	matches := []int{}
	for i,char := range line {
		if char != '#' {
			matches = append(matches, i)
		}
	}
	a := matches[0]
	b := matches[len(matches)-1]
	
	n := strings.Split(line[a:b+1],"#")
	for i,l := range n {
		nl := strings.Split(l,"")
		if (left) {
			slices.SortFunc(nl, func(a, b string) int {
				return cmp.Compare(b, a)
			})
		} else {
			slices.Sort(nl)
		}

		n[i] = strings.Join(nl,"")
	}
	return line[:a] + strings.Join(n,"#") + line[b+1:]
}

func spinCycle(input []string) []string {
	// N
	input = tools.Transpose(input)
	for i,line := range input {
		input[i] = rollRocks(line,true)
	}
	// W
	input = tools.Transpose(input)
	for i,line := range input {
		input[i] = rollRocks(line,true)
	}
	// S
	input = tools.Transpose(input)
	for i,line := range input {
		input[i] = rollRocks(line,false)
	}
	// E
	input = tools.Transpose(input)
	for i,line := range input {
		input[i] = rollRocks(line,false)
	}
	return input
}

func weighRocks(input []string) int {
	result := 0
	for _,line := range input {
		oMatches := []int{}
		for i,char := range line {
			if char == 'O' {
				oMatches = append(oMatches, i)
			}
		}
		for _,oMatch := range oMatches {
			result += len(line) - oMatch
		}
	}
	return result
}

func part1(input []string) int {
	result := 0
	transp := tools.Transpose(input)
	for _,line := range transp {
		newLine := rollRocks(line,true)
		result += weighRocks([]string{newLine})
	}
	return result
}

func part2(input []string) int {
	cycles := 1000000000
	cache := make(map[string]string)
	var cacheEntry, answer string
	for _, line := range input {
		cacheEntry += line
	}

	loopStart := -1
	loopEntries := []string{}

	for cycle:=0;cycle<cycles;cycle++ {
		if v, ok := cache[cacheEntry]; ok {
			if loopStart == -1 {
				loopStart = cycle
			} else {
				if v == loopEntries[0] {
					break
				}
			}
			loopEntries = append(loopEntries, v)
			cacheEntry = v
			answer = v
			continue
		}



		input = spinCycle(input)
		answer = ""
		for _,line := range input {
			answer += line
		}
		// if cacheEntry == answer {
		// 	fmt.Println("match found")
		// 	break
		// }

		cache[cacheEntry] = answer
		cacheEntry = answer
	}

	c := (cycles - 1 - loopStart) % len(loopEntries)
	answer = loopEntries[c]
	start := 0
	for i := 0; i < len(input); i++ {
		input[i] = answer[start : start+len(input[0])]
		start += len(input[0])
	}

	field := tools.Transpose(input)

	return weighRocks(field)
}

func main() {
	data,err := tools.ReadInput("./input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input := strings.Split(data, "\n")
	
	// tStart := time.Now()
	p1 := part1(input)
	// t1 := time.Now()
	p2 := part2(input)
	// t2 := time.Now()
	// d1 := t1.Sub(tStart)
	// d2 := t2.Sub(t1)

	fmt.Println("The answer to part 1 is",p1)
	fmt.Println("The answer to part 2 is",p2)

	// fmt.Println("Part 1 computed in",d1)
	// fmt.Println("Part 2 computed in",d2)
}