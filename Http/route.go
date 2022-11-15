package Http

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IContainer"
)

type route struct {
	method      []string
	path        string
	middlewares []IContainer.MagicalFunc
	handler     IContainer.MagicalFunc
}

func (route *route) Middlewares() []IContainer.MagicalFunc {
	return route.middlewares
}

func (route *route) Method() []string {
	return route.method
}

func (route *route) Path() string {
	return route.path
}

func (route *route) Handler() IContainer.MagicalFunc {
	return route.handler
}
