package task4

func Calculate(a, b int) int {
	if b == 0 {
		return a
	}
	return Calculate(b, a%b)
}
