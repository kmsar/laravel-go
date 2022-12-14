package linq

type joinEnumerator[S1, S2, T any, K comparable] struct {
	eOut  Enumerator[S1]
	eIn   Enumerator[S2]
	ksOut func(S1) (K, error)
	ksIn  func(S2) (K, error)
	rSel  func(S1, S2) (T, error)

	s1  *S1
	ks1 K
	ms2 map[K][]S2
	i   int
}

// Join correlates the elements of two sequences based on matching keys.
// nums := linq.Range(3, 6)
// strs := linq.FromSlice(
// 	[]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"})

// e := linq.Join(nums, strs,
// 	func(n int) (int, error) { return n, nil },
// 	func(s string) (int, error) { return len(s), nil },
// 	func(n int, s string) (string, error) { return fmt.Sprintf("%d:%s", n, s), nil })

// r, err := linq.ToSlice(e)
// exp := []string{
// 	"3:one", "3:two", "3:six", "3:ten",
// 	"4:four", "4:five", "4:nine",
// 	"5:three", "5:seven", "5:eight",
// }
func Join[S1, S2, T any, K comparable, E1 IEnumerable[S1], E2 IEnumerable[S2]](
	outer E1,
	inner E2,
	outerKeySelector func(S1) (K, error),
	innerKeySelector func(S2) (K, error),
	resultSelector func(S1, S2) (T, error),
) Enumerable[T] {
	return func() Enumerator[T] {
		return &joinEnumerator[S1, S2, T, K]{
			eOut:  outer(),
			eIn:   inner(),
			ksOut: outerKeySelector,
			ksIn:  innerKeySelector,
			rSel:  resultSelector,
		}
	}
}

func (e *joinEnumerator[S1, S2, T, K]) Next() (def T, _ error) {
	if e.s1 == nil {
		s1, err := e.eOut.Next()
		if err != nil {
			return def, err
		}
		ks1, err := e.ksOut(s1)
		if err != nil {
			return def, err
		}
		e.s1 = &s1
		e.ks1 = ks1
	}

	if e.ms2 == nil {
		m, err := innerMap(e.eIn, e.ksIn)
		if err != nil {
			return def, err
		}
		e.ms2 = m
	}

	s := e.ms2[e.ks1]

	if e.i >= len(s) {
		e.i = 0
		e.s1 = nil
		return e.Next()
	}

	i := e.i
	e.i++

	return e.rSel(*e.s1, s[i])
}

func innerMap[S2 any, K comparable](e Enumerator[S2], ks func(S2) (K, error)) (map[K][]S2, error) {
	m := make(map[K][]S2)
	for {
		s2, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return m, nil
			}
			return nil, err
		}

		k, err := ks(s2)
		if err != nil {
			return nil, err
		}
		m[k] = append(m[k], s2)
	}
}
