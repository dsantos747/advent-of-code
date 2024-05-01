package day7

import (
	"strconv"
	"strings"
)

// Credit to reddit user u/coussej - was stumped on this one

type output struct {
	result uint16
	known  bool
	source []string
}

type network map[string]*output

func makeNetwork(input []string) network {
	network := network{}

	for _, line := range input {
		split := strings.Split(line, " ")
		dest := split[len(split)-1]
		input := split[:len(split)-2]
		network[dest] = &output{0, false, input}
	}
	return network
}

// By providing the target wire ("a"), this function uses recursion to calculate the value of all wires
func (n network) getNetworkOutput(key string) uint16 {
	if !n[key].known {
		var v1 uint16

		switch len(n[key].source) {
		case 1: // direct assignment
			val, err := strconv.Atoi(n[key].source[0])
			if err != nil {
				v1 = n.getNetworkOutput(n[key].source[0])
			} else {
				v1 = uint16(val)
			}
			n[key].result = v1

		case 2: // Two word instruction == NOT (e.g. NOT lx)
			val, err := strconv.Atoi(n[key].source[1])
			if err != nil {
				v1 = n.getNetworkOutput(n[key].source[1])
			} else {
				v1 = uint16(val)
			}
			n[key].result = uint16(^v1)

		case 3: // All remaining operations
			var v2 uint16
			val, err := strconv.Atoi(n[key].source[0])
			if err != nil {
				v1 = n.getNetworkOutput(n[key].source[0])
			} else {
				v1 = uint16(val)
			}
			val, err = strconv.Atoi(n[key].source[2])
			if err != nil {
				v2 = n.getNetworkOutput(n[key].source[2])
			} else {
				v2 = uint16(val)
			}

			switch n[key].source[1] { // Determine the appropriate operation
			case "AND":
				n[key].result = v1 & v2
			case "OR":
				n[key].result = v1 | v2
			case "RSHIFT":
				n[key].result = v1 >> v2
			case "LSHIFT":
				n[key].result = v1 << v2
			}
		}

	}
	n[key].known = true
	return n[key].result
}

func part1(input []string) int {
	network := makeNetwork(input)

	result := network.getNetworkOutput("a")
	return int(result)
}

func part2(input []string, p1 int) int {
	network := makeNetwork(input)

	network["b"].source = []string{strconv.Itoa(p1)}

	result := network.getNetworkOutput("a")
	return int(result)
}

func Solve(data string) (*int, *int, error) {
	input := strings.Split(data, "\n")

	p1 := part1(input)
	p2 := part2(input, p1)

	return &p1, &p2, nil
}
