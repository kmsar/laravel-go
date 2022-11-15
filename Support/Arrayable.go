package Support

// Offset In Go, type constraints must be interfaces.
//https://go.dev/blog/intro-generics
//Until recently, the Go spec said that an interface defines a method set,
//which is roughly the set of methods enumerated in the interface.
//Any type that implements all those methods implements that interface.
type Offset interface {
	int64 | int | ~string
}

type ArrayAccess interface {
	offsetExists(offset Offset) bool
	offsetGet(offset Offset) interface{}
	offsetSet(offset Offset, value interface{})
	offsetUnset(offset Offset)
}
