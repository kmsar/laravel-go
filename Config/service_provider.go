package Config

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConfig"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEnv"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
)

func Service(env, path string, config map[string]IConfig.ConfigProvider) IFoundation.ServiceProvider {
	return &ServiceProvider{
		app:             nil,
		Env:             env,
		Paths:           []string{path},
		Sep:             "=",
		ConfigProviders: config,
	}
}

type ServiceProvider struct {
	app             IFoundation.IApplication
	Env             string
	Paths           []string
	Sep             string
	ConfigProviders map[string]IConfig.ConfigProvider
}

func (this *ServiceProvider) Boot(application IFoundation.IApplication) {
	//TODO implement me
	panic("implement me")
}

func (this *ServiceProvider) Stop() {

}

func (this *ServiceProvider) Start() error {
	return nil
}

func (this *ServiceProvider) Register(application IFoundation.IApplication) {
	this.app = application

	application.NamedSingleton("env", func() IEnv.Env {
		return NewEnv(this.Paths, this.Sep)
	})

	application.NamedSingleton("config", func(env IEnv.Env) IConfig.Config {
		configInstance := NewConfig(env, this.ConfigProviders)

		for key, provider := range this.ConfigProviders {
			configInstance.Set(key, provider(env))
		}
		return configInstance
	})
}
