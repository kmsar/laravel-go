package linq

type exceptEnumerator[T any, H comparable] struct {
	fst  Enumerator[T]
	snd  Enumerator[T]
	eq   func(T, T) (bool, error)
	hash func(T) (H, error)
	hmap *hashMap[H, T]
}

// Except produces the set difference of two sequences by using the specified comparer functions.
// e1 := linq.FromSlice([]string{"Mercury", "Venus", "Earth", "Jupiter"})
// e2 := linq.FromSlice([]string{"Mercury", "Earth", "Mars", "Jupiter"})

// e := linq.Except(e1, e2,
// 	func(a, b string) (bool, error) { return a == b, nil },
// 	func(a string) (int, error) { return len(a), nil })

// r, err := linq.ToSlice(e)
// exp := []string{"Venus"}
// if !reflect.DeepEqual(r, exp) 
func Except[T any, E IEnumerable[T]](first, second E, equals func(T, T) (bool, error), getHashCode func(T) (int, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &exceptEnumerator[T, int]{
			fst:  first(),
			snd:  second(),
			eq:   equals,
			hash: getHashCode,
		}
	}
}

// ExceptBy produces the set difference of two sequences according to a specified key selector function.
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

// e := linq.ExceptBy(
// 	linq.FromSlice(p1),
// 	linq.FromSlice(p2),
// 	func(p Planet) (string, error) { return p.Name, nil })

// r, err := linq.ToSlice(e)
// exp := []Planet{{"Venus", Rock, 2}}
func ExceptBy[T any, K comparable, E IEnumerable[T]](first, second E, keySelector func(v T) (K, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &exceptEnumerator[T, K]{
			fst:  first(),
			snd:  second(),
			eq:   alwaysEqual[T],
			hash: keySelector,
		}
	}
}

func (e *exceptEnumerator[T, H]) Next() (def T, _ error) {
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
		if !has {
			return v, nil
		}
	}
}
