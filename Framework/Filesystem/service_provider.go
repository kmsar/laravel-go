package Filesystem

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IConfig"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFilesystem"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
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
