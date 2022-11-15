package Iterator

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Generic/Chans"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Generic/KeyValue"
)

// Iterator is used to iterate over the map.
type Iterator[K, V any] struct {
	Res *Chans.Receiver[KeyValue.KeyValue[K, V]]
}

// Next returns the next key and value pair. The bool result reports
// whether the values are valid. If the values are not valid, we have
// reached the end.
func (it *Iterator[K, V]) Next() (K, V, bool) {
	kv, ok := it.Res.Next()
	return kv.K, kv.V, ok
}
