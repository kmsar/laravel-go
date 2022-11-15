package linq

type concatEnumerator[T any] struct {
	fst Enumerator[T]
	snd Enumerable[T]
}

// Concat concatenates two sequences.
// e1 := linq.FromSlice([]int{4, 5, 6})
// e2 := linq.FromSlice([]int{1, 2, 3})

// r, err := linq.ToSlice(
// 	linq.Concat(e1, e2))
// if err != nil {
// 	t.Fatalf("%v", err)
// }
// exp := []int{4, 5, 6, 1, 2, 3}
// if !reflect.DeepEqual(r, exp) {
// 	t.Fatalf("%v, wants %v", r, exp)
// }
func Concat[T any, E IEnumerable[T]](first, second E) Enumerable[T] {
	return func() Enumerator[T] {
		return &concatEnumerator[T]{fst: first(), snd: Enumerable[T](second)}
	}
}

func (e *concatEnumerator[T]) Next() (def T, _ error) {
	v, err := e.fst.Next()
	if err == nil {
		return v, nil
	}
	if !isEOC(err) {
		return def, err
	}
	if e.snd == nil {
		return def, EOC
	}
	e.fst, e.snd = e.snd(), nil
	return e.Next()
}
