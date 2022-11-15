package OrderedMaps

import (
	"strings"
)

// Set m to an ordered map from string to string,
// using strings.Compare as the comparison function.
var m = New[string, string](strings.Compare)

// Add adds the pair a, b to m.
func Add(a, b string) {
	m.Insert(a, b)
	d := NewStringStringArray()
	d.Get()
	d.Set("a", "b")

}
