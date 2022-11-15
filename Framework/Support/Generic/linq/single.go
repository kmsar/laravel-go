package linq

// Single returns the only element of a sequence, and return an error InvalidOperation if more than one such element exists.
// src := linq.FromSlice([]int{1, 2, 3})

// r, err := linq.Single(linq.Where(src,
// 	func(v int) (bool, error) {
// 		return v%2 == 0, nil
// 	}))

// exp := 2
// if r != exp {
// 	t.Fatalf("%v, wants %v", r, exp)
// }

// _, err = linq.Single(linq.Where(src,
// 	func(v int) (bool, error) {
// 		return v%2 != 0, nil
// 	}))
// if !errors.Is(err, linq.InvalidOperation)
func Single[T any, E IEnumerable[T]](src E) (def T, _ error) {
	e := src()
	v, err := e.Next()
	if err != nil {
		if isEOC(err) {
			err = InvalidOperation
		}
		return def, err
	}
	_, err = e.Next()
	if err == nil {
		return def, InvalidOperation
	}
	if !isEOC(err) {
		return def, err
	}
	return v, nil
}

// SingleOrDefault returns the only element of a sequence, or a specified default value if no such element exists; this function returns an error InvalidOperation if more than one element satisfies the condition.
// src := linq.FromSlice([]int{1, 2, 3})
// def := 42
// r, err := linq.SingleOrDefault(linq.Where(src,
// 	func(v int) (bool, error) { return v%2 == 0, nil }),
// 	def)
// exp := 2

// r, err = linq.SingleOrDefault(linq.Empty[int](), def)
// exp = def

// _, err = linq.SingleOrDefault(linq.Where(src,
// 	func(v int) (bool, error) { return v%2 != 0, nil }),
// 	def)
// errors.Is(err, linq.InvalidOperation
func SingleOrDefault[T any, E IEnumerable[T]](src E, defaultValue T) (def T, _ error) {
	e := src()
	v, err := e.Next()
	if err != nil {
		if isEOC(err) {
			return defaultValue, nil
		}
		return def, err
	}
	_, err = e.Next()
	if err == nil {
		return def, InvalidOperation
	}
	if !isEOC(err) {
		return def, err
	}
	return v, nil
}
