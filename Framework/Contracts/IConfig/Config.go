package IConfig

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IEnv"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

type ConfigProvider func(env IEnv.Env) interface{}

type Config interface {
	Support.Getter
	Support.FieldsProvider

	// Load  configuration based on given field provider.
	Load(provider Support.FieldsProvider)

	// Reload  configuration based on given field provider.
	Reload()

	// Merge  the given configuration values.
	Merge(key string, config Config)

	// Get  the specified configuration value.
	Get(key string, defaultValue ...interface{}) interface{}

	// Set set a given configuration value.
	Set(key string, value interface{})

	// Unset Destroy the specified configuration value.
	Unset(key string)

	// GetConfig get the specified configuration instance.
	GetConfig(key string) Config
}
