package tools

import (
	"math"
	"strconv"
	"strings"
)

func HexToInt(hex string) int {
	result := 0
	var char byte
	var val int

	hex = strings.ToLower(hex)

	for i := 0; i < len(hex); i++ {

		char = hex[len(hex)-1-i]

		if char == 'f' {
			val = 15
		} else if char == 'e' {
			val = 14
		} else if char == 'd' {
			val = 13
		} else if char == 'c' {
			val = 12
		} else if char == 'b' {
			val = 11
		} else if char == 'a' {
			val = 10
		} else {
			val, _ = strconv.Atoi(string(char))
		}

		result += val * int(math.Pow(16, float64(i)))
	}

	return result
}
