package main

import (
	// "AOC23/tools"
	"fmt"
	"slices"
	"strconv"
	"strings"

	tools "github.com/dsantos747/advent-of-code-2023/api/tools"
)

type RuleSet struct {
	seeds   []int
	seedLoc []int
	soil   Rule
	fert   Rule
	water   Rule
	light   Rule
	temp   Rule
	hum   Rule
	loc   Rule
}

type Rule struct {
	dest   []int
	source []int
	spread []int
}

func parseSeeds(seeds string) []int {
	seeds = strings.Split(seeds, ": ")[1]
	seeds_split := strings.Split(seeds, " ")
	var seedSlice []int
	for _, seed := range seeds_split {
		s, _ := strconv.Atoi(seed)
		seedSlice = append(seedSlice, s)
	}
	return seedSlice
}

func parseMap(m string) Rule {
	lines := strings.Split(m, "\n")
	pm := &Rule{}
	for i := 1; i < len(lines); i++ {
		vals := strings.Split(lines[i], " ")
		dest, _ := strconv.Atoi(vals[0])
		source, _ := strconv.Atoi(vals[1])
		spread, _ := strconv.Atoi(vals[2])
		pm.dest = append(pm.dest, dest)
		pm.source = append(pm.source, source)
		pm.spread = append(pm.spread, spread)
	}
	return *pm
}

func (a *RuleSet) seedToLoc(c int) int {
	c = a.soil.NextPos(c)
	c = a.fert.NextPos(c)
	c = a.water.NextPos(c)
	c = a.light.NextPos(c)
	c = a.temp.NextPos(c)
	c = a.hum.NextPos(c)
	c = a.loc.NextPos(c)
	return c
}

func (p *Rule) NextPos(c int) int {
	for i, source := range p.source {
		spread := p.spread[i]
		dest := p.dest[i]
		if c >= source && c < source+spread {
			return dest + (c - source)
		}
	}
	return c
}

func (a *RuleSet) SeedRangeToLocs() {
	for i := 0; i < len(a.seeds); i += 2 {
		seed := a.seeds[i]
		totalSeeds := a.seeds[i+1]
		minLoc := int(^uint(0) >> 1)
		for j := seed; j < seed+totalSeeds; j++ {
			loc := a.seedToLoc(j)
			if loc < minLoc {
				minLoc = loc
			}
		}
		a.seedLoc = append(a.seedLoc, minLoc)
	}
}



func main() {
	input,err := tools.ReadInput("../input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	data := strings.Split(input, "\n\n")

	var almanac RuleSet

	almanac.seeds = parseSeeds(data[0])
	almanac.soil = parseMap(data[1])
	almanac.fert = parseMap(data[2])
	almanac.water = parseMap(data[3])
	almanac.light = parseMap(data[4])
	almanac.temp = parseMap(data[5])
	almanac.hum = parseMap(data[6])
	almanac.loc = parseMap(data[7])
	almanac.SeedRangeToLocs()
	min := slices.Min(almanac.seedLoc)
	fmt.Println("The lowest valid location is",strconv.Itoa(min))
}