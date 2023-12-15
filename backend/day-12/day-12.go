package main

import (
	// "AOC23/tools"
	"fmt"
	"strconv"
	"strings"

	tools "github.com/dsantos747/advent-of-code-2023/tools"
)

func recurse(chars string, groups []int, a, b int, cache [][]int) int {

	if !(a < len(chars)) { // Reached end of string
		if !(b < len(groups)) { // Reached end of groups arr
			return 1
		} else {
			return 0
		}
	}

	// Return answer if cached
	if (cache[a][b] != -1) {
		return cache[a][b]
	}

	result := 0

	if (chars[a] == '.') {
		result = recurse(chars,groups,a+1,b,cache)
	}
	if (chars[a] == '?') {
		result += recurse(chars,groups,a+1,b,cache)
	}
	if (b < len(groups)) {
		c := 0
		for i := a; i < len(chars); i++ {
			if c > groups[b] || chars[i] == '.' || c == groups[b] && chars[i] == '?' {
				break
			}
			c += 1
		}

		if c == groups[b] {
			if a+c < len(chars) && chars[a+c] != '#' {
				result += recurse(chars, groups, a+c+1, b+1, cache)
			} else {
				result += recurse(chars, groups, a+c, b+1, cache)
			}
		}
	}

	cache[a][b] = result
	return result
}

func part1(input []string) (int) {
	result := 0
	for _,line := range input {
		chars := strings.Split(line," ")[0]
		groupsStr := strings.Split(strings.Split(line," ")[1],",")
		groups := []int{}
		for _,g := range groupsStr {
			val,_ := strconv.Atoi(g)
			groups = append(groups, val)
		}
		var cache [][]int
		for a:=0; a<len(chars); a++ {
			cache = append(cache, make([]int, len(groups)+1))
			for b:=0; b<len(groups)+1; b++ {
				cache[a][b] = -1
			}
		}
		result += recurse(chars,groups,0,0,cache)
	}
	return result
}

func part2(input []string) (int) {
	result := 0
	for _,line := range input {
		str := strings.Split(line," ")
		chars := str[0] + "?" + str[0] + "?" + str[0] + "?" + str[0] + "?" + str[0]
		groupStr := str[1] + "," + str[1] + "," + str[1] + "," + str[1] + "," + str[1]
		groupSplit := strings.Split(groupStr,",")
		groups := []int{}
		for _,g := range groupSplit {
			val,_ := strconv.Atoi(g)
			groups = append(groups, val)
		}
		var cache [][]int
		for a:=0; a<len(chars); a++ {
			cache = append(cache, make([]int, len(groups)+1))
			for b:=0; b<len(groups)+1; b++ {
				cache[a][b] = -1
			}
		}
		result += recurse(chars,groups,0,0,cache)
	}
	return result
}

func main() {
	data,err := tools.ReadInput("./input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input := strings.Split(data, "\n")

	p1 := part1(input)
	fmt.Println("The answer to part 1 is",p1)

	p2 := part2(input)
	fmt.Println("The answer to part 2 is",p2)
}