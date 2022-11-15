package gate

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IAuth"
	"github.com/kmsar/laravel-go/Framework/Contracts/IHttp"
	"github.com/kmsar/laravel-go/Framework/Contracts/IPipeline"
)

func Authorize(ability string, arguments ...interface{}) interface{} {
	return func(request IHttp.HttpContext, next IPipeline.Pipe, gate IAuth.Gate) interface{} {
		gate.Authorize(ability, arguments...)
		return next(request)
	}
}
