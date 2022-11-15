package IAuth

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

// GuardDriver guard driver
type GuardDriver func(name string, config Support.Fields, ctx Support.Context, provider UserProvider) Guard

// UserProviderDriver User Provider Driver
type UserProviderDriver func(config Support.Fields) UserProvider

type Auth interface {

	// ExtendUserProvider Extended User Provider.
	ExtendUserProvider(name string, provider UserProviderDriver)

	// ExtendGuard Extended guard.
	ExtendGuard(name string, guard GuardDriver)

	// Guard Get a guard instance by name.
	Guard(name string, ctx Support.Context) Guard

	// Get a user provider instance by name.
	UserProvider(name string) UserProvider
}

type Authenticatable interface {
	// GetId Get the ID for the currently authenticated user.
	GetId() string
}

type Guard interface {

	// Once Set the current user.
	Once(user Authenticatable)

	// User 获取当前认证的用户
	// Get the currently authenticated user.
	User() Authenticatable

	// GetId Get the ID for the currently authenticated user.
	GetId() string

	// Check Determine if the current user is authenticated.
	Check() bool

	// Guest Determine if the current user is a guest.
	Guest() bool

	// Login Log a user into the application.
	Login(user Authenticatable) interface{}

	Logout() error

	Error() error
}

type UserProvider interface {

	// RetrieveById Retrieve a user by their unique identifier.
	RetrieveById(identifier string) Authenticatable
}

type Authorizable interface {
	// Can Determine if the entity has a given ability.
	Can(ability string, arguments ...interface{}) bool
}

// GateChecker permission checker.
type GateChecker func(user Authorizable, data ...interface{}) bool

// GateHook permission hook.
type GateHook func(user Authorizable, ability string, data ...interface{}) bool

// Policy Permission policy, a set of checkers.
type Policy map[string]GateChecker

type Gate interface {

	// Allows determined if the given ability should be granted for the current user.
	Allows(ability string, arguments ...interface{}) bool

	// Denies Determine if the given ability should be denied for the current user.
	Denies(ability string, arguments ...interface{}) bool

	// Check Determine if all the given abilities should be granted for the current user.
	Check(abilities []string, arguments ...interface{}) bool

	// Any Determine if any one of the given abilities should be granted for the current user.
	Any(abilities []string, arguments ...interface{}) bool

	// Authorize Determine if the given ability should be granted for the current user.
	Authorize(ability string, arguments ...interface{})

	// Inspect the user for the given ability.
	Inspect(ability string, arguments ...interface{}) IHttp.IHttpResponse

	// ForUser Get a guard instance for the given user.
	ForUser(user Authorizable) Gate
}

type GateFactory interface {

	// Has determined if a given ability has been defined.
	Has(ability string) bool

	// Define a new ability.
	Define(ability string, callback GateChecker) GateFactory

	// Policy define a policy class for a given class type.
	Policy(class Support.Class, policy Policy) GateFactory

	// Before Register a callback to run before all Gate checks.
	Before(callable GateHook) GateFactory

	// After Register a callback to run after all Gate checks.
	After(callable GateHook) GateFactory

	// Check Determine if all the given abilities should be granted for the current user.
	Check(user Authorizable, ability string, arguments ...interface{}) bool

	// Abilities Get all the defined abilities.
	Abilities() []string
}
