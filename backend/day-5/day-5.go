package day5

import (
	"slices"
	"strconv"
	"strings"
)

type RuleSet struct {
	seeds   []int
	seedLoc []int
	soil    Rule
	fert    Rule
	water   Rule
	light   Rule
	temp    Rule
	hum     Rule
	loc     Rule
}

type Rule struct {
	dest   []int
	source []int
	size   []int
}

type interval struct {
	start int // inclusive
	end   int // exclusive
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
		size, _ := strconv.Atoi(vals[2])
		pm.dest = append(pm.dest, dest)
		pm.source = append(pm.source, source)
		pm.size = append(pm.size, size)
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

func (a *RuleSet) seedRangeToLoc(r []interval) []interval {
	r = a.soil.NextPosRange(r)
	r = a.fert.NextPosRange(r)
	r = a.water.NextPosRange(r)
	r = a.light.NextPosRange(r)
	r = a.temp.NextPosRange(r)
	r = a.hum.NextPosRange(r)
	r = a.loc.NextPosRange(r)
	return r
}

func (p *Rule) NextPos(c int) int {
	for i, source := range p.source {
		size := p.size[i]
		dest := p.dest[i]
		if c >= source && c < source+size {
			return c + dest - source
		}
	}
	return c
}

func (p *Rule) NextPosRange(r []interval) []interval {
	a := []interval{}

	for i, source := range p.source {
		size := p.size[i]
		dest := p.dest[i]
		src_end := source + size
		nr := []interval{}
		for len(r) > 0 {
			curr := r[len(r)-1]
			r = r[:len(r)-1]
			before := interval{curr.start, min(curr.end, source)}
			inter := interval{max(curr.start, source), min(src_end, curr.end)}
			after := interval{max(curr.start, src_end), curr.end}
			if before.end > before.start {
				nr = append(nr, before)
			}
			if inter.end > inter.start {
				a = append(a, interval{inter.start - source + dest, inter.end - source + dest})
			}
			if after.end > after.start {
				nr = append(nr, after)
			}
		}
		r = nr
	}
	a = append(a, r...)
	return a
}

func (a *RuleSet) RangeLocs() {
	for i := 0; i < len(a.seeds); i += 2 {
		start := a.seeds[i]
		size := a.seeds[i+1]
		minLoc := int(^uint(0) >> 1)
		r := a.seedRangeToLoc([]interval{{start, start + size}})

		for _, rm := range r {
			minLoc = min(minLoc, rm.start)
		}
		a.seedLoc = append(a.seedLoc, minLoc)
	}
}

func (a *RuleSet) SeedLocs() {
	for _, seed := range a.seeds {
		minLoc := int(^uint(0) >> 1)
		loc := a.seedToLoc(seed)
		minLoc = min(minLoc, loc)
		a.seedLoc = append(a.seedLoc, minLoc)
	}
}

func part1(input []string) int {
	var almanac RuleSet
	almanac.seeds = parseSeeds(input[0])
	almanac.soil = parseMap(input[1])
	almanac.fert = parseMap(input[2])
	almanac.water = parseMap(input[3])
	almanac.light = parseMap(input[4])
	almanac.temp = parseMap(input[5])
	almanac.hum = parseMap(input[6])
	almanac.loc = parseMap(input[7])
	almanac.SeedLocs()
	return slices.Min(almanac.seedLoc)
}

func part2(input []string) int {
	var almanac RuleSet
	almanac.seeds = parseSeeds(input[0])
	almanac.soil = parseMap(input[1])
	almanac.fert = parseMap(input[2])
	almanac.water = parseMap(input[3])
	almanac.light = parseMap(input[4])
	almanac.temp = parseMap(input[5])
	almanac.hum = parseMap(input[6])
	almanac.loc = parseMap(input[7])
	almanac.RangeLocs()
	return slices.Min(almanac.seedLoc)
}

func Solve(data string) (int, int, error) {
	input := strings.Split(data, "\n\n")

	p1 := part1(input)
	p2 := part2(input)

	return p1, p2, nil
}
