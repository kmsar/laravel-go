package Slices

// ArrayPop pops the element off the end of array
func ArrayPop[T comparable](s *[]T) T {
	var t T
	if s == nil || len(*s) == 0 {
		return t
	}

	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]

	return e
}
