package Slices

import (
	"golang.org/x/exp/constraints"
	"sort"
)

// Rsort sorts an array in reverse order (highest to lowest)
func Rsort[T constraints.Ordered](array []T) []T {
	if len(array) == 0 {
		return array
	}

	sort.Slice(array, func(i int, j int) bool {
		return array[i] > array[j]
	})

	return array
}
