package linq

import "golang.org/x/exp/constraints"

// Min returns the minimum value in a sequence of values.
// src := []byte{
// 	189, 47, 155, 170, 155, 99, 136, 3, 161, 231,
// }

// r, err := linq.Min(linq.FromSlice(src))
// exp := byte(3)
func Min[T constraints.Ordered, E IEnumerable[T]](src E) (def T, _ error) {
	e := src()
	min, err := e.Next()
	if err != nil {
		if isEOC(err) {
			err = InvalidOperation
		}
		return def, err
	}
	for {
		v, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return min, nil
			}
			return def, err
		}
		if err != nil {
			return def, err
		}
		if v < min {
			min = v
		}
	}
}

// MinBy returns the minimum value in a generic sequence according to a specified key selector function.
// src := []byte{
// 	189, 47, 155, 170, 155, 99, 136, 3, 161, 231,
// }

// r, err = linq.MinBy(linq.FromSlice(src), func(v byte) (float64, error) {
// 	return math.Abs(float64(v) - 100), nil
// })
// exp = byte(99)


// r, err = linq.MinByFunc(linq.FromSlice(src), func(a, b byte) (bool, error) {
// 	if a%2 != b%2 {
// 		return a%2 == 0, nil
// 	}
// 	return a < b, nil
// })
// exp = byte(136)
func MinBy[T any, K constraints.Ordered, E IEnumerable[T]](src E, keySelector func(T) (K, error)) (def T, _ error) {
	e := src()
	min, err := e.Next()
	if err != nil {
		if isEOC(err) {
			err = InvalidOperation
		}
		return def, err
	}
	mink, err := keySelector(min)
	if err != nil {
		return def, err
	}
	for {
		v, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return min, nil
			}
			return def, err
		}
		k, err := keySelector(v)
		if err != nil {
			return def, err
		}
		if k < mink {
			min, mink = v, k
		}
	}
}

// MinByFunc returns the minimum value in a generic sequence according to a comparer function.
func MinByFunc[T any, E IEnumerable[T]](src E, less func(a, b T) (bool, error)) (def T, _ error) {
	e := src()
	min, err := e.Next()
	if err != nil {
		if isEOC(err) {
			err = InvalidOperation
		}
		return def, err
	}
	for {
		v, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return min, nil
			}
			return def, err
		}
		b, err := less(v, min)
		if err != nil {
			return def, err
		}
		if b {
			min = v
		}
	}
}
