package linq

// First returns the first element in a sequence
// r, err := linq.First(
// 	linq.Where(linq.FromSlice(src),
// 		func(v int) (bool, error) {
// 			return (v > 5 && v%3 == 0), nil
// 		}))
// exp := 6
func First[T any, E IEnumerable[T]](src E) (def T, _ error) {
	e := src()
	v, err := e.Next()
	if err != nil {
		if isEOC(err) {
			err = InvalidOperation
		}
		return def, err
	}
	return v, nil
}

// FirstOrDefault returns the first element of the sequence, or a specified default value if no such element is found.
// src := []int{1, 2, 3}
// def := 42
// r, err := linq.FirstOrDefault(linq.Empty[int](), def)
// exp := def

// r, err = linq.FirstOrDefault(linq.FromSlice(src), def)
// exp = 1
func FirstOrDefault[T any, E IEnumerable[T]](src E, defaultValue T) (T, error) {
	v, err := First(src)
	if err != nil {
		if isInvalidOperation(err) {
			err = nil
		}
		return defaultValue, err
	}
	return v, nil
}
