package Validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IValidation"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Field"
)

var Validator = validator.New()

func Struct(data interface{}) error {
	return Validator.Struct(data)
}

func Valid(data interface{}, rules Support.Fields) Support.Fields {
	switch param := data.(type) {
	case Support.Fields:
		return Validator.ValidateMap(param, rules)
	case Support.FieldsProvider:
		return Validator.ValidateMap(param.Fields(), rules)
	}

	fields, err := Field.ConvertToFields(data)
	if err != nil {
		panic(Exception{
			param:  fields,
			errors: nil,
			string: "unsupported parameter type",
		})
	}

	return Validator.ValidateMap(fields, rules)
}

func Form(validatable IValidation.Validatable) Support.Fields {
	return Validator.ValidateMap(validatable.Fields(), validatable.Rules())
}

func VerifyForm(validatable IValidation.Validatable) {
	if errs := Form(validatable); len(errs) > 0 {
		panic(NewException(validatable.Fields(), validatable.Rules()))
	}
}

func VerifyStruct(data interface{}) {
	if err := Struct(data); err != nil {
		panic(err)
	}
}

func Verify(data interface{}, rules Support.Fields) {
	if errs := Valid(data, rules); len(errs) > 0 {
		var fields, _ = Field.ConvertToFields(data)
		panic(NewException(fields, rules))
	}
}
