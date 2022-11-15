package Redis

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

type SubscribeException struct {
	error
	fields Support.Fields
}

func (s SubscribeException) Error() string {
	return s.error.Error()
}

func (s SubscribeException) Fields() Support.Fields {
	return s.fields
}
