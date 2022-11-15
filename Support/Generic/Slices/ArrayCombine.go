package Slices

// ArrayCombine creates an array by using one array for keys and another for its values
func ArrayCombine[K, V comparable](keys []K, values []V) map[K]V {
	if len(keys) != len(values) {
		return nil
	}
	m := make(map[K]V, len(keys))
	for i, v := range keys {
		m[v] = values[i]
	}
	return m
}
