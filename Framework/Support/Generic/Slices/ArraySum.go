package Slices

import "golang.org/x/exp/constraints"

// ArraySum returns the sum of values in an array
func ArraySum[T constraints.Ordered](array []T) T {
	var sum T
	for _, v := range array {
		sum += v
	}

	return sum
}
