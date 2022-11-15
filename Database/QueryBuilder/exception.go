package querybuilder

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
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
