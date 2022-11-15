package Hashing

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConfig"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHashing"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Field"
)

type ServiceProvider struct {
}

func (this ServiceProvider) Stop() {

}

func (this ServiceProvider) Start() error {
	return nil
}

func (this ServiceProvider) Register(container IFoundation.IApplication) {
	container.NamedSingleton("hash", func(config IConfig.Config) IHashing.HasherFactory {
		return &Factory{
			config: config,
			hashes: make(map[string]IHashing.Hasher),
			drivers: map[string]IHashing.HasherProvider{
				"bcrypt": func(config Support.Fields) IHashing.Hasher {
					return &Bcrypt{
						cost: Field.GetIntField(config, "cost", 14),
						salt: Field.GetStringField(config, "salt"),
					}
				},
				"md5": func(config Support.Fields) IHashing.Hasher {
					return &Md5{
						salt: Field.GetStringField(config, "salt"),
					}
				},
			},
		}
	})

	container.NamedSingleton("hashing", func(factory IHashing.HasherFactory) IHashing.Hasher {
		return factory
	})
}
