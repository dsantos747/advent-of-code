package day7

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type Round struct {
	hand  string
	bid   int
	score int
}

func handScore(hand map[rune]int) int {
	if len(hand) == 1 {
		return 6
	} else if len(hand) == 2 {
		for _, val := range hand {
			if val == 4 {
				return 5
			} else if val == 3 {
				return 4
			}
		}
	} else if len(hand) == 3 {
		for _, val := range hand {
			if val == 3 {
				return 3
			} else if val == 2 {
				return 2
			}
		}
	} else if len(hand) == 4 {
		return 1
	}
	return 0
}

func cardScore(hand string) int {
	var cardMap = map[rune]int{'A': 12, 'K': 11, 'Q': 10, 'J': 9, 'T': 8, '9': 7, '8': 6, '7': 5, '6': 4, '5': 3, '4': 2, '3': 1, '2': 0}
	score := 0
	cardCount := make(map[rune]int)
	for i, letter := range hand {
		score += cardMap[letter] * int(math.Pow(100, float64(4-i)))
		cardCount[letter]++
	}
	score += handScore(cardCount) * int(math.Pow(100, 5))
	return score
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

	return handScore(hand)
}

func p2cardScore(hand string) int {
	var cardMap = map[rune]int{'A': 12, 'K': 11, 'Q': 10, 'J': 0, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1}
	score := 0
	cardCount := make(map[rune]int)
	for i, letter := range hand {
		score += cardMap[letter] * int(math.Pow(100, float64(4-i)))
		cardCount[letter]++
	}
	score += p2findHand(cardCount) * int(math.Pow(100, 5))
	return score
}

func part1(input []string) int {
	score := 0
	var camelCards []Round
	for _, game := range input {
		var r Round
		r.hand = strings.Fields(game)[0]
		r.bid, _ = strconv.Atoi(strings.Fields(game)[1])
		r.score = cardScore(r.hand)
		camelCards = append(camelCards, r)
	}

	sort.Slice(camelCards, func(i, j int) bool {
		return camelCards[i].score < camelCards[j].score
	})

	for rank, game := range camelCards {
		score += game.bid * (rank + 1)
	}
	return score
}

func part2(input []string) int {
	score := 0
	var camelCards []Round
	for _, game := range input {
		var r Round
		r.hand = strings.Fields(game)[0]
		r.bid, _ = strconv.Atoi(strings.Fields(game)[1])
		r.score = p2cardScore(r.hand)
		camelCards = append(camelCards, r)
	}

	sort.Slice(camelCards, func(i, j int) bool {
		return camelCards[i].score < camelCards[j].score
	})

	for rank, game := range camelCards {
		score += game.bid * (rank + 1)
	}
	return score
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input)

	return &p1, &p2, nil
}
