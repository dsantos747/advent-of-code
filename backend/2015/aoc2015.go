package aoc2023

import (
	"fmt"

	day1 "github.com/dsantos747/advent-of-code/2015/day-1"
	day2 "github.com/dsantos747/advent-of-code/2015/day-2"
	day3 "github.com/dsantos747/advent-of-code/2015/day-3"
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
