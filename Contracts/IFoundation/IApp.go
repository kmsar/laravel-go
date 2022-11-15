package IFoundation

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IContainer"
)

type IApplication interface {
	//
	//// Container returns the service container
	//Container() IContainer.IContainer
	//
	//// SetContainer set the service container
	//SetContainer(container IContainer.IContainer)

	//Instance(concrete interface{}) interface{}

	IsEnvironment(environments ...string) bool

	//// The Log method gives you an instance of a logger. You can write your log
	//// messages to this instance.
	//Log(channels ...string) LoggerFacade
	//
	//// Db returns the database facade. If no parameters are provided, it will use
	//// the default connection.
	//Db(connection ...string) Database

	IContainer.IContainer

	// IsProduction Determine if it is a production environment.
	IsProduction() bool

	// Debug Determine whether to enable debugging.
	// If debugging is enabled, the log will print some debugging information.
	Debug() bool

	// Environment Get the current operating environment.
	Environment() string

	// RegisterServices Register the application service.
	RegisterServices(provider ...ServiceProvider)

	// Start application start.
	Start() map[string]error

	// Stop application stop.
	Stop()
}
