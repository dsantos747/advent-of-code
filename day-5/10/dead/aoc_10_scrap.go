package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type RuleSet struct {
	seeds   [][]int
	soil   []Rule
	fert   []Rule
	water   []Rule
	light   []Rule
	temp   []Rule
	hum   []Rule
	loc   []Rule
}

type Rule struct {
	dest, source, spread int
}

func convertNumber(num int, rules []Rule) int {
    for _, rule := range rules {
        if (num >= rule.source) && (num < rule.source+rule.spread) { // num <=  or < ?
            return num + (rule.dest - rule.source)
        }
    }
    return num
}

func parseInput(game string, index int, almanac RuleSet) (int, RuleSet){
	if (game == "") {
		return index, almanac
	}
	keys := []string{"seeds","soil","fert","water","light","temp","hum","loc"}
	reMap := regexp.MustCompile(`map:`)

	if (reMap.Match([]byte(game))) {
		index++
	} else {
		if (index == 0) {
			game = game[7:]
			splitGame := strings.Split(game," ")
			for i := 0; i < len(splitGame); i+=2 {
				num1,_ := strconv.Atoi(splitGame[i])
				num2,_ := strconv.Atoi(splitGame[i+1])
				almanac.seeds = append(almanac.seeds, []int{num1, num2})
			}
		} else {
			splitGame := strings.Split(game," ")
			dest,_ := strconv.Atoi(splitGame[0])
			source,_ := strconv.Atoi(splitGame[1])
			spread,_ := strconv.Atoi(splitGame[2])
			// var dest, source, spread int
			// _, err := fmt.Sscanf(game, "%d %d %d", &dest, &source, &spread)
			// if err != nil {
			// 	fmt.Println("Error parsing numbers:", err)
			// 	return index, almanac
			// }
			rule := Rule{dest, source, spread}
			switch keys[index] {
			case "soil":
				almanac.soil = append(almanac.soil, rule)
			case "fert":
				almanac.fert = append(almanac.fert, rule)
			case "water":
				almanac.water = append(almanac.water, rule)
			case "light":
				almanac.light = append(almanac.light, rule)
			case "temp":
				almanac.temp = append(almanac.temp, rule)
			case "hum":
				almanac.hum = append(almanac.hum, rule)
			case "loc":
				almanac.loc = append(almanac.loc, rule)
			}
		}
	}

	return index, almanac
}



func main() {
	input,err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error loading input", err)
		return
	}
	defer input.Close()

	var minLoc int = int(^uint(0) >> 1)
	var index int = 0
	var almanac RuleSet

	sc := bufio.NewScanner(input)
	for sc.Scan() {
		index,almanac = parseInput(sc.Text(),index,almanac)
	}

	for _,seedRange := range almanac.seeds {
			for i:= 0; i <seedRange[1]; i++ {
				seed := seedRange[0]+i
				// That's it. You've lost your mind
				loc := convertNumber(convertNumber(convertNumber(convertNumber(convertNumber(convertNumber(convertNumber(seed,almanac.soil),almanac.fert),almanac.water),almanac.light),almanac.temp),almanac.hum),almanac.loc)
				if (loc < minLoc) {
					minLoc = loc
				}
			}
		fmt.Println("processed a seed range")
	}
	returnMin := strconv.Itoa(minLoc)

	fmt.Println("The lowest valid location is",returnMin)
}