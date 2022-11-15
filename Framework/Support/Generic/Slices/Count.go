package Slices

import "reflect"

// Count counts all elements in an array or map
func Count(v any) int {
	if v == nil {
		return 0
	}
	return reflect.ValueOf(v).Len()
}
