package Serialization

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IConfig"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Contracts/ISerialize"
)

type ServiceProvider struct {
	app IFoundation.IApplication
}

func (this *ServiceProvider) Register(application IFoundation.IApplication) {
	this.app = application
	application.NamedSingleton("serialization", func() ISerialize.Serialization {
		return New()
	})
	application.NamedSingleton("serializer", func(config IConfig.Config, serialization ISerialize.Serialization) ISerialize.Serializer {
		return serialization.Method(config.Get("serialization").(Config).Default)
	})
	application.NamedSingleton("class.serializer", func(config IConfig.Config) ISerialize.ClassSerializer {
		var serializer = NewClassSerializer(application)
		for _, class := range config.Get("serialization").(Config).Class {
			serializer.Register(class)
		}
		return serializer
	})
}

func (this *ServiceProvider) Start() error {
	return nil
}
func (this *ServiceProvider) Boot(application IFoundation.IApplication) {
	//TODO implement me
	panic("implement me")
}
func (this *ServiceProvider) Stop() {
}
