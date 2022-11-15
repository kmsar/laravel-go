package linq

import "golang.org/x/exp/constraints"

// Sum calculates the sum of a integer sequence.
// src := []int16{100, 200, 300, 400, 500}
// r, err := linq.Sum(linq.FromSlice(src))
// if err != nil {
// 	t.Fatalf("%v", err)
// }
// exp := 1500
func Sum[T constraints.Integer, E IEnumerable[T]](src E) (int, error) {
	sum := 0
	err := ForEach(src, func(v T) error {
		sum += int(v)
		return nil
	})
	if err != nil {
		return 0, err
	}
	return sum, nil
}

// Sumf calculates the sum of a floating number sequence.
// linq.Select(linq.FromSlice(src), func(v int16) (float32, error) { return float32(v) / 16, nil }))
// if err != nil {
// 	t.Fatalf("%v", err)
// }
// expf := float64(1500) / 16
func Sumf[T constraints.Float, E IEnumerable[T]](src E) (float64, error) {
	var sum float64
	err := ForEach(src, func(v T) error {
		sum += float64(v)
		return nil
	})
	if err != nil {
		return 0, err
	}
	return sum, nil
}
