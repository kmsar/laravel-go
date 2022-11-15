package Filesystem

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConfig"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFilesystem"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
)

type ServiceProvider struct {
}

func (this ServiceProvider) Stop() {

}

func (this ServiceProvider) Start() error {
	return nil
}

func (this ServiceProvider) Register(container IFoundation.IApplication) {
	container.NamedSingleton("filesystem", func(config IConfig.Config) IFilesystem.FileSystemFactory {
		return New(config.Get("filesystem").(Config))
	})

	container.NamedSingleton("system.default", func(factory IFilesystem.FileSystemFactory) IFilesystem.FileSystem {
		return factory
	})

	//container.Singleton("system.qiniu", func(factory contracts.FileSystemFactory) *adapters.Qiniu {
	//	var adapter, _ = factory.Disk("qiniu").(*adapters.Qiniu)
	//	return adapter
	//})
}
