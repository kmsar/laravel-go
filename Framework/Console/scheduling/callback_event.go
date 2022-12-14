package scheduling

import (
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IConsole"
	"github.com/kmsar/laravel-go/Framework/Hashing"
)

func NewCallbackEvent(mutex *Mutex, callback interface{}, timezone string) IConsole.CallbackEvent {
	return &CallbackEvent{
		Event:       NewEvent(mutex, callback, timezone),
		description: "",
	}
}

type CallbackEvent struct {
	*Event
	description string
}

func (this *CallbackEvent) Description(description string) IConsole.CallbackEvent {
	this.description = description
	return this
}

func (this *CallbackEvent) MutexName() string {
	if this.mutexName == "" {
		return fmt.Sprintf("goal.schedule-%s", Hashing.Md5Hash(this.expression+this.description))
	}
	return this.mutexName
}
