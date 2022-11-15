package Foundation

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Container"
	"github.com/kmsar/laravel-go/Framework/Contracts/IContainer"
)

type IApplication interface {

	// Container returns the service container
	Container() Container.Container

	// SetContainer set the service container
	SetContainer(container Container.Container)

	//Instance(concrete interface{}) interface{}

	IsEnvironment(environments ...string) bool

	//// The Log method gives you an instance of a logger. You can write your log
	//// messages to this instance.
	//Log(channels ...string) LoggerFacade
	//
	//// Db returns the database facade. If no parameters are provided, it will use
	//// the default connection.
	//Db(connection ...string) Database

	IContainer.Container

	// IsProduction Determine if it is a production environment.
	IsProduction() bool

	// Debug Determine whether to enable debugging.
	// If debugging is enabled, the log will print some debugging information.
	Debug() bool

	// Environment Get the current operating environment.
	Environment() (string, error)

	// RegisterServices Register the application service.
	RegisterServices(provider ...ServiceProvider)

	// Start application start.
	Start() map[string]error

	// Stop application stop.
	Stop()
}
