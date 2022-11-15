package linq

type chunkEnumerator[T any] struct {
	src  Enumerator[T]
	size int
}

// Chunk splits the elements of a sequence into chunks of size at most `size`.

// src := []int{1, 2, 3, 4, 5, 6, 7, 8}
// e := linq.Chunk(linq.FromSlice(src), 3)
// r, err := linq.ToSlice(e)
// if err != nil {
// 	t.Fatalf("%v", err)
// }

// exp := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}

// if !reflect.DeepEqual(r, exp) {
// 	t.Fatalf("%v, wants %v", r, exp)
// }
func Chunk[T any, E IEnumerable[T]](src E, size int) Enumerable[[]T] {
	return func() Enumerator[[]T] {
		return &chunkEnumerator[T]{
			src:  src(),
			size: size,
		}
	}
}

func (e *chunkEnumerator[T]) Next() ([]T, error) {
	s := make([]T, 0, e.size)

	for i := 0; i < e.size; i++ {
		v, err := e.src.Next()
		if err != nil {
			if isEOC(err) {
				break
			}
			return nil, err
		}
		s = append(s, v)
	}

	if len(s) == 0 {
		return nil, EOC
	}

	return s, nil
}
