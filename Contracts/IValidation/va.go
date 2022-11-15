package IValidation

import "github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"

// Validatable 可验证的表单
type Validatable interface {
	Support.FieldsProvider

	Rules() Support.Fields
}

type ShouldValidate interface {
	Validatable

	ShouldVerify() bool
}
