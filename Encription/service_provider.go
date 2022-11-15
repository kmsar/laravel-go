package Encription

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConfig"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEncryption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEnv"
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
	container.NamedSingleton("encryption", func(config IConfig.Config, env IEnv.Env) IEncryption.EncryptorFactory {
		factory := &Factory{encryptors: make(map[string]IEncryption.Encryptor)}

		factory.Extend("default", AES(env.GetString("app.key")))

		return factory
	})

	container.NamedSingleton("encryption.default", func(factory IEncryption.EncryptorFactory) IEncryption.Encryptor {
		return factory.Driver("default")
	})
}
