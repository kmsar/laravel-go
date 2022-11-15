package IValidation

import "github.com/kmsar/laravel-go/Framework/Contracts/Support"

// Validatable 可验证的表单
type Validatable interface {
	Support.FieldsProvider

	Rules() Support.Fields
}

type ShouldValidate interface {
	Validatable

	ShouldVerify() bool
}
