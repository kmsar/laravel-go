package class

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Utils"
	"github.com/pkg/errors"
	"reflect"
)

type Interface struct {
	reflect.Type

	fields map[string]reflect.StructField
}

func (i *Interface) GetType() reflect.Type {
	return i.Type
}

func Define(arg interface{}) Support.Interface {
	argType := reflect.TypeOf(arg)
	if argType.Kind() == reflect.Ptr {
		argType = argType.Elem()
	}
	class := &Interface{Type: argType}
	if argType.Kind() != reflect.Interface {
		panic(TypeException{
			errors.New(" interface type exception"),
			map[string]interface{}{
				"class": class.ClassName(),
			},
		})
	}
	return class
}

func (i *Interface) ClassName() string {
	return Utils.GetTypeKey(i)
}

func (i *Interface) IsSubClass(class interface{}) bool {
	if value, ok := class.(reflect.Type); ok {
		return value.ConvertibleTo(i.Type)
	}

	return reflect.TypeOf(class).ConvertibleTo(i.Type)
}

func (i *Interface) Implements(class reflect.Type) bool {
	switch value := class.(type) {
	case *Interface:
		return i.Type.Implements(value.Type)
	case *Class:
		return false
	}

	return i.Type.Implements(class)
}
