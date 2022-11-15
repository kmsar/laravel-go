package class

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

type TypeException struct {
	error
	fields Support.Fields
}

func (e TypeException) Error() string {
	return e.error.Error()
}

func (e TypeException) Fields() Support.Fields {
	return e.fields
}
