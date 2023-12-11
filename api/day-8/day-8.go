package main

import (
	// "AOC23/tools"
	"fmt"
	"strconv"
	"strings"

	tools "github.com/dsantos747/advent-of-code-2023/api/tools"
)

func hcf(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func checkLcm(slice []int) int {
	lcm := 1
	for _,a := range slice {
		if (a == 0) {
			return 0
		}
	}
	for _,num := range slice {
		lcm = lcm * num / hcf(lcm,num)
	}
	return lcm

}


func part1(data []string) int {
	dirs := strings.ReplaceAll(strings.ReplaceAll(data[0],"L","0"),"R","1")
	routes := data[2:]

	mapMap := make(map[string][]string)
	for _,route := range routes {
		r := strings.Split(strings.ReplaceAll(strings.ReplaceAll(route,"(",""),")",""), "=")
		key := strings.TrimSpace(r[0])
		branches := strings.Split(r[1], ",")
		brLeft := strings.TrimSpace(branches[0])
		brRight := strings.TrimSpace(branches[1])
		mapMap[key] = []string{brLeft, brRight}
	}

	nextKey := "AAA"
	stepCount := 0
	for {
		for _,d := range dirs {
			dir,_ := strconv.Atoi(string(d))
			nextKey = mapMap[nextKey][dir]
			stepCount++
			if (nextKey=="ZZZ") {
				return stepCount
			}
	
		}
	}
}

func part2(data []string) int {
	dirs := strings.ReplaceAll(strings.ReplaceAll(data[0],"L","0"),"R","1")
	routes := data[2:]

	mapMap := make(map[string][]string)
	var keyList []string
	for _,route := range routes {
		r := strings.Split(strings.ReplaceAll(strings.ReplaceAll(route,"(",""),")",""), "=")
		key := strings.TrimSpace(r[0])
		if (string(key[2]) == "A") {
			keyList = append(keyList, string(key))
		}
		branches := strings.Split(r[1], ",")
		brLeft := strings.TrimSpace(branches[0])
		brRight := strings.TrimSpace(branches[1])
		mapMap[key] = []string{brLeft, brRight}
	}

	countSlice := make([]int, len(keyList))
	stepCount := 0
	for {
		for _,d := range dirs {
			dir,_ := strconv.Atoi(string(d))
			stepCount++

			for i,key := range keyList {
				keyList[i] = mapMap[key][dir]
				if (string(keyList[i][2]) == "Z") {
					countSlice[i] = stepCount
				}
				
			}
			lcm := checkLcm(countSlice)
			if (lcm > 0) {
				return lcm
			}

		}
	}
}

func main() {

	input,err := tools.ReadInput("./input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	data := strings.Split(input, "\n")

	// Note - comment this out if using p2 test input
	p1 := part1(data)
	fmt.Println("The answer to part 1 is",p1)

	p2 := part2(data)
	fmt.Println("The answer to part 2 is",p2)
}