package Cache

import (
	"github.com/kmsar/laravel-go/Framework/Cache/drivers"
	"github.com/kmsar/laravel-go/Framework/Contracts/ICache"
	"github.com/kmsar/laravel-go/Framework/Contracts/IConfig"
	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Contracts/IRedis"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Field"
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
