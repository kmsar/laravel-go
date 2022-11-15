package Auth

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IAuth"
	"github.com/kmsar/laravel-go/Framework/Contracts/IConfig"
	"github.com/kmsar/laravel-go/Framework/Contracts/IHttp"
	"github.com/kmsar/laravel-go/Framework/Contracts/IPipeline"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Exceptions"
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
