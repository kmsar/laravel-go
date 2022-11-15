package Slices

// ArrayColumn returns the values from a single column in the input array
func ArrayColumn[T comparable](input []map[string]T, columnKey string) []T {
	columns := make([]T, 0, len(input))
	for _, val := range input {
		if v, ok := val[columnKey]; ok {
			columns = append(columns, v)
		}
	}
	return columns
}
