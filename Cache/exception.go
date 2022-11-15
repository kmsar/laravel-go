package Cache

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

type DriverException struct {
	error
	fields Support.Fields
}

func (this DriverException) Error() string {
	return this.error.Error()
}

func (this DriverException) Fields() Support.Fields {
	return this.fields
}
