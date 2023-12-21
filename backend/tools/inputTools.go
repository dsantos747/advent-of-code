package tools

import "strings"

func findSingleSubstring(input []string, character string) (i, j int) {
	for i, line := range input {
		if strings.Contains(line, character) {
			for j, char := range line {
				if string(char) == character {
					return i, j
				}
			}
		}
	}
	return 0, 0
}
