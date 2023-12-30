package tools

import (
	"fmt"
	"strings"
)

type ij struct {
	i, j int
}

// Visualisation
func PrintPath(input []string, path []ij) {
	var vis [][]string
	for _, line := range input {
		vis = append(vis, strings.Split(line, ""))
	}
	for _, coord := range path {
		vis[coord.i][coord.j] = "\033[0;32m█\033[0m"
	}
	for _, line := range vis {
		fmt.Println(line)
	}
}
