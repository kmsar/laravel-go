package linq

// Contains cetermines whether a sequence contains a specified element.
//r, err = linq.Contains(linq.FromSlice([]int{1, 2, 3, 4, 5}), 6)
func Contains[T comparable, E IEnumerable[T]](src E, val T) (bool, error) {
	e := src()
	for {
		t, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return false, nil
			}
			return false, err
		}

		if t == val {
			return true, nil
		}
	}
}

// ContainsFunc determines whether a sequence contains a specified element by using a specified comparer function.
// type Product struct {
// 	Name string
// 	Code int
// }
// src := []Product{
// 	{Name: "apple", Code: 9},
// 	{Name: "orange", Code: 4},
// 	{Name: "lemon", Code: 12},
// }
// e := linq.FromSlice(src)
// v := Product{Name: "apple", Code: 1}
// r, err := linq.ContainsFunc(e, v, func(a, b Product) (bool, error) {
// 	return a.Name == b.Name, nil
// })
func ContainsFunc[T any, E IEnumerable[T]](src E, val T, equals func(T, T) (bool, error)) (bool, error) {
	e := src()
	for {
		t, err := e.Next()
		if err != nil {
			if isEOC(err) {
				return false, nil
			}
			return false, err
		}

		eq, err := equals(t, val)
		if err != nil {
			return false, err
		}
		if eq {
			return true, nil
		}
	}
}
