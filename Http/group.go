package Http

import (
	"errors"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Container"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Http/echo"

	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IContainer"
	IRouting "github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IRouting"
)

var (
	MethodTypeError = errors.New("http method type unknown")
)

type group struct {
	prefix      string
	middlewares []IContainer.MagicalFunc
	routes      []IRouting.Route
	groups      []IRouting.RouteGroup
}

func NewGroup(prefix string, middlewares ...interface{}) IRouting.RouteGroup {
	return &group{
		prefix:      prefix,
		routes:      make([]IRouting.Route, 0),
		groups:      make([]IRouting.RouteGroup, 0),
		middlewares: convertToMiddlewares(middlewares...),
	}
}

// AddRoute 添加一条路由
func (group *group) AddRoute(route IRouting.Route) IRouting.RouteGroup {
	group.routes = append(group.routes, route)

	return group
}

// Group 添加一个子组
func (group *group) Group(prefix string, middlewares ...interface{}) IRouting.RouteGroup {
	var groupInstance = NewGroup(group.prefix+prefix, middlewares...)

	group.groups = append(group.groups, groupInstance)

	return groupInstance
}

// Add 添加路由，method 只允许字符串或者字符串数组
func (group *group) Add(method interface{}, path string, handler interface{}, middlewares ...interface{}) IRouting.RouteGroup {
	methods := make([]string, 0)
	switch r := method.(type) {
	case string:
		methods = []string{r}
	case []string:
		methods = r
	default:
		panic(MethodTypeError)
	}
	group.AddRoute(&route{
		method:      methods,
		path:        group.prefix + path,
		middlewares: convertToMiddlewares(middlewares...),
		handler:     Container.NewMagicalFunc(handler),
	})

	return group
}

func (group *group) Get(path string, handler interface{}, middlewares ...interface{}) IRouting.RouteGroup {
	return group.Add(echo.GET, path, handler, middlewares...)
}

func (group *group) Post(path string, handler interface{}, middlewares ...interface{}) IRouting.RouteGroup {
	return group.Add(echo.POST, path, handler, middlewares...)
}

func (group *group) Delete(path string, handler interface{}, middlewares ...interface{}) IRouting.RouteGroup {
	return group.Add(echo.DELETE, path, handler, middlewares...)
}

func (group *group) Put(path string, handler interface{}, middlewares ...interface{}) IRouting.RouteGroup {
	return group.Add(echo.PUT, path, handler, middlewares...)
}

func (group *group) Trace(path string, handler interface{}, middlewares ...interface{}) IRouting.RouteGroup {
	return group.Add(echo.TRACE, path, handler, middlewares...)
}

func (group *group) Patch(path string, handler interface{}, middlewares ...interface{}) IRouting.RouteGroup {
	return group.Add(echo.PATCH, path, handler, middlewares...)
}

func (group *group) Options(path string, handler interface{}, middlewares ...interface{}) IRouting.RouteGroup {
	return group.Add(echo.OPTIONS, path, handler, middlewares...)
}

func (group *group) Middlewares() []IContainer.MagicalFunc {
	return group.middlewares
}

func (group *group) Groups() []IRouting.RouteGroup {
	return group.groups
}

func (group *group) Routes() []IRouting.Route {
	return group.routes
}
