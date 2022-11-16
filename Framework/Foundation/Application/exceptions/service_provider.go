package exceptions

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
)

type ServiceProvider struct {
	DontReportExceptions []IExeption.Exception
}

func (provider ServiceProvider) Start() error {
	return nil
}
func (this *ServiceProvider) Boot(application IFoundation.IApplication) {
	//TODO implement me
	panic("implement me")
}
func (provider ServiceProvider) Stop() {
}

func (provider ServiceProvider) Register(container IFoundation.IApplication) {

	container.NamedSingleton("exception.handler", func() IExeption.ExceptionHandler {
		return NewDefaultHandler(provider.DontReportExceptions)
	})
}
