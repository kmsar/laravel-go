package Slices

type Empty struct{}

// ArrayUnique removes duplicate values from an array,
// if the input is empty then return the original input
func ArrayUnique[T comparable](arr []T) []T {
	set := make(map[T]Empty)

	for _, v := range arr {
		set[v] = Empty{}
	}

	result := make([]T, 0)
	for k := range set {
		result = append(result, k)
	}

	return result
}
