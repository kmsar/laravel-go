package linq

type repeatEnumerator[T any] struct {
	v T
	n int
}

// Repeat generates a sequence that contains one repeated value.
// r, err := linq.ToSlice(linq.Repeat(rune('A'), 5))
// exp := []rune{'A', 'A', 'A', 'A', 'A'}
func Repeat[T any](v T, n int) Enumerable[T] {
	return func() Enumerator[T] {
		return &repeatEnumerator[T]{v: v, n: n}
	}
}

func (e *repeatEnumerator[T]) Next() (def T, _ error) {
	if e.n <= 0 {
		return def, EOC
	}
	e.n--
	return e.v, nil
}
