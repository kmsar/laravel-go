package linq

// Count returns a number that represents how many elements in the specified sequence
    //	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// r, err := linq.Count(linq.FromSlice(src))
	// exp := 9
	// if r != exp
func Count[T any, E IEnumerable[T]](src E) (int, error) {
	c := 0
	err := ForEach(src, func(v T) error {
		c++
		return nil
	})
	if err != nil {
		return 0, err
	}
	return c, nil
}
