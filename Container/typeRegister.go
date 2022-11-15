package Container

import (
	"errors"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Exceptions"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Utils"
	"reflect"
	"sync"
)

type TypeRegister struct {
	types sync.Map
}

func (t *TypeRegister) SetOrGet(refType reflect.Type) string {
	name := Utils.GetTypeKey(refType)
	if typ, ok := t.types.Load(name); ok {
		if typ != refType {
			Exceptions.Throw("another type with same name exist")
		}
		return name
	}
	t.types.Store(name, refType)
	return name
}

func (t *TypeRegister) Get(name string) (reflect.Type, error) {
	if typ, ok := t.types.Load(name); ok {
		return typ.(reflect.Type), nil
	}
	return nil, errors.New("no one")
}

func (t *TypeRegister) getName(refType reflect.Type) string {
	var name string
	t.types.Range(func(k, v interface{}) bool {
		if refType == v.(reflect.Type) {
			name = k.(string)
			return false
		}
		return true
	})
	return name
}
