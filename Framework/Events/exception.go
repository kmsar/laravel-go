package Events

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IEvent"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
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
