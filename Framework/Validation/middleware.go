package Validation

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IPipeline"
	"github.com/kmsar/laravel-go/Framework/Contracts/IValidation"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

func VerifyValidatable(request Support.Context, pipe IPipeline.Pipe) interface{} {

	if form, ok := request.(IValidation.Validatable); ok {
		VerifyForm(form)
	}

	return pipe(request)
}
