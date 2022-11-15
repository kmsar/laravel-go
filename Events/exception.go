package Events

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEvent"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

type EventException struct {
	error
	fields Support.Fields
	event  IEvent.Event
}

func (e EventException) Error() string {
	return e.error.Error()
}

func (e EventException) Fields() Support.Fields {
	return e.fields
}
