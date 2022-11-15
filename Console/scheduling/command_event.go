package scheduling

import (
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConsole"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Hashing"
)

func NewCommandEvent(command string, mutex *Mutex, callback interface{}, timezone string) IConsole.CommandEvent {
	return &CommandEvent{
		Event:   NewEvent(mutex, callback, timezone),
		command: command,
	}
}

type CommandEvent struct {
	*Event
	command string
}

func (this *CommandEvent) Command(command string) IConsole.CommandEvent {
	this.command = command
	return this
}

func (this *CommandEvent) MutexName() string {
	if this.mutexName == "" {
		return fmt.Sprintf("goal.schedule-%s", Hashing.Md5Hash(this.expression+this.command))
	}
	return this.mutexName
}
