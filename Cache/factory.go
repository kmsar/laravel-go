package Cache

import (
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Field"

	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/ICache"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
)

type Factory struct {
	config           Config
	exceptionHandler IExeption.ExceptionHandler
	stores           map[string]ICache.CacheStore
	drivers          map[string]ICache.CacheStoreProvider
}

func (this *Factory) getName(names ...string) string {
	if len(names) > 0 {
		return names[0]
	}
	return this.config.Default

}

func (this Factory) getConfig(name string) Support.Fields {
	return this.config.Stores[name]
}

func (this *Factory) Store(names ...string) ICache.CacheStore {
	name := this.getName(names...)
	if cacheStore, existsStore := this.stores[name]; existsStore {
		return cacheStore
	}

	this.stores[name] = this.make(name)

	return this.stores[name]
}

func (this *Factory) Extend(driver string, cacheStoreProvider ICache.CacheStoreProvider) {
	this.drivers[driver] = cacheStoreProvider
}

func (this *Factory) make(name string) ICache.CacheStore {
	config := this.getConfig(name)
	driver := Field.GetStringField(config, "driver")
	driveProvider, existsProvider := this.drivers[driver]
	if !existsProvider {
		panic(DriverException{
			fmt.Errorf("不支持的缓存驱动：%s", driver),
			config,
		})
	}
	return driveProvider(config)
}
