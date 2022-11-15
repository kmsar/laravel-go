package linq

// ElementAt returns the element at a specified index in a sequence.
// src := []int{1, 2, 3, 4, 5}

// r, err := linq.ElementAt(linq.FromSlice(src), 3)
// exp := 4
// _, err = linq.ElementAt(linq.FromSlice(src), 10)
func ElementAt[T any, E IEnumerable[T]](src E, n int) (def T, err error) {
	e := src()
	var v T
	for i := 0; i <= n; i++ {
		v, err = e.Next()
		if err != nil {
			if isEOC(err) {
				err = OutOfRange
			}
			return def, err
		}
	}
	return v, nil
}

// ElementAtOrDefault returns the element at a specified index in a sequence or a default value if the index is out of range.
//	r, err = linq.ElementAtOrDefault(linq.FromSlice(src), 10)
func ElementAtOrDefault[T any, E IEnumerable[T]](src E, n int) (T, error) {
	v, err := ElementAt(src, n)
	if isOutOfRange(err) {
		err = nil
	}
	return v, err
}
