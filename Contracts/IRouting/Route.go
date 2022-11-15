package IRouting

import "github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IContainer"

type Route interface {

	// Middlewares Get the middlewares attached to the route.
	Middlewares() []IContainer.MagicalFunc

	// Method Get the request method attached to the route.
	Method() []string

	// Path Get the request path attached to the route.
	Path() string

	// Handler Get the route handler attached to the route.
	Handler() IContainer.MagicalFunc
}

type RouteGroup interface {

	// Get Register a new GET route with the route group.
	Get(path string, handler interface{}, middlewares ...interface{}) RouteGroup

	// Post Register a new POST route with the route group.
	Post(path string, handler interface{}, middlewares ...interface{}) RouteGroup

	// Delete Register a new DELETE route with the routing group.
	Delete(path string, handler interface{}, middlewares ...interface{}) RouteGroup

	// Put Register a new PUT route with the routing group.
	Put(path string, handler interface{}, middlewares ...interface{}) RouteGroup

	// Patch Register a new PATCH route with the routing group.
	Patch(path string, handler interface{}, middlewares ...interface{}) RouteGroup

	// Options Register a new OPTIONS route with the routing group.
	Options(path string, handler interface{}, middlewares ...interface{}) RouteGroup

	// Trace Register a new TRACE route with the routing group.
	Trace(path string, handler interface{}, middlewares ...interface{}) RouteGroup

	// Middlewares Get the middlewares attached to the route.
	Middlewares() []IContainer.MagicalFunc

	// Group Create a route group with shared attributes.
	Group(prefix string, middlewares ...interface{}) RouteGroup

	// Routes get route.
	Routes() []Route

	// Groups get routing group.
	Groups() []RouteGroup
}

type Router interface {
	Static(path string, directory string)

	// Get Register a new GET route with the router.
	Get(path string, handler interface{}, middlewares ...interface{})

	// Post Register a new POST route with the router.
	Post(path string, handler interface{}, middlewares ...interface{})

	// Delete Register a new DELETE route with the router.
	Delete(path string, handler interface{}, middlewares ...interface{})

	// Put Register a new PUT route with the router.
	Put(path string, handler interface{}, middlewares ...interface{})

	// Patch Register a new PATCH route with the router.
	Patch(path string, handler interface{}, middlewares ...interface{})

	// Options Register a new OPTIONS route with the router.
	Options(path string, handler interface{}, middlewares ...interface{})

	// Trace Register a new TRACE route with the router.
	Trace(path string, handler interface{}, middlewares ...interface{})

	// Use use middleware.
	Use(middlewares ...interface{})

	// Group Create a route group with shared attributes.
	Group(prefix string, middlewares ...interface{}) RouteGroup

	// Start start httpserver.
	Start(address string) error

	// Close close httpserver.
	Close() error
}
