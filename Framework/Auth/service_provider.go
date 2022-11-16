package Auth

import (
	"github.com/kmsar/laravel-go/Framework/Auth/guards"
	"github.com/kmsar/laravel-go/Framework/Auth/providers"
	"github.com/kmsar/laravel-go/Framework/Contracts/IAuth"
	"github.com/kmsar/laravel-go/Framework/Contracts/IConfig"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Contracts/IRedis"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

type ServiceProvider struct {
}

func (this *ServiceProvider) Boot(application IFoundation.IApplication) {
	//TODO implement me
	panic("implement me")
}
func (this ServiceProvider) Start() error {
	return nil
}

func (this ServiceProvider) Stop() {
}

func (this ServiceProvider) Register(container IFoundation.IApplication) {
	container.NamedSingleton("auth", func(config IConfig.Config, factory IRedis.RedisFactory) IAuth.Auth {
		authConfig := config.Get("auth").(Config)

		return &Auth{
			authConfig: authConfig,
			guardDrivers: map[string]IAuth.GuardDriver{
				"jwt": func(name string, config Support.Fields, ctx Support.Context, provider IAuth.UserProvider) IAuth.Guard {
					guard := guards.JwtGuard(name, config, ctx, provider)

					if factory != nil { // 有 redis 的话
						if redisConnName, ok := config["redis"].(string); ok {
							guard.SetRedis(factory.Connection(redisConnName))
						} else {
							guard.SetRedis(factory.Connection())
						}
					}

					return guard
				},
				"session": guards.SessionGuard,
			},
			userDrivers: map[string]IAuth.UserProviderDriver{
				"db": providers.DBDriver,
			},
			userProviders: make(map[string]IAuth.UserProvider),
		}
	})
	container.NamedBind("auth.guard", func(config IConfig.Config, auth IAuth.Auth, ctx Support.Context) IAuth.Guard {
		return auth.Guard(config.Get("auth").(Config).Defaults.Guard, ctx)
	})
}
