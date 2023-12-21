package tools

func HCF(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(slice []int) int {
	lcm := 1
	for _, a := range slice {
		if a == 0 {
			return 0
		}
	}
	for _, num := range slice {
		lcm = lcm * num / HCF(lcm, num)
	}
	return lcm
}

func mod(a, n int) int {
	return ((a % n) + n) % n
}
