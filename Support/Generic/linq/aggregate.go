package linq

// Aggregate applies an accumulator function over a sequence.
//	src := []byte{100, 100, 200, 200}
	// r, err := linq.Aggregate(
	// 	linq.FromSlice(src),
	// 	0,
	// 	func(acc int, v byte) (int, error) {
	// 		return acc + int(v), nil
	// 	})
	// if err != nil {
	// 	t.Fatalf("%v", err)
	// }
	// exp := 600
	// if r != exp {
	// 	t.Fatalf("%v, wants %v", r, exp)
	// }
func Aggregate[S, T any, E IEnumerable[S]](src E, seed T, fn func(acc T, v S) (T, error)) (T, error) {
	acc := seed
	err := ForEach(src, func(v S) (err error) {
		acc, err = fn(acc, v)
		return err
	})
	return acc, err
}
