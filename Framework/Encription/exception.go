package Encription

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
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
