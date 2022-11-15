package Http

import (
	"github.com/kmsar/laravel-go/Framework/Container"
	"github.com/kmsar/laravel-go/Framework/Contracts/IContainer"
)

func convertToMiddlewares(middlewares ...interface{}) (results []IContainer.MagicalFunc) {
	for _, middleware := range middlewares {
		magicalFunc, isMiddleware := middleware.(IContainer.MagicalFunc)
		if !isMiddleware {
			magicalFunc = Container.NewMagicalFunc(middleware)
		}
		if magicalFunc.NumOut() != 1 {
			panic(MiddlewareError)
		}
		results = append(results, magicalFunc)
	}
	return
}
