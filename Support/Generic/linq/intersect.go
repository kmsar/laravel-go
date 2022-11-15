package linq

type intersectEnumerator[T any, H comparable] struct {
	fst  Enumerator[T]
	snd  Enumerator[T]
	eq   func(T, T) (bool, error)
	hash func(T) (H, error)
	hmap *hashMap[H, T]
}

// Intersect produces the set intersection of two sequences by using the specified comparer functions.
// fst := linq.FromSlice([]int{1, 2, 3, 4, 5, 6, 7})
// snd := linq.FromSlice([]int{0, 2, 4, 6, 8, 10})

// e := linq.Intersect(
// 	fst, snd,
// 	func(a, b int) (bool, error) { return a == b, nil },
// 	func(a int) (int, error) { return a / 3, nil })

// r, err := linq.ToSlice(e)
// exp := []int{2, 4, 6}
func Intersect[T any, E IEnumerable[T]](first, second E, equals func(T, T) (bool, error), getHashCode func(T) (int, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &intersectEnumerator[T, int]{
			fst:  first(),
			snd:  second(),
			eq:   equals,
			hash: getHashCode,
		}
	}
}

// IntersectBy produces the set intersection of two sequences according to a specified key selector function.
// type PlanetType int
// const (
// 	Rock PlanetType = iota
// 	Ice
// 	Gas
// 	Liquid
// )
// type Planet struct {
// 	Name         string
// 	Type         PlanetType
// 	OrderFromSun int
// }
// p1 := []Planet{
// 	{"Marcury", Rock, 1},
// 	{"Venus", Rock, 2},
// 	{"Earth", Rock, 3},
// 	{"Jupiter", Gas, 5},
// }
// p2 := []Planet{
// 	{"Marcury", Rock, 1},
// 	{"Earth", Rock, 3},
// 	{"Mars", Rock, 4},
// 	{"Jupiter", Gas, 5},
// }

// e := linq.IntersectBy(
// 	linq.FromSlice(p1),
// 	linq.FromSlice(p2),
// 	func(p Planet) (string, error) { return p.Name, nil })

// r, err := linq.ToSlice(e)
// exp := []Planet{
// 	{"Marcury", Rock, 1},
// 	{"Earth", Rock, 3},
// 	{"Jupiter", Gas, 5},
// }
func IntersectBy[T any, K comparable, E IEnumerable[T]](first, second E, keySelector func(v T) (K, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &intersectEnumerator[T, K]{
			fst:  first(),
			snd:  second(),
			eq:   alwaysEqual[T],
			hash: keySelector,
		}
	}
}

func (e *intersectEnumerator[T, H]) Next() (def T, _ error) {
	if e.hmap == nil {
		hm := newHashMap(e.hash, e.eq)
		if err := hm.addAll(e.snd); err != nil {
			return def, err
		}
		e.hmap = hm
	}

	for {
		v, err := e.fst.Next()
		if err != nil {
			return def, err
		}
		has, err := e.hmap.has(v)
		if err != nil {
			return def, err
		}
		if has {
			return v, nil
		}
	}
}
