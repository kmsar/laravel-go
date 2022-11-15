package IHashing

import "github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"

type HasherProvider func(config Support.Fields) Hasher

type HasherFactory interface {
	Hasher

	// Driver 获取驱动程序实例
	// Get a driver instance.
	Driver(driver string) Hasher

	// Extend Register a custom driver creator Closure.
	Extend(driver string, hasherProvider HasherProvider)
}

type Hasher interface {

	// Info Get information about the given hashed value.
	Info(hashedValue string) Support.Fields

	// Make Hash the given value.
	Make(value string, options Support.Fields) string

	// Check  the given plain value against a hash.
	Check(value, hashedValue string, options Support.Fields) bool
}
