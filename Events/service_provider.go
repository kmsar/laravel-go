package Events

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEvent"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
)

var dispatcher IEvent.EventDispatcher

type ServiceProvider struct {
}

func Dispatch(event IEvent.Event) bool {
	if dispatcher != nil {
		dispatcher.Dispatch(event)
		return true
	}
	return false
}

func (this ServiceProvider) Stop() {

}

func (this ServiceProvider) Start() error {
	return nil
}

func (provider ServiceProvider) Register(container IFoundation.IApplication) {
	container.NamedSingleton("events", func(handler IExeption.ExceptionHandler) IEvent.EventDispatcher {
		dispatcher = NewDispatcher(handler)
		return dispatcher
	})
}
