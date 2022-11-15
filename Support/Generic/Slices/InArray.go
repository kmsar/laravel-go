package Slices

// InArray checks if a value exists in an array
//
// needle is the element to search, haystack is the slice to be searched
func InArray[T comparable](needle T, haystack []T) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}

	return false
}
