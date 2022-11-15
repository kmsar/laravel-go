package mapset

// Set is the primary interface provided by the mapset package.  It
// represents an unordered set of data and a large number of
// operations that can be applied to that set.
type Set[T comparable] interface {
	// Add Adds an element to the set. Returns whether
	// the item was added.
	Add(val T) bool

	// Cardinality Returns the number of elements in the set.
	Cardinality() int

	// Clear Removes all elements from the set, leaving
	// the empty set.
	Clear()

	// Clone Returns a clone of the set using the same
	// implementation, duplicating all keys.
	Clone() Set[T]

	// Contains Returns whether the given items
	// are all in the set.
	Contains(val ...T) bool

	// Difference Returns the difference between this set
	// and other. The returned set will contain
	// all elements of this set that are not also
	// elements of other.
	//
	// Note that the argument to Difference
	// must be of the same type as the receiver
	// of the method. Otherwise, Difference will
	// panic.
	Difference(other Set[T]) Set[T]

	// Equal Determines if two sets are equal to each
	// other. If they have the same cardinality
	// and contain the same elements, they are
	// considered equal. The order in which
	// the elements were added is irrelevant.
	//
	// Note that the argument to Equal must be
	// of the same type as the receiver of the
	// method. Otherwise, Equal will panic.
	Equal(other Set[T]) bool

	// Intersect Returns a new set containing only the elements
	// that exist only in both sets.
	//
	// Note that the argument to Intersect
	// must be of the same type as the receiver
	// of the method. Otherwise, Intersect will
	// panic.
	Intersect(other Set[T]) Set[T]

	// IsProperSubset Determines if every element in this set is in
	// the other set but the two sets are not equal.
	//
	// Note that the argument to IsProperSubset
	// must be of the same type as the receiver
	// of the method. Otherwise, IsProperSubset
	// will panic.
	IsProperSubset(other Set[T]) bool

	// IsProperSuperset Determines if every element in the other set
	// is in this set but the two sets are not
	// equal.
	//
	// Note that the argument to IsSuperset
	// must be of the same type as the receiver
	// of the method. Otherwise, IsSuperset will
	// panic.
	IsProperSuperset(other Set[T]) bool

	// IsSubset Determines if every element in this set is in
	// the other set.
	//
	// Note that the argument to IsSubset
	// must be of the same type as the receiver
	// of the method. Otherwise, IsSubset will
	// panic.
	IsSubset(other Set[T]) bool

	// IsSuperset Determines if every element in the other set
	// is in this set.
	//
	// Note that the argument to IsSuperset
	// must be of the same type as the receiver
	// of the method. Otherwise, IsSuperset will
	// panic.
	IsSuperset(other Set[T]) bool

	// Each Iterates over elements and executes the passed func against each element.
	// If passed func returns true, stop iteration at the time.
	Each(func(T) bool)

	// Iter Returns a channel of elements that you can
	// range over.
	Iter() <-chan T

	// Iterator Returns an Iterator object that you can
	// use to range over the set.
	Iterator() *Iterator[T]

	// Remove a single element from the set.
	Remove(i T)

	// String Provides a convenient string representation
	// of the current state of the set.
	String() string

	// SymmetricDifference Returns a new set with all elements which are
	// in either this set or the other set but not in both.
	//
	// Note that the argument to SymmetricDifference
	// must be of the same type as the receiver
	// of the method. Otherwise, SymmetricDifference
	// will panic.
	SymmetricDifference(other Set[T]) Set[T]

	// Union Returns a new set with all elements in both sets.
	//
	// Note that the argument to Union must be of the
	// same type as the receiver of the method.
	// Otherwise, IsSuperset will panic.
	Union(other Set[T]) Set[T]

	// Pop removes and returns an arbitrary item from the set.
	Pop() (T, bool)

	// ToSlice Returns the members of the set as a slice.
	ToSlice() []T

	// MarshalJSON will marshal the set into a JSON-based representation.
	MarshalJSON() ([]byte, error)

	// UnmarshalJSON will unmarshal a JSON-based byte slice into a full Set datastructure.
	// For this to work, set subtypes must implemented the Marshal/Unmarshal interface.
	UnmarshalJSON(b []byte) error
}

// NewSet creates and returns a new set with the given elements.
// Operations on the resulting set are thread-safe.
func NewSet[T comparable](vals ...T) Set[T] {
	s := newThreadSafeSet[T]()
	for _, item := range vals {
		s.Add(item)
	}
	return &s
}

// NewThreadUnsafeSet creates and returns a new set with the given elements.
// Operations on the resulting set are not thread-safe.
func NewThreadUnsafeSet[T comparable](vals ...T) Set[T] {
	s := newThreadUnsafeSet[T]()
	for _, item := range vals {
		s.Add(item)
	}
	return &s
}
