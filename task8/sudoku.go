package task8

type Sudoku [9]sudokuSet

type sudokuSet [9]int

func (s Sudoku) row(i int) sudokuSet {
	return s[i]
}

func (s Sudoku) col(i int) sudokuSet {
	var column [9]int
	for j := 0; j < 9; j++ {
		column[j] = s[j][i]
	}
	return column
}

func (s Sudoku) box(i int) sudokuSet {
	var box [9]int
	for j := 0; j < 9; j++ {
		box[j] = s[i/3*3+j/3][i%3*3+j%3]
	}
	return box
}

func (s Sudoku) Check() bool {
	for i := 0; i < 9; i++ {
		if !s.row(i).check() || !s.col(i).check() || !s.box(i).check() {
			return false
		}
	}
	return true
}

func (ss sudokuSet) check() bool {
	var x [9]bool
	for _, v := range ss {
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
