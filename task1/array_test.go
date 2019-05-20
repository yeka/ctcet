package task1

import "testing"

func TestSearchClosest(t *testing.T) {
	array := []float64{-1.5, 0, 4.4, 5, 6, 7}
	testTable := map[float64]float64{
		4.5: 4.4,
		5.5: 6,
	}

	for v, expect := range testTable {
		actual := SearchClosest(v, array)
		if actual != expect {
			t.Errorf("Unexpected result %v, expect %v\n", actual, expect)
		}
	}
}
