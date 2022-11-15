package Redis

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
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
