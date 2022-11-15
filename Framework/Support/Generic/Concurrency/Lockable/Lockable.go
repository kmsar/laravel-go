package Lockable

import "sync"

// A Lockable is a value that may be safely simultaneously accessed
// from multiple goroutines via the Get and Set methods.
type Lockable[T any] struct {
	T
	mu sync.Mutex
}

// Get returns the value stored in a Lockable.
func (l *Lockable[T]) Get() T {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.T
}

// Set sets the value in a Lockable.
func (l *Lockable[T]) Set(v T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.T = v
}
