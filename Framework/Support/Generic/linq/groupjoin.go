package linq

type groupJoinEnumerator[S1, S2, T any, K comparable] struct {
	eOut  Enumerator[S1]
	eIn   Enumerator[S2]
	ksOut func(S1) (K, error)
	ksIn  func(S2) (K, error)
	rSel  func(S1, Enumerable[S2]) (T, error)

	ms2 *hashMap[K, S2]
}

// GroupJoin correlates the elements of two sequences based on equality of keys and groups the results.
// type Person struct {
// 	Name string
// }

// type Pet struct {
// 	Name  string
// 	Owner *Person
// }

// alice := &Person{"Alice"}
// bob := &Person{"Bob"}
// charlie := &Person{"Charlie"}

// abby := &Pet{"Abby", alice}
// bailey := &Pet{"Bailey", bob}
// bella := &Pet{"Bella", bob}
// cody := &Pet{"Cody", charlie}

// people := linq.FromSlice([]*Person{alice, bob, charlie})
// pets := linq.FromSlice([]*Pet{cody, bella, abby, bailey})

// r, err := linq.ToSlice(
// 	linq.GroupJoin(
// 		people, pets,
// 		func(p *Person) (*Person, error) { return p, nil },
// 		func(p *Pet) (*Person, error) { return p.Owner, nil },
// 		func(person *Person, pets linq.Enumerable[*Pet]) ([]string, error) {
// 			names, err := linq.ToSlice(
// 				linq.Select(pets, func(p *Pet) (string, error) { return p.Name, nil }))
// 			if err != nil {
// 				return nil, err
// 			}
// 			return append([]string{person.Name}, names...), nil
// 		}))
// if err != nil {
// 	t.Fatalf("%v", err)
// }

// exp := [][]string{
// 	{"Alice", "Abby"},
// 	{"Bob", "Bella", "Bailey"},
// 	{"Charlie", "Cody"},
// }
// if !reflect.DeepEqual(r, exp) {
// 	t.Fatalf("%v, wants %v", r, exp)
// }
func GroupJoin[S1, S2, T any, K comparable, E1 IEnumerable[S1], E2 IEnumerable[S2]](
	outer E1,
	inner E2,
	outerKeySelector func(S1) (K, error),
	innerKeySelector func(S2) (K, error),
	resultSelector func(S1, Enumerable[S2]) (T, error),
) Enumerable[T] {
	return func() Enumerator[T] {
		return &groupJoinEnumerator[S1, S2, T, K]{
			eOut:  outer(),
			eIn:   inner(),
			ksOut: outerKeySelector,
			ksIn:  innerKeySelector,
			rSel:  resultSelector,
		}
	}
}

func (e *groupJoinEnumerator[S1, S2, T, K]) Next() (def T, _ error) {
	out, err := e.eOut.Next()
	if err != nil {
		return def, err
	}
	if e.ms2 == nil {
		e.ms2 = newHashMap(e.ksIn, alwaysNotEqual[S2])
		if err := e.ms2.addAll(e.eIn); err != nil {
			return def, err
		}
	}
	k, err := e.ksOut(out)
	if err != nil {
		return def, err
	}

	return e.rSel(out, FromSlice(e.ms2.gets(k)))
}
