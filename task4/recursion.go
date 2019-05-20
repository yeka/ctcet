package task4

func Calculate(a, b int) int {
	for {
		if b == 0 {
			return a
		}
		a, b = b, a % b
	}
}
