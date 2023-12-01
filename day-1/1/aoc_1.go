package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input,err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error loading input", err)
		return
	}
	defer input.Close()

	var calibSum int = 0
	pattern := `[0-9]`

	re,err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex", err)
		return
	}

	sc := bufio.NewScanner(input) // Create scanner instance
	for sc.Scan() { // Scan over each line of input (input.txt) and print
		matches := re.FindAll([]byte(sc.Text()),-1)
		calib,_ := strconv.Atoi(string(matches[0]) + string(matches[len(matches)-1]))
		calibSum += calib
	}

	fmt.Println("The sum of all calibration values is",calibSum)
}