package aoc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	day1 "github.com/dsantos747/advent-of-code-2023/day-1"
	day10 "github.com/dsantos747/advent-of-code-2023/day-10"
	day11 "github.com/dsantos747/advent-of-code-2023/day-11"
	day2 "github.com/dsantos747/advent-of-code-2023/day-2"
	day3 "github.com/dsantos747/advent-of-code-2023/day-3"
	day4 "github.com/dsantos747/advent-of-code-2023/day-4"
	day5 "github.com/dsantos747/advent-of-code-2023/day-5"
	day6 "github.com/dsantos747/advent-of-code-2023/day-6"
	day7 "github.com/dsantos747/advent-of-code-2023/day-7"
	day8 "github.com/dsantos747/advent-of-code-2023/day-8"
	day9 "github.com/dsantos747/advent-of-code-2023/day-9"
	"github.com/rs/cors"
)

type RequestBody struct {
	Day   int    `json:"day"`
	Input string `json:"input"`
}

type DayFunc func(input string) (int, int, error)

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
}

func init() {
	functions.HTTP("HelloAOC", HelloAOC)
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloAOC)

	// Use CORS middleware
	handler := cors.Default().Handler(mux)

	// Start the server
	http.ListenAndServe(":8080", handler)
}

func HelloAOC(w http.ResponseWriter, r *http.Request) {

	var reqBody RequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqBody)
	if err != nil {
		// fmt.Fprintf(w, fmt.Sprintf("Failed to decode JSON data: %v", err))

		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}

	day := reqBody.Day - 1
	input := reqBody.Input

	if day < 0 || day >= len(DayFunctions) {
		http.Error(w, "Invalid day index", http.StatusBadRequest)
		return
	}

	p1, p2, err := DayFunctions[day](input)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing day function: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"p1": p1,
		"p2": p2,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
