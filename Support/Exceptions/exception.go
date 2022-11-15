package Exceptions

import (
	"errors"
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

func WithError(err error, fields Support.Fields) IExeption.Exception {
	if e, isException := err.(IExeption.Exception); isException {
		return e
	}
	return New(err.Error(), fields)
}

func WithRecover(err interface{}, fields Support.Fields) IExeption.Exception {
	switch e := err.(type) {
	case IExeption.Exception:
		return e
	case error:
		return WithError(e, fields)
	case string:
		return New(e, fields)
	case fmt.Stringer:
		return New(e.String(), fields)
	}
	return New(fmt.Sprintf("%v", err), fields)
}

func WithPrevious(err error, fields Support.Fields, previous IExeption.Exception) Exception {
	return Exception{err, fields, previous}
}

func New(err string, fields Support.Fields) Exception {
	return Exception{errors.New(err), fields, nil}
}

func Throw(err interface{}) {
	if err != nil {
		panic(ResolveException(err))
	}
}

type Exception struct {
	error
	fields   Support.Fields
	previous IExeption.Exception
}

func (e Exception) Fields() Support.Fields {
	return e.fields
}
