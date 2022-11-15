package linq

// All determines whether all elements of a sequence satisfy a condition.
// isEven := func(v int) (bool, error) {
// 	return v%2 == 0, nil
// }

// tests := []struct {
// 	src []int
// 	exp bool
// }{
// 	{[]int{2, 4, 6, 8, 10}, true},
// 	{[]int{2, 4, 6, 7, 10}, false},
// }

// for _, test := range tests {
// 	e := linq.FromSlice(test.src)
// 	r, err := linq.All(e, isEven)
// 	if err != nil {
// 		t.Fatalf("%v", err)
// 	}
// 	if r != test.exp {
// 		t.Fatalf("wants: %v, got %v", test.exp, r)
// 	}
// }
func All[T any, E IEnumerable[T]](src E, pred func(v T) (bool, error)) (bool, error) {
	e := src()
	for {
		v, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return true, nil
			}
			return false, err
		}
		ok, err := pred(v)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
}
