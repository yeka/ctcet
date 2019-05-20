package task1

import (
	"math"
)

func SearchClosest(v float64, array []float64) float64 {
	i := -1
	for k, x := range array {
		if i == -1 || math.Abs(v-x) < math.Abs(v-array[i]) || (math.Abs(v-x) == math.Abs(v-array[i]) && x > array[i]) {
			i = k
		}
	}
	return array[i]
}
