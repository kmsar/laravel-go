package application

import (
	"github.com/qbhy/parallel"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConfig"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IContainer"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Utils"
	"reflect"
)

const EnvProduction = "production"

type application struct {
	IContainer.IContainer
	services []IFoundation.ServiceProvider
}

func (this *application) IsEnvironment(environments ...string) bool {
	//TODO implement me
	panic("implement me")
}

func (this *application) Environment() string {
	return this.Get("config").(IConfig.Config).Get("app").(Config).Env
}

func (this *application) IsProduction() bool {
	return this.Environment() == EnvProduction
}

func (this *application) Debug() bool {
	return this.Get("config").(IConfig.Config).Get("app").(Config).Debug
}

func (this *application) Start() map[string]error {
	errors := make(map[string]error)
	queue := parallel.NewParallel(len(this.services))

	for _, service := range this.services {
		(func(service IFoundation.ServiceProvider) {
			_ = queue.Add(func() interface{} {
				return service.Start()
			})
		})(service)
	}

	results := queue.Wait()
	for serviceIndex, result := range results {
		if err, isErr := result.(error); isErr {
			errors[Utils.GetTypeKey(reflect.TypeOf(this.services[serviceIndex]))] = err
		}
	}

	return errors
}

func (this *application) Stop() {

	for serviceIndex := len(this.services) - 1; serviceIndex > -1; serviceIndex-- {
		this.services[serviceIndex].Stop()
	}
}

func (this *application) RegisterServices(services ...IFoundation.ServiceProvider) {
	this.services = append(this.services, services...)

	for _, service := range services {
		service.Register(this)
	}
}
