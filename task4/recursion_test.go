package task4

import "testing"

func TestRecursion(t *testing.T) {
	if x := Calculate(1071, 1029); x != 21 {
		t.Errorf("Unexpected result %v, expect 21\n", x)
	}
}
