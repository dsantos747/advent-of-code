package aoc2023

import (
	"fmt"

	day1 "github.com/dsantos747/advent-of-code/2023/day-1"
	day10 "github.com/dsantos747/advent-of-code/2023/day-10"
	day11 "github.com/dsantos747/advent-of-code/2023/day-11"
	day12 "github.com/dsantos747/advent-of-code/2023/day-12"
	day13 "github.com/dsantos747/advent-of-code/2023/day-13"
	day14 "github.com/dsantos747/advent-of-code/2023/day-14"
	day15 "github.com/dsantos747/advent-of-code/2023/day-15"
	day16 "github.com/dsantos747/advent-of-code/2023/day-16"
	day17 "github.com/dsantos747/advent-of-code/2023/day-17"
	day18 "github.com/dsantos747/advent-of-code/2023/day-18"
	day19 "github.com/dsantos747/advent-of-code/2023/day-19"
	day2 "github.com/dsantos747/advent-of-code/2023/day-2"
	day20 "github.com/dsantos747/advent-of-code/2023/day-20"
	day21 "github.com/dsantos747/advent-of-code/2023/day-21"
	day22 "github.com/dsantos747/advent-of-code/2023/day-22"
	day23 "github.com/dsantos747/advent-of-code/2023/day-23"
	day24 "github.com/dsantos747/advent-of-code/2023/day-24"
	day25 "github.com/dsantos747/advent-of-code/2023/day-25"
	day3 "github.com/dsantos747/advent-of-code/2023/day-3"
	day4 "github.com/dsantos747/advent-of-code/2023/day-4"
	day5 "github.com/dsantos747/advent-of-code/2023/day-5"
	day6 "github.com/dsantos747/advent-of-code/2023/day-6"
	day7 "github.com/dsantos747/advent-of-code/2023/day-7"
	day8 "github.com/dsantos747/advent-of-code/2023/day-8"
	day9 "github.com/dsantos747/advent-of-code/2023/day-9"
)

type RequestBody struct {
	Day   int    `json:"day"`
	Input string `json:"input"`
}

type DayFunc func(input string) (*int, *int, error)

var DayFunctions = []DayFunc{
	day1.Solve,
	day2.Solve,
	day3.Solve,
	day4.Solve,
	day5.Solve,
	day6.Solve,
	day7.Solve,
	day8.Solve,
	day9.Solve,
	day10.Solve,
	day11.Solve,
	day12.Solve,
	day13.Solve,
	day14.Solve,
	day15.Solve,
	day16.Solve,
	day17.Solve,
	day18.Solve,
	day19.Solve,
	day20.Solve,
	day21.Solve,
	day22.Solve,
	day23.Solve,
	day24.Solve,
	day25.Solve,
}

func SolveDay(day int, input string) (*map[string]int, error) {
	if day < 0 || day >= len(DayFunctions) {
		err := fmt.Errorf("invalid day index")
		return nil, err
	}

	p1, p2, err := DayFunctions[day](input)
	if err != nil {
		err = fmt.Errorf("error executing day function: %w", err)
		return nil, err
	}

	response := map[string]int{
		"p1": *p1,
		"p2": *p2,
	}

	return &response, nil
}
