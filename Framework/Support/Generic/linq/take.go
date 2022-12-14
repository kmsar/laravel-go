package linq

type takeEnumerator[T any] struct {
	src Enumerator[T]
	cnt int
}

// Take returns a specified number of contiguous elements from the start of a sequence.
// tests := []struct {
// 	src linq.Enumerable[int]
// 	cnt int
// 	exp []int
// }{
// 	{linq.Range(0, 5), 3, []int{0, 1, 2}},
// 	{linq.Range(0, 5), 8, []int{0, 1, 2, 3, 4}},
// }
// for _, test := range tests {
// 	r, err := linq.ToSlice(
// 		linq.Take(test.src, test.cnt))
// 	if err != nil {
// 		t.Fatalf("%v: %v", test.cnt, err)
// 	}
// 	if !reflect.DeepEqual(r, test.exp) {
// 		t.Fatalf("%v, wants %v", r, test.exp)
// 	}
// }
func Take[T any, E IEnumerable[T]](src E, count int) Enumerable[T] {
	return func() Enumerator[T] {
		return &takeEnumerator[T]{src: src(), cnt: count}
	}
}

func (e *takeEnumerator[T]) Next() (def T, _ error) {
	if e.cnt <= 0 {
		return def, EOC
	}
	e.cnt--
	return e.src.Next()
}

type takeWhileEnumerator[T any] struct {
	src  Enumerator[T]
	pred func(T) (bool, error)
}

// TakeWhile returns elements from a sequence as long as a specified condition is true, and then skips the remaining elements.
// src := linq.Range(0, 10)
// r, err := linq.ToSlice(
// 	linq.TakeWhile(src, func(n int) (bool, error) { return n < 5, nil }))
// exp := []int{0, 1, 2, 3, 4}
func TakeWhile[T any, E IEnumerable[T]](src E, pred func(T) (bool, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &takeWhileEnumerator[T]{src: src(), pred: pred}
	}
}

func (e *takeWhileEnumerator[T]) Next() (def T, _ error) {
	v, err := e.src.Next()
	if err != nil {
		return def, err
	}
	ok, err := e.pred(v)
	if err != nil {
		return def, err
	}
	if !ok {
		return def, EOC
	}
	return v, nil
}

type takeLastEnumerator[T any] struct {
	src Enumerator[T]
	cnt int
	buf []T
	ofs int
	i   int
}

// TakeLast returns a new enumerable collection that contains the last count elements from source.
// tests := []struct {
// 	src linq.Enumerable[int]
// 	n   int
// 	exp []int
// }{
// 	{linq.Range(0, 10), 5, []int{5, 6, 7, 8, 9}},
// 	{linq.Range(0, 3), 5, []int{0, 1, 2}},
// }
// for i, test := range tests {
// 	r, err := linq.ToSlice(
// 		linq.TakeLast(test.src, test.n))
// 	if err != nil {
// 		t.Fatalf("%v: %v", i, err)
// 	}
// 	if !reflect.DeepEqual(r, test.exp) {
// 		t.Fatalf("%v, wants %v", r, test.exp)
// 	}
// }
func TakeLast[T any, E IEnumerable[T]](src E, count int) Enumerable[T] {
	return func() Enumerator[T] {
		return &takeLastEnumerator[T]{src: src(), cnt: count}
	}
}

func (e *takeLastEnumerator[T]) Next() (def T, _ error) {
	if e.buf == nil {
		e.buf = make([]T, e.cnt)
		i := 0
		for ; ; i++ {
			v, err := e.src.Next()
			if err != nil {
				if isEOC(err) {
					break
				}
				return def, err
			}
			e.buf[i%e.cnt] = v
		}
		if i < e.cnt {
			e.buf = e.buf[:i]
		} else {
			e.ofs = i
		}
	}
	i := e.i
	if i >= len(e.buf) {
		return def, EOC
	}
	e.i++
	return e.buf[(e.ofs+i)%len(e.buf)], nil
}
