package Slices

// ArrayPad pads array to the specified length with a value
func ArrayPad[T comparable](array []T, size int, value T) []T {
	if size == 0 || (size > 0 && size < len(array)) || (size < 0 && size > -len(array)) {
		return array
	}
	n := size
	if size < 0 {
		n = -size
	}
	n -= len(array)
	tmp := make([]T, n)
	for i := 0; i < n; i++ {
		tmp[i] = value
	}
	if size > 0 {
		return append(array, tmp...)
	}
	return append(tmp, array...)
}
