package exceptions

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
)

type ServiceProvider struct {
	DontReportExceptions []IExeption.Exception
}

func (provider ServiceProvider) Start() error {
	return nil
}

func (provider ServiceProvider) Stop() {
}

func (provider ServiceProvider) Register(container IFoundation.IApplication) {

	container.NamedSingleton("exception.handler", func() IExeption.ExceptionHandler {
		return NewDefaultHandler(provider.DontReportExceptions)
	})
}
