package Slices

// ArrayKeyExists is alias of KeyExists()
func ArrayKeyExists[K, V comparable](k K, m map[K]V) bool {
	return KeyExists(k, m)
}
