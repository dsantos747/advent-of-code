package aoc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	day1 "github.com/dsantos747/advent-of-code-2023/day-1"
	"github.com/rs/cors"
)

type RequestBody struct {
	Day   int    `json:"day"`
	Input string `json:"input"`
}

type DayFunc func(input string) (int, int, error)

var DayFunctions = []DayFunc{
	day1.Solve,
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

	fmt.Fprintf(w, "The answer to part 1 is %d.<br/>The answer to part 2 is %d.", p1, p2)

}
