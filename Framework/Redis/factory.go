package Redis

import (
	"fmt"
	goredis "github.com/go-redis/redis/v9"
	"github.com/kmsar/laravel-go/Framework/Support/Field"
	"time"

	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
	"github.com/kmsar/laravel-go/Framework/Contracts/IRedis"
	"github.com/kmsar/laravel-go/Framework/Support/Utils"
	"sync"
)

type Factory struct {
	config           Config
	exceptionHandler IExeption.ExceptionHandler
	connections      map[string]IRedis.RedisConnection
	mutex            sync.Mutex
}

func (this *Factory) Connection(names ...string) IRedis.RedisConnection {
	name := Utils.DefaultString(names, this.config.Default)

	if connection, existsConnection := this.connections[name]; existsConnection {
		return connection
	}

	config := this.config.Stores[name]

	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.connections[name] = &Connection{
		exceptionHandler: this.exceptionHandler,
		client: goredis.NewClient(&goredis.Options{
			// The network type, either tcp or unix.
			// Default is tcp.
			Network: Field.GetStringField(config, "network", "tcp"),
			//	// host:port address.
			Addr: fmt.Sprintf("%s:%s",
				Field.GetStringField(config, "host", "127.0.0.1"),
				Field.GetStringField(config, "port", "6379"),
			),
			// Dialer creates new network connection and has priority over
			// Network and Addr options.
			Dialer: nil,

			//// Hook that is called when new connection is established.
			OnConnect: nil,
			Username:  Field.GetStringField(config, "username"),
			Password:  Field.GetStringField(config, "password"),
			// CredentialsProvider allows the username and password to be updated
			// before reconnecting. It should return the current username and password.
			//CredentialsProvider :nil,

			//// Database to be selected after connecting to the server.
			DB: Field.GetIntField(config, "db", 0),

			//// Default is 3 retries; -1 (not 0) disables retries.
			MaxRetries:      Field.GetIntField(config, "retries", 3),
			MinRetryBackoff: 8 * time.Millisecond,
			MaxRetryBackoff: 512 * time.Millisecond,
			DialTimeout:     5 * time.Second,
			ReadTimeout:     0,
			WriteTimeout:    0,

			// ContextTimeoutEnabled controls whether the client respects context timeouts and deadlines.
			// See https://redis.uptrace.dev/guide/go-redis-debugging.html#timeouts
			ContextTimeoutEnabled: false,

			// Type of connection pool.
			// true for FIFO pool, false for LIFO pool.
			// Note that FIFO has slightly higher overhead compared to LIFO,
			// but it helps closing idle connections faster reducing the pool size.
			PoolFIFO: false,
			// Maximum number of socket connections.
			// Default is 10 connections per every available CPU as reported by runtime.GOMAXPROCS.
			PoolSize: 10,
			// Default is ReadTimeout + 1 second.
			//PoolTimeout time.Duration

			// Minimum number of idle connections which is useful when establishing
			// new connection is slow.
			MinIdleConns: 0,

			// Maximum number of idle connections.
			//MaxIdleConns int

			// ConnMaxIdleTime is the maximum amount of time a connection may be idle.
			// Should be less than server's timeout.
			//
			// Expired connections may be closed lazily before reuse.
			// If d <= 0, connections are not closed due to a connection's idle time.
			//
			// Default is 5 minutes. -1 disables idle timeout check.
			ConnMaxIdleTime: 5 * time.Minute,

			// ConnMaxLifetime is the maximum amount of time a connection may be reused.
			//
			// Expired connections may be closed lazily before reuse.
			// If <= 0, connections are not closed due to a connection's age.
			//
			// Default is to not close idle connections.
			ConnMaxLifetime: -1,

			PoolTimeout: 0,

			//// Limiter interface used to implement circuit breaker or rate limiter.
			Limiter: nil,

			// TLS Config to use. When set, TLS will be negotiated.
			//TLSConfig *tls.Config

			// Enables read only queries on slave/follower nodes.
			//readOnly bool
		}),
	}

	return this.connections[name]
}
