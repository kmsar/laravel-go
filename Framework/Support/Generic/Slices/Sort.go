package Slices

import (
	"golang.org/x/exp/constraints"
	"sort"
)

// Sort sorts an array (lowest to highest)
func Sort[T constraints.Ordered](array []T) []T {
	if len(array) == 0 {
		return array
	}

	sort.Slice(array, func(i int, j int) bool {
		return array[i] < array[j]
	})

	return array
}
