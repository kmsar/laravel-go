package KeyValue

// KeyValue is a pair of key and value used when iterating.
type KeyValue[K, V any] struct {
	K K
	V V
}
