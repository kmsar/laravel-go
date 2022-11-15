package Slices

// Append appends the contents of t to the end of s and returns the result.
// If s has enough capacity, it is extended in place; otherwise a
// new array is allocated and returned.
func Append[T any](s []T, t ...T) []T {
	lens := len(s)
	tot := lens + len(t)
	if tot < 0 {
		panic("Append: cap out of range")
	}
	if tot > cap(s) {
		news := make([]T, tot, tot+tot/2)
		copy(news, s)
		s = news
	}
	s = s[:tot]
	copy(s[lens:], t)
	return s
}
