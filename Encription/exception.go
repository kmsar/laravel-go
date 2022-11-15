package Encription

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

type EncryptException struct {
	error
	fields Support.Fields
}

func (e EncryptException) Error() string {
	return e.error.Error()
}

func (e EncryptException) Fields() Support.Fields {
	return e.fields
}
