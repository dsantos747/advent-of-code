package day4

import (
	"fmt"
	"math"
	"strings"

	"crypto/md5"
	"encoding/hex"
)

func part1(input string) int {
	for i := 1; i <= math.MaxInt64; i++ {
		h := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
		hash := hex.EncodeToString(h[:])

		if strings.HasPrefix(hash, "00000") {
			return i
		}
	}

	return math.MaxInt64
}

func part2(input string) int {

	for i := 1; i <= math.MaxInt64; i++ {
		h := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
		hash := hex.EncodeToString(h[:])

		if strings.HasPrefix(hash, "000000") {
			return i
		}
	}

	return math.MaxInt64
}

func Solve(data string) (*int, *int, error) {
	p1 := part1(data)
	p2 := part2(data)

	return &p1, &p2, nil
}
