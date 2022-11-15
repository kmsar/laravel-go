package Redis

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConfig"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IRedis"
	"sync"
)

var (
	factory IRedis.RedisFactory
	cli     IRedis.RedisConnection
)

func Default() IRedis.RedisConnection {
	if cli == nil {
		cli = factory.Connection()
	}
	return cli
}

func DefaultFactory() IRedis.RedisFactory {
	return factory
}

type ServiceProvider struct {
}

func (this ServiceProvider) Stop() {

}

func (this ServiceProvider) Start() error {
	return nil
}

func (this ServiceProvider) Register(app IFoundation.IApplication) {

	app.NamedSingleton("redis.factory", func(config IConfig.Config, handler IExeption.ExceptionHandler) IRedis.RedisFactory {
		factory = &Factory{
			config:           config.Get("redis").(Config),
			exceptionHandler: handler,
			connections:      make(map[string]IRedis.RedisConnection),
			mutex:            sync.Mutex{},
		}

		return factory
	})

	app.NamedSingleton("redis", func(factory IRedis.RedisFactory) IRedis.RedisConnection {
		return factory.Connection()
	})

	app.NamedSingleton("redis.connection", func(redis IRedis.RedisConnection) *Connection {
		return redis.(*Connection)
	})
}
