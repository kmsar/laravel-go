package Cache

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Cache/drivers"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/ICache"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConfig"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IRedis"
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
	container.NamedSingleton("cache", func(
		config IConfig.Config,
		redis IRedis.RedisFactory,
		handler IExeption.ExceptionHandler) ICache.CacheFactory {
		factory := &Factory{
			config:           config.Get("cache").(Config),
			exceptionHandler: handler,
			stores:           make(map[string]ICache.CacheStore),
			drivers: map[string]ICache.CacheStoreProvider{
				"memory": drivers.NewMemory,
			},
		}

		factory.Extend("redis", func(cacheConfig Support.Fields) ICache.CacheStore {
			return drivers.NewRedisCache(
				redis.Connection(Field.GetStringField(cacheConfig, "connection")),
				Field.GetStringField(cacheConfig, "prefix"),
			)
		})

		return factory
	})
	container.NamedSingleton("cache.store", func(factory ICache.CacheFactory) ICache.CacheStore {
		return factory.Store()
	})
}
