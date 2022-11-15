package linq

// Any determines whether any element of a sequence satisfies a condition.
// isOdd := func(v int) (bool, error) {
// 	return v%2 != 0, nil
// }

// tests := []struct {
// 	src []int
// 	exp bool
// }{
// 	{[]int{2, 4, 6, 8, 10}, false},
// 	{[]int{2, 4, 6, 7, 10}, true},
// }

// for _, test := range tests {
// 	e := linq.FromSlice(test.src)
// 	r, err := linq.Any(e, isOdd)
// 	if err != nil {
// 		t.Fatalf("%v", err)
// 	}
// 	if r != test.exp {
// 		t.Fatalf("wants: %v, got %v", test.exp, r)
// 	}
// }
func Any[T any, E IEnumerable[T]](src E, pred func(v T) (bool, error)) (bool, error) {
	e := src()
	for {
		v, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return false, nil
			}
			return false, err
		}
		ok, err := pred(v)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}
}
