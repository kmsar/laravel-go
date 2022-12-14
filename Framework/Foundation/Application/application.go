package Application

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IConfig"
	"github.com/kmsar/laravel-go/Framework/Contracts/IContainer"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Support/Parallel"
	"github.com/kmsar/laravel-go/Framework/Support/Utils"

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
	queue := Parallel.NewParallel(len(this.services))

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
