package Validation

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

type Exception struct {
	param  Support.Fields
	errors Support.Fields
	string
}

func NewException(param Support.Fields, errors Support.Fields) Exception {
	return Exception{param, errors, "param validation failed"}
}

func (this Exception) Error() string {
	return this.string
}

func (this Exception) Fields() Support.Fields {
	return this.param
}

func (this Exception) GetErrors() Support.Fields {
	return this.errors
}
