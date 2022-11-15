package linq

type distinctEnumerator[T any, K comparable] struct {
	src  Enumerator[T]
	hmap *hashMap[K, T]
}

// Distinct returns distinct elements from a sequence by using the specified comparer functions.
// ages := []int{21, 46, 46, 55, 17, 21, 55, 55, 25, 25}

// e := linq.Distinct(
// 	linq.FromSlice(ages),
// 	func(a, b int) (bool, error) { return a == b, nil },
// 	func(a int) (int, error) { return a / 10, nil })

// r, err := linq.ToSlice(e)
// if err != nil {
// 	t.Fatalf("%v", err)
// }

// exp := []int{21, 46, 55, 17, 25}
// if !reflect.DeepEqual(r, exp) {
// 	t.Fatalf("%v, wants %v", r, exp)
// }
func Distinct[T any, E IEnumerable[T]](src E, equals func(T, T) (bool, error), getHashCode func(T) (int, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &distinctEnumerator[T, int]{
			src:  src(),
			hmap: newHashMap(getHashCode, equals),
		}
	}
}

// DistinctBy returns distinct elements from a sequence according to a specified key selector function.
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
// planets := []Planet{
// 	{"Marcury", Rock, 1},
// 	{"Venus", Rock, 2},
// 	{"Earth", Rock, 3},
// 	{"Mars", Rock, 4},
// 	{"Jupiter", Gas, 5},
// 	{"Saturn", Gas, 6},
// 	{"Uranus", Liquid, 7},
// 	{"Neptune", Liquid, 8},
// 	{"Pluto", Ice, 9}, // dwarf planet
// }

// e := linq.DistinctBy(linq.FromSlice(planets), func(p Planet) (PlanetType, error) {
// 	return p.Type, nil
// })
// r, err := linq.ToSlice(linq.Select(e, func(p Planet) (string, error) { return p.Name, nil }))

// exp := []string{"Marcury", "Jupiter", "Uranus", "Pluto"}
// if !reflect.DeepEqual(r, exp)
func DistinctBy[T any, K comparable, E IEnumerable[T]](src E, keySelector func(v T) (K, error)) Enumerable[T] {
	return func() Enumerator[T] {
		return &distinctEnumerator[T, K]{
			src:  src(),
			hmap: newKeyMap(keySelector),
		}
	}
}

func (e *distinctEnumerator[T, int]) Next() (def T, _ error) {
	for {
		v, err := e.src.Next()
		if err != nil {
			return def, err
		}

		added, err := e.hmap.add(v)
		if err != nil {
			return def, err
		}
		if added {
			return v, nil
		}
	}
}
