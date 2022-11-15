package Slices

// ArrayShift shifts an element off the beginning of array
func ArrayShift[T comparable](s *[]T) T {
	var t T
	if s == nil || len(*s) == 0 {
		return t
	}

	f := (*s)[0]
	*s = (*s)[1:]

	return f
}
