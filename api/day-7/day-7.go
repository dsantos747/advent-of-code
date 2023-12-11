package main

import (
	// "AOC23/tools"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	tools "github.com/dsantos747/advent-of-code-2023/api/tools"
)

type Round struct {
	hand string
	bid int
	score int
}

func cardConvert(letter string) int {
	if (string(letter) == "A") {
		return 12
	} else if (string(letter) == "K") {
		return 11
	} else if (string(letter) == "Q") {
		return 10
	} else if (string(letter) == "J") {
		return 9
	} else if (string(letter) == "T") {
		return 8
	} else {
		val,_ := strconv.Atoi(string(letter))
		return val - 2
	}
}

func findHand(hand map[rune]int) int {
	if (len(hand) == 1) {
		return 6
	} else if (len(hand) == 2) {
		for _,val := range hand {
			if (val == 4) {
				return 5
			} else if (val == 3) {
				return 4
			}
		}
	} else if (len(hand) == 3) {
		for _,val := range hand {
			if (val == 3) {
				return 3
			} else if (val == 2) {
				return 2
			}
		}
	} else if (len(hand) == 4) {
		return 1
	} 
	return 0
}

func cardScore(hand string) int {
	score := 0
	cardCount := make(map[rune]int)
	for i,letter := range hand {
		score += cardConvert(string(letter)) * int(math.Pow(100,float64(4-i)))
		cardCount[letter]++
	}
	score += findHand(cardCount) * int(math.Pow(100,5))
	return score
}

func p2cardConvert(letter string) int {
	if (string(letter) == "A") {
		return 12
	} else if (string(letter) == "K") {
		return 11
	} else if (string(letter) == "Q") {
		return 10
	} else if (string(letter) == "J") {
		return 0
	} else if (string(letter) == "T") {
		return 9
	} else {
		val,_ := strconv.Atoi(string(letter))
		return val - 1
	}
}

func p2findHand(hand map[rune]int) int {
	if count, ok := hand[[]rune("J")[0]]; ok {
		delete(hand, []rune("J")[0])
		var maxKey rune
		var maxValue int

		for key, value := range hand {
			if maxValue == 0 || value > maxValue {
				maxKey = key
				maxValue = value
			}
		}
		hand[maxKey] += count
	}
	if (len(hand) == 1) {
		return 6
	} else if (len(hand) == 2) {
		for _,val := range hand {
			if (val == 4) {
				return 5
			} else if (val == 3) {
				return 4
			}
		}
	} else if (len(hand) == 3) {
		for _,val := range hand {
			if (val == 3) {
				return 3
			} else if (val == 2) {
				return 2
			}
		}
	} else if (len(hand) == 4) {
		return 1
	} 
	return 0
}

func p2cardScore(hand string) int {
	score := 0
	cardCount := make(map[rune]int)
	for i,letter := range hand {
		score += p2cardConvert(string(letter)) * int(math.Pow(100,float64(4-i)))
		cardCount[letter]++
	}
	score += p2findHand(cardCount) * int(math.Pow(100,5))
	return score
}

func part1(input []string) int {
	score := 0
	var camelCards []Round
	for _,game := range input {
		var r Round
		r.hand = strings.Fields(game)[0]
		r.bid,_ = strconv.Atoi(strings.Fields(game)[1])
		r.score = cardScore(r.hand)
		camelCards = append(camelCards, r)
	}

	sort.Slice(camelCards, func(i, j int) bool {
		return camelCards[i].score < camelCards[j].score
	})

	for rank,game := range camelCards {
		score += game.bid * (rank+1)
	}
	return score
}

func part2(input []string) int {
	score := 0
	var camelCards []Round
	for _,game := range input {
		var r Round
		r.hand = strings.Fields(game)[0]
		r.bid,_ = strconv.Atoi(strings.Fields(game)[1])
		r.score = p2cardScore(r.hand)
		camelCards = append(camelCards, r)
	}

	sort.Slice(camelCards, func(i, j int) bool {
		return camelCards[i].score < camelCards[j].score
	})

	for rank,game := range camelCards {
		score += game.bid * (rank+1)
	}
	return score
}

func main() {
	input,err := tools.ReadInput("./input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	data := strings.Split(input, "\n")
	p1 := part1(data)
	fmt.Println("The answer to part 1 is",p1)

	p2 := part2(data)
	fmt.Println("The answer to part 2 is",p2)
}