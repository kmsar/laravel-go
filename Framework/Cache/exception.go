package Cache

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
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
