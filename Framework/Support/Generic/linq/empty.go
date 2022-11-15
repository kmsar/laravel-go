package linq

type emptyEnumerator[T any] struct{}

// Empty returns an empty IEnumerable[T] that has the specified type argument.
// r, err := linq.ToSlice(linq.Empty[uint]())
// if err != nil {
// 	t.Fatalf("%v", err)
// }
// if len(r) != 0 {
// 	t.Fatalf("not empty: %#v", r)
// }
// if reflect.TypeOf(r) != reflect.TypeOf([]uint{}) {
// 	t.Fatalf("%T != []uint", r)
// }
func Empty[T any]() Enumerable[T] {
	return func() Enumerator[T] {
		return emptyEnumerator[T]{}
	}
}

func (emptyEnumerator[T]) Next() (def T, _ error) {
	return def, EOC
}
