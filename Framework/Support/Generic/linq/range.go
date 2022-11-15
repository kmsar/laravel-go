package linq

type rangeEnumerator struct {
	st  int
	cnt int
	i   int
}

// Range generates a sequence of integral numbers within a specified range.
// e := linq.Range(3, 5)
// r, err := linq.ToSlice(e)
// exp := []int{3, 4, 5, 6, 7}
func Range(start, count int) Enumerable[int] {
	return func() Enumerator[int] {
		return &rangeEnumerator{st: start, cnt: count}
	}
}

func (e *rangeEnumerator) Next() (int, error) {
	if e.i >= e.cnt {
		return 0, EOC
	}

	r := e.st + e.i
	e.i++
	return r, nil
}
