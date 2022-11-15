package linq

type skipEnumerator[T any] struct {
	src Enumerator[T]
	cnt int
}

// Skip bypasses a specified number of elements in a sequence and then returns the remaining elements.
// tests := []struct {
// 	src linq.Enumerable[int]
// 	n   int
// 	exp []int
// }{
// 	{linq.Range(0, 5), 0, []int{0, 1, 2, 3, 4}},
// 	{linq.Range(0, 5), 2, []int{2, 3, 4}},
// 	{linq.Range(0, 5), 5, []int{}},
// 	{linq.Range(0, 5), 6, []int{}},
// }
// for _, test := range tests {
// 	e := linq.Skip(test.src, test.n)
// 	r, err := linq.ToSlice(e)
// 	if err != nil {
// 		t.Fatalf("Skip(%v): %v", test.n, err)
// 	}
// 	if !reflect.DeepEqual(r, test.exp) {
// 		t.Fatalf("Skip(%v): %v, wants %v", test.n, r, test.exp)

func Skip[T any, E IEnumerable[T]](src E, count int) Enumerable[T] {
	return func() Enumerator[T] {
		return &skipEnumerator[T]{src: src(), cnt: count}
	}
}

func (e *skipEnumerator[T]) Next() (def T, _ error) {
	for ; e.cnt > 0; e.cnt-- {
		_, err := e.src.Next()
		if err != nil {
			return def, err
		}
	}
	return e.src.Next()
}

type skipWhileEnumerator[T any] struct {
	src     Enumerator[T]
	pred    func(T) (bool, error)
	skipped bool
}

// SkipWhile bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements.
// src := linq.Range(0, 10)
// r, err := linq.ToSlice(
// 	linq.SkipWhile(src, func(n int) (bool, error) { return n < 5, nil }))
// exp := []int{5, 6, 7, 8, 9}
func SkipWhile[T any, E IEnumerable[T]](src E, pred func(T) (bool, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &skipWhileEnumerator[T]{src: src(), pred: pred}
	}
}

func (e *skipWhileEnumerator[T]) Next() (def T, _ error) {
	if e.skipped {
		return e.src.Next()
	}
	for {
		v, err := e.src.Next()
		if err != nil {
			return def, err
		}
		ok, err := e.pred(v)
		if err != nil {
			return def, err
		}
		if !ok {
			e.skipped = true
			return v, nil
		}
	}
}

type skipLastEnumerator[T any] struct {
	src Enumerator[T]
	cnt int
	buf []T
	i   int
}

// SkipLast returns a new enumerable collection that contains the elements from source with the last count elements of the source collection omitted.
// src := linq.Range(0, 10)
// r, err := linq.ToSlice(
// 	linq.SkipLast(src, 3))
// exp := []int{0, 1, 2, 3, 4, 5, 6}
func SkipLast[T any, E IEnumerable[T]](src E, count int) Enumerable[T] {
	return func() Enumerator[T] {
		return &skipLastEnumerator[T]{src: src(), cnt: count}
	}
}

func (e *skipLastEnumerator[T]) Next() (def T, _ error) {
	if e.buf == nil {
		e.buf = make([]T, e.cnt)
		for i := 0; i < e.cnt; i++ {
			v, err := e.src.Next()
			if err != nil {
				return def, err
			}
			e.buf[i] = v
		}
	}

	i := e.i % e.cnt
	r := e.buf[i]
	v, err := e.src.Next()
	if err != nil {
		return def, err
	}
	e.buf[i] = v
	e.i++
	return r, nil
}
