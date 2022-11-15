package Maps

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Generic/Chans"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Generic/Iterator"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Generic/KeyValue"
)

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		items: map[K]V{},
	}
}

func NewStringStringArray() *Map[string, string] {
	return NewMap[string, string]()
}

type Map[K comparable, V any] struct {
	items map[K]V
}

// Count The built-in len function returns on the number of items in a map:
func (m *Map[K, V]) Count() int {
	return len(m.items)
}

// Get will return the value associated with the key.
// If the key doesn't exist, the second return value will be false.
func (m *Map[K, V]) Get(key K) (V, bool) {
	val, exists := m.items[key]
	return val, exists
}

// Set will store a key-value pair. If the key already exists,
// it will overwrite the existing key-value pair.
func (m *Map[K, V]) Set(key K, val V) {
	m.items[key] = val
}

// Delete will remove the key and its associated value.
func (m *Map[K, V]) Delete(key K) {
	delete(m.items, key)
}

func (m *Map[K, V]) Has(key K) bool {
	_, ok := m.items[key]
	return ok
}

// Iterator returns an iterator that does an in-order traversal of the map.
func (m *Map[K, V]) Iterator() *Iterator.Iterator[K, V] {
	type kv = KeyValue.KeyValue[K, V] // convenient shorthand
	sender, receiver := Chans.Ranger[kv]()
	go func() {
		for key, value := range m.items {
			sender.Send(kv{key, value})
		}
		sender.Close()
	}()

	return &Iterator.Iterator[K, V]{receiver}
}
