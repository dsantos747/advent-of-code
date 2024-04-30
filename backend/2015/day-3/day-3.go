package day3

type Pos struct {
	i, j int
}

var dirMap = map[rune]Pos{
	'>': {0, 1},
	'<': {0, -1},
	'^': {-1, 0},
	'v': {1, 0},
}

func part1(input string) int {
	total := 1
	curr := Pos{0, 0}
	visited := map[Pos]int{curr: 1}

	for _, x := range input {
		dir := dirMap[x]
		curr = Pos{curr.i + dir.i, curr.j + dir.j}
		visited[curr]++
		if visited[curr] == 1 {
			total++
		}
	}

	return total
}

func part2(input string) int {
	total := 1
	santa := Pos{0, 0}
	robot := Pos{0, 0}
	var curr Pos
	visited := map[Pos]int{santa: 1}

	for i, x := range input {
		dir := dirMap[x]
		if i%2 == 0 {
			curr = Pos{santa.i + dir.i, santa.j + dir.j}
			santa = curr
		} else {
			curr = Pos{robot.i + dir.i, robot.j + dir.j}
			robot = curr
		}
		visited[curr]++
		if visited[curr] == 1 {
			total++
		}
	}

	return total
}

func Solve(data string) (*int, *int, error) {

	p1 := part1(data)
	p2 := part2(data)

	return &p1, &p2, nil
}
