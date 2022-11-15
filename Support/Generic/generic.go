package Generic

import "golang.org/x/exp/constraints"

type Tree[T interface{}] struct {
	left, right *Tree[T]
	value       T
}

func (t *Tree[T]) Lookup(x T) *Tree[T] {
	return &Tree[T]{}
}

var stringTree Tree[string]

// Scale returns a copy of s with each element multiplied by c.
// This implementation has a problem, as we will see.
func Scale[E constraints.Integer](s []E, c E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

//x := GMin[int](2, 3)
//fmin := GMin[float64]
//m := fmin(2.71, 3.14)

//[S interface{~[]E}, E interface{}]
//Here S must be a slice type whose element type can be any type.
//Here S must be a slice type whose element type can be any type.
//
//Because this is a common case, the enclosing interface{} may be omitted for interfaces in constraint position, and we can simply write:
//
//[S ~[]E, E interface{}]
//[S ~[]E, E any]

// Scale1 The problem is that the Scale function returns a value of type []E where E is the element type of the argument slice. When we call Scale with a value of type Point, whose underlying type is []int32, we get back a value of type []int32, not type Point. This follows from the way that the generic code is written, but it’s not what we want.
//In order to fix this, we have to change the Scale function to use a type parameter for the slice type.
// Scale returns a copy of s with each element multiplied by c.
func Scale1[S ~[]E, E constraints.Integer](s S, c E) S {
	return Map[E, E](s, func(t1 E) E {
		return t1 * c
	})
	//var r S
	//r = make(S, len(s))
	//for i, v := range s {
	//	r[i] = (v * c)
	//}
	//return r
}

func Assertion[T any](v interface{}) (T, bool) {
	t, ok := v.(T)
	return t, ok
}

func Switch[T any](v interface{}) (T, bool) {
	switch v := v.(type) {
	case T:
		return v, true
	default:
		var zero T
		return zero, false
	}
}

func Switch2[T any](v interface{}) int {
	switch v.(type) {
	case T:
		return 0
	case string:
		return 1
	default:
		return 2
	}
}

//
//// S2a will be set to 0.
//var S2a = Switch2[string]("a string")
//
//// S2b will be set to 1.
//var S2b = Switch2[int]("another string")

//
//type Pair[T any] struct { f1, f2 T }
//var V = Pair{1, 2} // inferred as Pair[int]{1, 2}
