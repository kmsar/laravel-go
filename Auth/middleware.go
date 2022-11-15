package Auth

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IAuth"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConfig"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IPipeline"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Exceptions"
)

func Guard(guards ...string) interface{} {
	return func(request IHttp.IHttpRequest, next IPipeline.Pipe, auth IAuth.Auth, config IConfig.Config) interface{} {

		if len(guards) == 0 {
			guards = append(guards, config.Get("auth").(Config).Defaults.Guard)
		}

		for _, guard := range guards {
			if auth.Guard(guard, request).Guest() {
				panic(Exception{
					Exception: Exceptions.New(guard+" guard authentication failed", Support.Fields{
						"guards": guards,
					}),
				})
			}
		}

		return next(request)
	}
}
