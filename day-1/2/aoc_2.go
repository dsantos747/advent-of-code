package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func numFormat(word string) string {
	wordMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":  "6",
		"seven":  "7",
		"eight":  "8",
		"nine":  "9",
	}

	if number, ok := wordMap[word]; ok {
		return number
	}
	return word
}

func main() {
	input,err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error loading input", err)
		return
	}
	defer input.Close()

	var calibSum int = 0
	pattern := `[0-9]|one|two|three|four|five|six|seven|eight|nine`



	re,err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex", err)
		return
	}

	sc := bufio.NewScanner(input) // Create scanner instance
	for sc.Scan() { // Scan over each line of input (input.txt) and print
		matches := re.FindAll([]byte(sc.Text()),-1)
		first := numFormat(string(matches[0]))
		last := numFormat(string(matches[len(matches)-1]))
		calib,_ := strconv.Atoi(first + last)
		calibSum += calib
	}

	fmt.Println("The sum of all calibration values is",calibSum)
}