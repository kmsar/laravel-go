package class

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
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
