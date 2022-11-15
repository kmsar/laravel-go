package Slices

// Copy copies values from t to s, stopping when either slice is
// full, returning the number of values copied.
func Copy[T any](s, t []T) int {
	i := 0
	for ; i < len(s) && i < len(t); i++ {
		s[i] = t[i]
	}
	return i
}
