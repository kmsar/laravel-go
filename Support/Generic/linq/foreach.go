package linq

// ForEach performs the specified function on each element of the specified enumerator.
//	c := 0
	// s := 0
	// err := linq.ForEach(
	// 	linq.FromSlice([]int{1, 2, 3, 4, 5}),
	// 	func(n int) error {
	// 		c++
	// 		s += n
	// 		return nil
	// 	})
	// if c != 5 || s != 15 {
	// 	t.Fatalf("(c, s) = (%v, %v) wants (%v, %v)", c, s, 5, 15)
	// }
func ForEach[T any, E IEnumerable[T]](src E, f func(T) error) error {
	e := src()
	for {
		v, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return nil
			}
			return err
		}
		err = f(v)
		if err != nil {
			return err
		}
	}
}
