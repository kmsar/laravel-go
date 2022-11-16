package Encription

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IConfig"
	"github.com/kmsar/laravel-go/Framework/Contracts/IEncryption"
	"github.com/kmsar/laravel-go/Framework/Contracts/IEnv"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
)

type ServiceProvider struct {
}

func (this *ServiceProvider) Boot(application IFoundation.IApplication) {
	//TODO implement me
	panic("implement me")
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
