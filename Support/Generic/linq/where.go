package linq

type whereEnumerator[T any] struct {
	src  Enumerator[T]
	pred func(v T) (bool, error)
}

// Where filters a sequence of values based on a predicate.
// e1 := linq.FromSlice([]int16{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
// e2 := linq.Where(e1, func(v int16) (bool, error) { return v%2 != 0, nil })
// r, err := linq.ToSlice(e2)
// exp := []int16{9, 7, 5, 3, 1}
func Where[T any, E IEnumerable[T]](src E, pred func(v T) (bool, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &whereEnumerator[T]{src: src(), pred: pred}
	}
}

func (e *whereEnumerator[T]) Next() (def T, _ error) {
	for {
		v, err := e.src.Next()
		if err != nil {
			return def, err
		}

		ok, err := e.pred(v)
		if err != nil {
			return def, err
		}
		if ok {
			return v, nil
		}
	}
}
