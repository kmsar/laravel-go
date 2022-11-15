package querybuilder

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

type ParamException struct {
	error
	fields Support.Fields
}

func (p ParamException) Error() string {
	return p.error.Error()
}

func (p ParamException) Fields() Support.Fields {
	return p.fields
}
