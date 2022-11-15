package Http

import (
	"errors"
	"github.com/kmsar/laravel-go/Framework/Pipeline"

	IRouting "github.com/kmsar/laravel-go/Framework/Contracts/IRouting"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Http/echo"
	"github.com/kmsar/laravel-go/Framework/Support/Exceptions"

	"github.com/kmsar/laravel-go/Framework/Container"
	"github.com/kmsar/laravel-go/Framework/Contracts/IContainer"
	"github.com/kmsar/laravel-go/Framework/Contracts/IEvent"
	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"strings"
)

var (
	MiddlewareError = errors.New("middleware error")

	// magical middlewares
	exceptionHandler = Container.NewMagicalFunc(func(handler IExeption.ExceptionHandler, exception Exception) interface{} {
		return handler.Handle(exception)
	})
)

func New(container IFoundation.IApplication) IRouting.Router {
	router := &Router{
		app:         container,
		events:      container.Get("events").(IEvent.EventDispatcher),
		echo:        echo.New(),
		routes:      make([]IRouting.Route, 0),
		groups:      make([]IRouting.RouteGroup, 0),
		middlewares: make([]IContainer.MagicalFunc, 0),
	}

	router.Use(router.recovery)

	return router
}

type Router struct {
	events IEvent.EventDispatcher
	app    IFoundation.IApplication
	echo   *echo.Echo
	groups []IRouting.RouteGroup
	routes []IRouting.Route

	// 全局中间件
	middlewares []IContainer.MagicalFunc
}

func (this *Router) Group(prefix string, middlewares ...interface{}) IRouting.RouteGroup {
	groupInstance := NewGroup(prefix, middlewares...)

	this.groups = append(this.groups, groupInstance)

	return groupInstance
}

func (this *Router) Close() error {
	return this.echo.Close()
}

func (this *Router) Static(path, directory string) {
	if strings.HasPrefix(directory, "/") {
		directory = this.app.Get("path").(string) + "/" + directory
	}
	this.echo.Static(path, directory)
}

func (this *Router) Get(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.GET, path, handler, middlewares...)
}

func (this *Router) Post(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.POST, path, handler, middlewares...)
}

func (this *Router) Delete(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.DELETE, path, handler, middlewares...)
}

func (this *Router) Put(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.PUT, path, handler, middlewares...)
}

func (this *Router) Patch(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.PATCH, path, handler, middlewares...)
}

func (this *Router) Options(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.OPTIONS, path, handler, middlewares...)
}

func (this *Router) Trace(path string, handler interface{}, middlewares ...interface{}) {
	this.Add(echo.TRACE, path, handler, middlewares...)
}

func (this *Router) Use(middlewares ...interface{}) {
	for _, middleware := range middlewares {
		if magicalFunc, ok := middleware.(IContainer.MagicalFunc); ok {
			this.middlewares = append(this.middlewares, magicalFunc)
		} else if echoMiddleware, isEchoFunc := middleware.(echo.MiddlewareFunc); isEchoFunc {
			this.echo.Use(echoMiddleware)
		} else {
			this.middlewares = append(this.middlewares, Container.NewMagicalFunc(middleware))
		}
	}
}

func (this *Router) Add(method interface{}, path string, handler interface{}, middlewares ...interface{}) {
	methods := make([]string, 0)
	switch v := method.(type) {
	case string:
		methods = []string{v}
	case []string:
		methods = v
	default:
		panic(errors.New("method 只能接收 string 或者 []string"))
	}
	this.routes = append(this.routes, &route{
		method:      methods,
		path:        path,
		middlewares: convertToMiddlewares(middlewares...),
		handler:     Container.NewMagicalFunc(handler),
	})
}

func (this *Router) mountGroup(group IRouting.RouteGroup) {
	this.mountRoutes(group.Routes(), group.Middlewares()...)

	for _, routeGroup := range group.Groups() {
		this.mountGroup(routeGroup)
	}
}

// Start 启动 httpserver
func (this *Router) Start(address string) error {

	this.mountRoutes(this.routes)

	for _, routeGroup := range this.groups {
		this.mountGroup(routeGroup)
	}

	this.echo.HTTPErrorHandler = func(err error, context echo.Context) {
		if result := this.app.StaticCall(exceptionHandler, Exception{Exception: Exceptions.WithError(err, Support.Fields{
			"status": context.Response().Status,
		}), Request: NewRequest(context)})[0]; result != nil {
			HandleResponse(result, NewRequest(context))
		}
	}
	this.echo.Debug = this.app.Debug()

	return this.echo.Start(address)
}

// mountRoutes 装配路由
func (this *Router) mountRoutes(routes []IRouting.Route, middlewares ...IContainer.MagicalFunc) {
	for _, routeItem := range routes {
		(func(routeInstance IRouting.Route) {
			this.echo.Match(routeInstance.Method(), routeInstance.Path(), func(context echo.Context) error {
				request := NewRequest(context)
				defer func() {
					this.events.Dispatch(&RequestAfter{request})
				}()

				// 触发钩子
				this.events.Dispatch(&RequestBefore{request})

				pipes := append(this.middlewares, middlewares...)
				pipes = append(pipes, routeInstance.Middlewares()...)

				var result interface{}
				if len(pipes) == 0 {
					results := this.app.StaticCall(routeInstance.Handler(), request)
					if len(results) > 0 {
						result = results[0]
					}
				} else {
					result = Pipeline.Static(this.app).SendStatic(request).
						ThroughStatic(
							this.middlewares...,
						).
						ThroughStatic(
							append(middlewares, routeInstance.Middlewares()...)...,
						).
						ThenStatic(routeInstance.Handler())
				}

				this.events.Dispatch(&ResponseBefore{request})

				HandleResponse(result, request)

				return nil
			})
		})(routeItem)
	}
}
