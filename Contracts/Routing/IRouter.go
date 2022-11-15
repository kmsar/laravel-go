package Routing

type Route interface {

	// Get the middlewares attached to the route.
	Middlewares() []MagicalFunc

	// Get the request method attached to the route.
	Method() []string

	// Get the request path attached to the route.
	Path() string

	// Get the route handler attached to the route.
	Handler() MagicalFunc
}

type RouteGroup interface {

	// Register a new GET route with the route group.
	Get(path string, handler interface{}, middlewares ...interface{}) RouteGroup

	// Register a new POST route with the route group.
	Post(path string, handler interface{}, middlewares ...interface{}) RouteGroup

	// Register a new DELETE route with the routing group.
	Delete(path string, handler interface{}, middlewares ...interface{}) RouteGroup

	// Register a new PUT route with the routing group.
	Put(path string, handler interface{}, middlewares ...interface{}) RouteGroup

	// Register a new PATCH route with the routing group.
	Patch(path string, handler interface{}, middlewares ...interface{}) RouteGroup

	// Register a new OPTIONS route with the routing group.
	Options(path string, handler interface{}, middlewares ...interface{}) RouteGroup

	// Register a new TRACE route with the routing group.
	Trace(path string, handler interface{}, middlewares ...interface{}) RouteGroup

	// Get the middlewares attached to the route.
	Middlewares() []MagicalFunc

	// Create a route group with shared attributes.
	Group(prefix string, middlewares ...interface{}) RouteGroup

	// get route.
	Routes() []Route

	// get routing group.
	Groups() []RouteGroup
}

type Router interface {
	Static(path string, directory string)

	// Register a new GET route with the router.
	Get(path string, handler interface{}, middlewares ...interface{})

	// Register a new POST route with the router.
	Post(path string, handler interface{}, middlewares ...interface{})

	// Register a new DELETE route with the router.
	Delete(path string, handler interface{}, middlewares ...interface{})

	// Register a new PUT route with the router.
	Put(path string, handler interface{}, middlewares ...interface{})

	// Register a new PATCH route with the router.
	Patch(path string, handler interface{}, middlewares ...interface{})

	// Register a new OPTIONS route with the router.
	Options(path string, handler interface{}, middlewares ...interface{})

	// Register a new TRACE route with the router.
	Trace(path string, handler interface{}, middlewares ...interface{})

	// use middleware.
	Use(middlewares ...interface{})

	// Create a route group with shared attributes.
	Group(prefix string, middlewares ...interface{}) RouteGroup

	// start httpserver.
	Start(address string) error

	// close httpserver.
	Close() error
}
