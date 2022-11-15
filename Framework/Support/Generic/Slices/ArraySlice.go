package Slices

// ArraySlice extracts a slice of the array
func ArraySlice[T comparable](array []T, offset, length uint) []T {
	if offset > uint(len(array)) {
		return nil
	}
	end := offset + length
	if end < uint(len(array)) {
		return array[offset:end]
	}
	return array[offset:]
}
