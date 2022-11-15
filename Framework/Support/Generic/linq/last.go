package linq

// Last returns the last element of a sequence that satisfies a specified condition.
// src := []int{1, 2, 3, 4, 5, 6, 7}
// r, err := linq.Last(
// 	linq.Where(linq.FromSlice(src),
// 		func(v int) (bool, error) {
// 			return v%2 == 0, nil
// 		}))
// exp := 6
func Last[T any, E IEnumerable[T]](src E) (def T, _ error) {
	var last T
	var eocerr error = InvalidOperation
	e := src()
	for {
		v, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return last, eocerr
			}
			return def, err
		}
		last = v
		eocerr = nil
	}
}

// LastOrDefault returns the last element of a sequence that satisfies a condition, or a specified default value if no such element is found.
// src := []int{1, 2, 3}
// def := 42
// r, err := linq.LastOrDefault(linq.Empty[int](), def)
// exp := def

// r, err = linq.LastOrDefault(linq.FromSlice(src), def)
// exp = 3
func LastOrDefault[T any, E IEnumerable[T]](src E, defaultValue T) (T, error) {
	v, err := Last(src)
	if isInvalidOperation(err) {
		return defaultValue, nil
	}
	return v, err
}
