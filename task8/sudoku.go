package task8

func SudokuCheck(matrix [9][9]int) bool {
	for i := 0; i < 9; i++ {
		row := matrix[i]
		var column [9]int
		var box [9]int

		for j := 0; j < 9; j++ {
			column[j] = matrix[j][i]
			box[j] = matrix[i/3*3+j/3][i%3*3+j%3]
		}

		if !completeSet(row) || !completeSet(column) || !completeSet(box) {
			return false
		}
	}

	return true
}

// completeSet checks if array given contains all number from 1 to 9
func completeSet(arr [9]int) bool {
	var x [9]bool
	for _, v := range arr {
		if v < 1 && v > 9 {
			return false
		}
		x[v-1] = true
	}

	for _, v := range x {
		if !v {
			return false
		}
	}

	return true
}
