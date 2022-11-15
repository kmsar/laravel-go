package Validation

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IPipeline"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IValidation"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

func VerifyValidatable(request Support.Context, pipe IPipeline.Pipe) interface{} {

	if form, ok := request.(IValidation.Validatable); ok {
		VerifyForm(form)
	}

	return pipe(request)
}
