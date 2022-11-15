package linq

type selectEnumerator[S, T any] struct {
	src Enumerator[S]
	sel func(v S) (T, error)
}

// Select projects each element of a sequence into a new form.
// e1 := linq.FromSlice([]string{
// 	"a", "bb", "ccc", "dddd", "eeeee",
// })
// e2 := linq.Select(e1, func(s string) (int, error) { return len(s), nil })
// r, err := linq.ToSlice(e2)

// exp := []int{1, 2, 3, 4, 5}

func Select[S, T any, E IEnumerable[S]](src E, selector func(v S) (T, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &selectEnumerator[S, T]{src: src(), sel: selector}
	}
}

func (e *selectEnumerator[S, T]) Next() (def T, _ error) {
	v, err := e.src.Next()
	if err != nil {
		return def, err
	}

	return e.sel(v)
}
