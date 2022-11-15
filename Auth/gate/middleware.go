package gate

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IAuth"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IPipeline"
)

func Authorize(ability string, arguments ...interface{}) interface{} {
	return func(request IHttp.HttpContext, next IPipeline.Pipe, gate IAuth.Gate) interface{} {
		gate.Authorize(ability, arguments...)
		return next(request)
	}
}
