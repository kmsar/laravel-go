package ICache

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Container"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"time"
)

// CacheStoreProvider cache storage provider.
type CacheStoreProvider func(cacheConfig Support.Fields) CacheStore

type CacheFactory interface {

	// Store Get a cache store instance by name.
	Store(name ...string) CacheStore

	// Extend Extended cache instance
	Extend(drive string, cacheStoreProvider CacheStoreProvider)
}

type CacheStore interface {

	// Get Retrieve an item from the cache by key.
	Get(key string) interface{}

	// Many Retrieve multiple items from the cache by key.
	Many(keys []string) []interface{}

	// Put Store an item in the cache.
	Put(key string, value interface{}, seconds time.Duration) error

	// Add Store an item in the cache if the key does not exist.
	Add(key string, value interface{}, ttl ...time.Duration) bool

	// Pull Retrieve an item from the cache and delete it.
	Pull(key string, defaultValue ...interface{}) interface{}

	// PutMany Store multiple items in the cache.
	PutMany(values map[string]interface{}, seconds time.Duration) error

	// Increment  the value of an item in the cache.
	Increment(key string, value ...int64) (int64, error)

	// Decrement  the value of an item in the cache.
	Decrement(key string, value ...int64) (int64, error)

	// Forever Store an item in the cache indefinitely.
	Forever(key string, value interface{}) error

	// Forget Remove an item from the cache.
	Forget(key string) error

	// Flush Remove all items from the cache.
	Flush() error

	// GetPrefix get the cache key prefix.
	GetPrefix() string

	// Remember get an item from the cache, or execute the given Closure and store the result.
	Remember(key string, ttl time.Duration, provider Container.InstanceProvider) interface{}

	// RememberForever get an item from the cache, or execute the given Closure and store the result forever.
	RememberForever(key string, provider Container.InstanceProvider) interface{}
}
