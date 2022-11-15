package Application

import (
	"github.com/kmsar/laravel-go/Framework/Container"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
)

var (
	instance IFoundation.IApplication
)

func Singleton() IFoundation.IApplication {
	if instance != nil {
		return instance
	}

	instance = &application{
		IContainer: Container.New(),
		services:   make([]IFoundation.ServiceProvider, 0),
	}

	return instance
}

func SetSingleton(app IFoundation.IApplication) {
	instance = app
}

func Get(key string, args ...interface{}) interface{} {
	return Singleton().Get(key, args...)
}
