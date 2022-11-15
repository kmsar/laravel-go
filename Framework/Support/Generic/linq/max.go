package linq

import "golang.org/x/exp/constraints"

// Maxby returns the maximum value in a sequence of values.
// src := []byte{
// 	189, 47, 155, 170, 155, 99, 136, 3, 161, 231,
// }
// r, err := linq.Max(linq.FromSlice(src))
// exp := byte(231)
func Max[T constraints.Ordered, E IEnumerable[T]](src E) (def T, _ error) {
	e := src()
	max, err := e.Next()
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
				return max, nil
			}
			return def, err
		}
		if err != nil {
			return def, err
		}
		if v > max {
			max = v
		}
	}
}

// MaxBy returns the maximum value in a generic sequence according to a specified key selector function.
// src := []byte{
// 	189, 47, 155, 170, 155, 99, 136, 3, 161, 231,
// }
// r, err = linq.MaxBy(linq.FromSlice(src), func(v byte) (float64, error) {
// 	return math.Abs(float64(v) - 150), nil
// })
// exp = byte(3)

// r, err = linq.MaxByFunc(linq.FromSlice(src), func(a, b byte) (bool, error) {
// 	if a%2 != b%2 {
// 		return a%2 == 0, nil
// 	}
// 	return a > b, nil
// })
// exp = byte(170)
func MaxBy[T any, K constraints.Ordered, E IEnumerable[T]](src E, keySelector func(T) (K, error)) (def T, _ error) {
	e := src()
	max, err := e.Next()
	if err != nil {
		if isEOC(err) {
			err = InvalidOperation
		}
		return def, err
	}
	maxk, err := keySelector(max)
	if err != nil {
		return def, err
	}
	for {
		v, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return max, nil
			}
			return def, err
		}
		k, err := keySelector(v)
		if err != nil {
			return def, err
		}
		if k > maxk {
			max, maxk = v, k
		}
	}
}

// MaxByFunc returns the maximum value in a generic sequence according to a comparer function.
func MaxByFunc[T any, E IEnumerable[T]](src E, greater func(a, b T) (bool, error)) (def T, _ error) {
	e := src()
	max, err := e.Next()
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
				return max, nil
			}
			return def, err
		}
		b, err := greater(v, max)
		if err != nil {
			return def, err
		}
		if b {
			max = v
		}
	}
}
