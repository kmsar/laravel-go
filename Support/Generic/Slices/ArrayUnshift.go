package Slices

// ArrayUnshift prepends one or more elements to the beginning of a array,
// returns the new number of elements in the array.
func ArrayUnshift[T comparable](s *[]T, elements ...T) int {
	if s == nil {
		return 0
	}
	*s = append(elements, *s...)
	return len(*s)
}
