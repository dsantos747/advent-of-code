package tools

import "strings"

func FindSingleSubstring(input []string, character string) (i, j int) {
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

func InField(input []string, i, j int) bool {
	if i >= 0 && i < len(input) && j >= 0 && j < len(input[0]) {
		return true
	}
	return false
}
