package main

import (
	"fmt"
	"strings"

	tools "github.com/dsantos747/advent-of-code-2023/tools"
)

type Signal struct {
	src   string
	dest  string
	value int
}

type Module struct {
	name  string
	state int
	op    Operator
	ins   []Signal
	dests []string
}

type ModuleMap map[string]Module

type Queue []Signal

type Operator interface {
	Apply(moduleMap ModuleMap, module *Module, queue *Queue, count *[2]int, signal Signal)
}

type Broadcast struct{}

func (bc Broadcast) Apply(moduleMap ModuleMap, module *Module, queue *Queue, count *[2]int, signal Signal) {
	for _, dest := range module.dests {
		(*count)[0]++
		*queue = append(*queue, Signal{module.name, dest, 0})
	}
}

type FlipFlop struct{}

func (ff FlipFlop) Apply(moduleMap ModuleMap, module *Module, queue *Queue, count *[2]int, signal Signal) {
	if signal.value == 0 {
		if module.state == 0 {
			originalModule := moduleMap[signal.dest]
			originalModule.state = 1
			moduleMap[signal.dest] = originalModule
			for _, dest := range module.dests {
				(*count)[1]++
				*queue = append(*queue, Signal{module.name, dest, 1})
			}
		} else {
			originalModule := moduleMap[signal.dest]
			originalModule.state = 0
			moduleMap[signal.dest] = originalModule
			for _, dest := range module.dests {
				(*count)[0]++
				*queue = append(*queue, Signal{module.name, dest, 0})
			}
		}
	}
}

type Conjunct struct{}

func (co Conjunct) Apply(moduleMap ModuleMap, module *Module, queue *Queue, count *[2]int, signal Signal) {
	high := false
	for i, input := range module.ins {
		if input.src == signal.src {
			originalModule := moduleMap[signal.dest]
			originalModule.ins[i].value = signal.value
			moduleMap[signal.dest] = originalModule
			if signal.value == 0 {
				high = true
			}
		} else if input.value == 0 {
			high = true
		}
	}

	if high {
		for _, dest := range module.dests {
			(*count)[1]++
			*queue = append(*queue, Signal{module.name, dest, 1})
		}
	} else {
		for _, dest := range module.dests {
			(*count)[0]++
			*queue = append(*queue, Signal{module.name, dest, 0})
		}
	}
}

type Output struct{}

func (op Output) Apply(moduleMap ModuleMap, module *Module, queue *Queue, count *[2]int, signal Signal) {
}

func parse(input []string) ModuleMap {
	res := make(ModuleMap)
	var name string
	var op Operator
	var dests []string

	for _, line := range input {
		s1 := strings.Split(line, " -> ")
		dests = strings.Split(s1[1], ", ")
		if s1[0][0] == '%' {
			name = s1[0][1:]
			op = FlipFlop{}
		} else if s1[0][0] == '&' {
			name = s1[0][1:]
			op = Conjunct{}
		} else {
			name = "broadcaster"
			op = Broadcast{}
		}
		res[name] = Module{name: name, state: 0, op: op, ins: []Signal{}, dests: dests}
	}
	res["output"] = Module{name: "output", state: 0, op: Output{}, ins: []Signal{}, dests: []string{}} // Used for test2 input

	for _, mod := range res {
		for _, dest := range mod.dests {
			destMod := res[dest]
			destMod.ins = append(destMod.ins, Signal{src: mod.name, dest: destMod.name, value: 0})
			res[dest] = destMod
		}
	}

	return res
}

func part1(input []string) int {
	moduleMap := parse(input)
	var queue Queue
	count := [2]int{0, 0}

	for i := 0; i < 4000; i++ {
		queue = append(queue, Signal{"button", "broadcaster", 0})
		count[0]++
		for {
			if len(queue) == 0 {
				break
			}

			currModule := queue[0].dest
			module := moduleMap[currModule]
			if module.op != nil {
				moduleMap[currModule].op.Apply(moduleMap, &module, &queue, &count, queue[0])

			}

			queue = queue[1:]
		}
	}

	// fmt.Println("Counts:", count)
	product := count[0] * count[1]
	return product
}

func part2(input []string) int {
	// The catch with this one is that you're looking for when rx is sent a low signal, but by studying the input you see
	// rx is fed solely by a conjunction module (in my case, kj). Therefore, rx will only receive a low signal when all inputs
	// to kj are high. Therefore, it should be possible to figure out how many presses it takes for each of kj's inputs to be
	// high, then use LCM to determine when they will ALL be high.

	moduleMap := parse(input)
	var queue Queue
	count := [2]int{0, 0}
	inToRx := moduleMap["rx"].ins[0].src
	keyInputs := []string{}
	for _, in := range moduleMap[inToRx].ins {
		keyInputs = append(keyInputs, in.src)
	}

	targetLength := len(keyInputs)
	buttonCounts := []int{}
	for i := 0; ; i++ {

		queue = append(queue, Signal{"button", "broadcaster", 0})
		count[0]++
		for {
			if len(queue) == 0 {
				break
			}

			currModule := queue[0].dest
			module := moduleMap[currModule]

			if queue[0].dest == inToRx {
				for j, key := range keyInputs {
					for _, in := range module.ins {
						if in.src == key && in.value == 1 {
							buttonCounts = append(buttonCounts, i+1)
							keyInputs = append(keyInputs[:j], keyInputs[j+1:]...)
						}
					}
				}

			}

			if len(buttonCounts) == targetLength {
				break
			}

			if module.op != nil {
				moduleMap[currModule].op.Apply(moduleMap, &module, &queue, &count, queue[0])
			}

			queue = queue[1:]
		}
		if len(buttonCounts) == targetLength {
			break
		}
	}

	// fmt.Println(buttonCounts)
	lcm := tools.LCM(buttonCounts)
	// fmt.Println("Counts:", count)
	return lcm
}

func main() {
	data, err := tools.ReadInput("./input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input := strings.Split(data, "\n")

	p1 := part1(input)
	fmt.Println("The answer to part 1 is", p1)

	p2 := part2(input)
	fmt.Println("The answer to part 2 is", p2)
}
