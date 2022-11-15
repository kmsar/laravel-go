package mapset

import (
	"testing"
)

type yourType struct {
	name string
}

func Test_ExampleIterator(t *testing.T) {

	s := NewSet[*yourType](
		[]*yourType{
			&yourType{name: "Alise"},
			&yourType{name: "Bob"},
			&yourType{name: "John"},
			&yourType{name: "Nick"},
		}...,
	)

	var found *yourType
	it := s.Iterator()

	for elem := range it.C {
		if elem.name == "John" {
			found = elem
			it.Stop()
		}
	}

	if found == nil || found.name != "John" {
		t.Fatalf("expected iterator to have found `John` record but got nil or something else")
	}
}
