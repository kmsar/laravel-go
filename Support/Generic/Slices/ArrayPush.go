package Slices

// ArrayPush pushes one or more elements onto the end of array,
// returns the new number of elements in the array
func ArrayPush[T comparable](s *[]T, elements ...T) int {
	if s == nil {
		return 0
	}
	*s = append(*s, elements...)
	return len(*s)
}
