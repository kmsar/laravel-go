package Events

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IEvent"
	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
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
func (this *ServiceProvider) Boot(application IFoundation.IApplication) {
	//TODO implement me
	panic("implement me")
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
