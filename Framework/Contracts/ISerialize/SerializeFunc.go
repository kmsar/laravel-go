package ISerialize

import "github.com/kmsar/laravel-go/Framework/Contracts/Support"

type Serialization interface {

	// Method Get the serialized driver based on the given name
	Method(name string) Serializer

	// Extend Add serialization driver.
	Extend(name string, serializer Serializer)
}

type Serializer interface {
	// Serialize serialize the given data.
	Serialize(interface{}) string
	UnSerialize(string, interface{}) error
}

type ClassSerializer interface {

	// Register register parsing class.
	Register(class Support.Class)

	// Serialize serialize the given data.
	Serialize(interface{}) string

	// Parse  the serialized string.
	Parse(string) (interface{}, error)
}
