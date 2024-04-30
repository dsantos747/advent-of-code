package aoc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	aoc2023 "github.com/dsantos747/advent-of-code/2023"
)

type RequestBody struct {
	Year  int    `json:"year"`
	Day   int    `json:"day"`
	Input string `json:"input"`
}

type YearFunc func(day int, input string) (*map[string]int, error)

var YearFunctions = map[int]YearFunc{
	2023: aoc2023.SolveDay,
}

// Is this necessary?
func init() {
	functions.HTTP("HelloAOC", HelloAOC)
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", HelloAOC)

	// // Use CORS middleware
	// handler := cors.Default().Handler(mux)

	// // Start the server
	// http.ListenAndServe(":8080", handler)
}

func HelloAOC(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqBody)
	if err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}

	day := reqBody.Day - 1
	input := strings.ReplaceAll(string(reqBody.Input), "\r\n", "\n")

	if _, ok := YearFunctions[reqBody.Year]; !ok {
		http.Error(w, "invalid year provided", http.StatusBadRequest)
		return
	}

	response, err := YearFunctions[reqBody.Year](day, input)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
