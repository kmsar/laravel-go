package Auth

import (
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IAuth"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Exceptions"
	"github.com/kmsar/laravel-go/Framework/Support/Field"
)

type Auth struct {
	authConfig Config

	guardDrivers  map[string]IAuth.GuardDriver
	userProviders map[string]IAuth.UserProvider
	userDrivers   map[string]IAuth.UserProviderDriver
}

func (this *Auth) ExtendUserProvider(key string, provider IAuth.UserProviderDriver) {
	this.userDrivers[key] = provider
}

func (this *Auth) ExtendGuard(key string, guard IAuth.GuardDriver) {
	this.guardDrivers[key] = guard
}

func (this *Auth) Guard(key string, ctx Support.Context) IAuth.Guard {
	config := this.authConfig.Guards[key]
	driver := Field.GetStringField(config, "driver")

	if guardDriver, existsDriver := this.guardDrivers[driver]; existsDriver {
		return guardDriver(key, config, ctx, this.UserProvider(Field.GetStringField(config, "provider")))
	}

	panic(GuardException{
		Exception: Exceptions.New(fmt.Sprintf("unsupported guard driver：%s", driver), config),
	})
}

func (this *Auth) UserProvider(key string) IAuth.UserProvider {
	if userProvider, existsUserProvider := this.userProviders[key]; existsUserProvider {
		return userProvider
	}

	config := this.authConfig.Users[key]
	driver := Field.GetStringField(config, "driver")

	if userDriver, existsProvider := this.userDrivers[driver]; existsProvider {
		this.userProviders[key] = userDriver(config)
		return this.userProviders[key]
	}

	panic(UserProviderException{
		Exception: Exceptions.New(fmt.Sprintf("unsupported user driver：%s", driver), config),
	})
}
