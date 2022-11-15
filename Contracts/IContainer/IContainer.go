package IContainer

import (
	"reflect"
)

type IContainer interface {
	Singleton(resolver interface{})
	// Register a shared binding in the container.
	Transient(resolver interface{})
	//egister a binding with the container.
	Scoped(resolver interface{})
	Bind(interface{})

	NamedSingleton(key string, resolver interface{})
	// Register a shared binding in the container.
	NamedTransient(key string, resolver interface{})
	//egister a binding with the container.
	NamedScoped(key string, resolver interface{})
	NamedBind(string, interface{})
	// Register an existing instance as shared in the container.
	Instance(key string, instance interface{})

	// Determine whether to bind.
	HasBound(key string) bool

	// alias a type to a different name.
	Alias(key string, value string)
	// flush the container of all bindings and resolved instances.
	Flush()
	FlushScopedInstances()
	// get the given type from the container.
	Get(key string, args ...interface{}) interface{}
	GetTypeName(res any) string

	// call the given fn / class@method and inject its dependencies.
	Call(fn interface{}, args ...interface{}) []interface{}
	// call the given magical func / class@method and inject its dependencies.
	StaticCall(fn MagicalFunc, args ...interface{}) []interface{}
	// injects the given type from the container.
	DI(object interface{}, args ...interface{})
	Resolve(object interface{}, args ...interface{})

	Fill(structure interface{}, args ...interface{})

	//	 ContainerIf(abstract, concrete = null);
	// instance(abstract, instance);
	// when(concrete);
	// factory(abstract);

	// beforeResolving(abstract, Closure callback = null);
	// resolving(abstract, Closure callback = null);
	//afterResolving(abstract, Closure callback = null);

}

// Component injectable class.
type Component interface {
	Construct(container IContainer)
}

// MagicalFunc Magic methods that can be called from the container.
type MagicalFunc interface {

	// NumOut number of output parameters.
	NumOut() int

	// NumIn number of input parameters.
	NumIn() int

	// Call transfer.
	Call(in []reflect.Value) []reflect.Value

	// Arguments get all parameters.
	Arguments() []reflect.Type

	// Returns get all return types.
	Returns() []reflect.Type
}

// container instance provider
type InstanceProvider func() interface{}
