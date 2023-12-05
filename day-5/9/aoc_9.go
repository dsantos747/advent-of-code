package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func convertNumber(num int, rules []Rule) int {
    for _, rule := range rules {
        if (num >= rule.source) && (num <= rule.source+rule.spread) {
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
			for _,num := range splitGame {
				gameNum,_ := strconv.Atoi(num)
				almanac.seeds = append(almanac.seeds, gameNum)
			}
		} else {
			var dest, source, spread int
			_, err := fmt.Sscanf(game, "%d %d %d", &dest, &source, &spread)
			if err != nil {
				fmt.Println("Error parsing numbers:", err)
				return index, almanac
			}
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

type RuleSet struct {
	seeds   []int
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

func main() {
	input,err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error loading input", err)
		return
	}
	defer input.Close()

	var minLoc int = 999999999999999999
	var index int = 0
	var almanac RuleSet

	sc := bufio.NewScanner(input)
	for sc.Scan() {
		index,almanac = parseInput(sc.Text(),index,almanac)
	}

	for _,seed := range almanac.seeds {
		// That's it. You've lost your mind
		loc := convertNumber(convertNumber(convertNumber(convertNumber(convertNumber(convertNumber(convertNumber(seed,almanac.soil),almanac.fert),almanac.water),almanac.light),almanac.temp),almanac.hum),almanac.loc)
		if (loc < minLoc) {
			minLoc = loc
		}
	}

	fmt.Println("The lowest valid location is",minLoc)
}