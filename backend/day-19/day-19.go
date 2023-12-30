package day19

import (
	"strconv"
	"strings"
)

type Condition struct {
	category int
	operator string
	value    int
}

type Step struct {
	condition *Condition
	result    string
}

type WorkFlow map[string][]Step

type Range struct {
	min int
	max int
}

func parseParts(input string) [][]int {
	parts := strings.Split(input, "\n")

	var res [][]int
	for _, part := range parts {
		part = part[1 : len(part)-1]
		ratings := strings.Split(part, ",")
		var subRes []int
		for _, r := range ratings {
			val, _ := strconv.Atoi(r[2:])
			subRes = append(subRes, val)
		}
		res = append(res, subRes)
	}
	return res
}

func parseWF(input string) WorkFlow {
	wfs := strings.Split(input, "\n")

	res := make(WorkFlow)
	for _, wf := range wfs {
		w := strings.SplitAfter(wf, "{")
		key := w[0][:len(w[0])-1]
		steps := strings.Split(w[1][:len(w[1])-1], ",")
		var stepSlice []Step
		for _, step := range steps {
			if strings.Contains(step, ":") {
				s := strings.Split(step, ":")
				var cat int
				if s[0][0] == 'x' {
					cat = 0
				} else if s[0][0] == 'm' {
					cat = 1
				} else if s[0][0] == 'a' {
					cat = 2
				} else {
					cat = 3
				}
				val, _ := strconv.Atoi(s[0][2:])
				cond := &Condition{cat, string(s[0][1]), val}
				stepSlice = append(stepSlice, Step{condition: cond, result: s[1]})
			} else {
				stepSlice = append(stepSlice, Step{result: step})
			}
		}
		res[key] = stepSlice
	}
	return res
}

func (wf WorkFlow) evalPart(part []int, wfKey string) bool {
	for _, step := range wf[wfKey] {
		if step.condition != nil {

			if (*step.condition).operator == "<" {
				if part[(*step.condition).category] >= (*step.condition).value {
					continue
				}
			} else {
				if part[(*step.condition).category] <= (*step.condition).value {
					continue
				}
			}
		}
		if step.result == "A" {
			return true
		} else if step.result == "R" {
			return false
		} else {
			return wf.evalPart(part, step.result)
		}
	}
	return false
}

func (wf WorkFlow) evalCombinations(wfKey string, ranges []Range) int {
	if wfKey == "R" {
		return 0
	}
	if wfKey == "A" {
		prod := 1
		for _, r := range ranges {
			prod *= r.max - r.min + 1
		}
		return prod
	}
	result := 0
	for _, step := range wf[wfKey] {
		if step.condition == nil {
			result += wf.evalCombinations(step.result, ranges)
			continue
		}

		nextRanges := make([]Range, len(ranges))
		copy(nextRanges, ranges)

		if (*step.condition).operator == "<" {
			nextRanges[(*step.condition).category].max = (*step.condition).value - 1
			ranges[(*step.condition).category].min = (*step.condition).value
			result += wf.evalCombinations(step.result, nextRanges)
		} else {
			nextRanges[(*step.condition).category].min = (*step.condition).value + 1
			ranges[(*step.condition).category].max = (*step.condition).value
			result += wf.evalCombinations(step.result, nextRanges)
		}
	}
	return result
}

func part1(input []string) int {
	result := 0

	wf := parseWF(input[0])
	parts := parseParts(input[1])

	for _, part := range parts {
		if wf.evalPart(part, "in") {
			sum := 0
			for _, rat := range part {
				sum += rat
			}
			result += sum
		}

	}

	return result
}

func part2(input []string) int {
	wf := parseWF(input[0])
	initRanges := []Range{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}}

	result := wf.evalCombinations("in", initRanges)

	return result
}

func Solve(data string) (int, int, error) {
	input := strings.Split(data, "\n\n")

	p1 := part1(input)
	p2 := part2(input)

	return p1, p2, nil
}
