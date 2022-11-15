package linq

type unionEnumerator[T any, H comparable] struct {
	fst  Enumerator[T]
	snd  Enumerator[T]
	hmap *hashMap[H, T]
}

// Union produces the set union of two sequences by using the specified comparer functions.
// fst := linq.FromSlice([]int{2, 3, 4, 5})
// snd := linq.FromSlice([]int{0, 2, 4, 6, 8})

// e := linq.Union(
// 	fst, snd,
// 	func(a, b int) (bool, error) { return a == b, nil },
// 	func(a int) (int, error) { return a / 3, nil })

// r, err := linq.ToSlice(e)

// exp := []int{2, 3, 4, 5, 0, 6, 8}
func Union[T any, E IEnumerable[T]](first, second E, equals func(T, T) (bool, error), getHashCode func(T) (int, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &unionEnumerator[T, int]{
			fst:  first(),
			snd:  second(),
			hmap: newHashMap(getHashCode, equals),
		}
	}
}

// UnionBy produces the set union of two sequences according to a specified key selector function.
// type Planet struct {
// 	Name         string
// 	OrderFromSun int
// }
// fst := []Planet{
// 	{"Marcury", 1},
// 	{"Venus", 2},
// 	{"Earth", 3},
// 	{"Mars", 4},
// 	{"Jupiter", 5},
// }
// snd := []Planet{
// 	{"Mars", 4},
// 	{"Jupiter", 5},
// 	{"Saturn", 6},
// 	{"Uranus", 7},
// 	{"Neptune", 8},
// }

// e := linq.UnionBy(
// 	linq.FromSlice(fst), linq.FromSlice(snd),
// 	func(v Planet) (int, error) { return v.OrderFromSun, nil })

// r, err := linq.ToSlice(e)
// if err != nil {
// 	t.Fatalf("%v", err)
// }
// exp := []Planet{
// 	{"Marcury", 1},
// 	{"Venus", 2},
// 	{"Earth", 3},
// 	{"Mars", 4},
// 	{"Jupiter", 5},
// 	{"Saturn", 6},
// 	{"Uranus", 7},
// 	{"Neptune", 8},
// }
func UnionBy[T any, K comparable, E IEnumerable[T]](first, second E, keySelector func(v T) (K, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &unionEnumerator[T, K]{
			fst:  first(),
			snd:  second(),
			hmap: newKeyMap(keySelector),
		}
	}
}

func (e *unionEnumerator[T, H]) Next() (def T, _ error) {
	if e.fst != nil {
		v, err := e.fst.Next()
		if err == nil {
			if _, err = e.hmap.add(v); err != nil {
				return def, err
			}
			return v, nil
		}
		if !isEOC(err) {
			return def, err
		}
		e.fst = nil
	}

	for {
		v, err := e.snd.Next()
		if err != nil {
			return def, err
		}
		has, err := e.hmap.has(v)
		if err != nil {
			return def, err
		}
		if !has {
			return v, nil
		}
	}
}
