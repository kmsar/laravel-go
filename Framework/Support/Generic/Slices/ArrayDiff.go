package Slices

// ArrayDiff computes the difference of arrays
func ArrayDiff[T comparable](array1, array2 []T) []T {
	var res []T
	for _, v := range array1 {
		if !InArray(v, array2) {
			res = append(res, v)
		}
	}
	return res
}
