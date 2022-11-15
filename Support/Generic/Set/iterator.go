package mapset

// Iterator defines an iterator over a Set, its C channel can be used to range over the Set's
// elements.
type Iterator[T comparable] struct {
	C    <-chan T
	stop chan struct{}
}

// Stop stops the Iterator, no further elements will be received on C, C will be closed.
func (i *Iterator[T]) Stop() {
	// Allows for Stop() to be called multiple times
	// (close() panics when called on already closed channel)
	defer func() {
		recover()
	}()

	close(i.stop)

	// Exhaust any remaining elements.
	for range i.C {
	}
}

// newIterator returns a new Iterator instance together with its item and stop channels.
func newIterator[T comparable]() (*Iterator[T], chan<- T, <-chan struct{}) {
	itemChan := make(chan T)
	stopChan := make(chan struct{})
	return &Iterator[T]{
		C:    itemChan,
		stop: stopChan,
	}, itemChan, stopChan
}
