package Utils

import (
	"errors"
	"github.com/kmsar/laravel-go/Framework/Contracts/IContainer"
	"github.com/kmsar/laravel-go/Framework/Support/Utils/Convert"
	"reflect"
	"strings"
)

var (
	UnCallable = errors.New("UnCallable！")
)

func GetCallable(arg interface{}) (reflect.Value, error) {
	argValue := reflect.ValueOf(arg)

	if argValue.Kind() == reflect.Func {
		return argValue, nil
	}

	return argValue, UnCallable
}

// IsSameStruct
func IsSameStruct(v1, v2 interface{}) bool {
	var (
		f1 reflect.Type
		f2 reflect.Type
		ok bool
	)

	if f1, ok = v1.(reflect.Type); !ok {
		f1 = reflect.TypeOf(v1)
	}

	if f2, ok = v2.(reflect.Type); !ok {
		f2 = reflect.TypeOf(v2)
	}

	return f1.PkgPath() == f2.PkgPath() && f1.Name() == f2.Name()
}

// ConvertToTypes 把变量转换成反射类型
func ConvertToTypes(args ...interface{}) []reflect.Type {
	types := make([]reflect.Type, 0)
	for _, arg := range args {
		types = append(types, reflect.TypeOf(arg))
	}
	return types
}

// IsInstanceIn InstanceIn
func IsInstanceIn(v interface{}, types ...reflect.Type) bool {
	for _, e := range types {
		if IsSameStruct(e, v) {
			return true
		}
	}
	return false
}

// EachStructField
func EachStructField(value reflect.Value, data interface{}, handler func(reflect.StructField, reflect.Value)) {
	dataType := reflect.TypeOf(data)

	for i := 0; i < dataType.NumField(); i++ {
		handler(dataType.Field(i), value.Field(i))
	}
}

// EachSlice
func EachSlice(value reflect.Value, handler func(int, reflect.Value)) {
	sliceLen := value.Len()

	for i := 0; i < sliceLen; i++ {
		handler(i, value.Index(i))
	}
}

// EachMap  map
func EachMap(value reflect.Value, handler func(key reflect.Value, value reflect.Value)) {
	for _, key := range value.MapKeys() {
		handler(key, value.MapIndex(key))
	}
}

// GetTypeKey
func GetTypeKey(p reflect.Type) string {
	if p.Kind() == reflect.Ptr {
		p = p.Elem()
	}

	pkgPath := p.PkgPath()

	if pkgPath == "" {
		return "" + p.Name()
	}
	return pkgPath + "." + p.Name()

}

// NotNil  nil
func NotNil(args ...interface{}) interface{} {
	for _, arg := range args {
		switch argValue := arg.(type) {
		case IContainer.InstanceProvider:
			arg = argValue()
		case func() interface{}:
			arg = argValue()
		}
		if arg != nil {
			return arg
		}
	}
	return nil
}

// ParseStructTag
func ParseStructTag(rawTag reflect.StructTag) map[string][]string {
	results := make(map[string][]string, 0)
	for _, tagString := range strings.Split(string(rawTag), " ") {
		tag := strings.Split(tagString, ":")
		if len(tag) > 1 {
			results[tag[0]] = strings.Split(strings.ReplaceAll(tag[1], `"`, ""), ",")
		} else {
			results[tag[0]] = nil
		}
	}
	return results
}

// ConvertToValue  interface  reflect.Value
func ConvertToValue(argType reflect.Type, arg interface{}) reflect.Value {
	switch argType.Kind() {
	case reflect.String:
		return reflect.ValueOf(Convert.ConvertToString(arg, ""))
	case reflect.Int:
		return reflect.ValueOf(Convert.ConvertToInt(arg, 0))
	case reflect.Int64:
		return reflect.ValueOf(Convert.ConvertToInt64(arg, 0))
	case reflect.Float64:
		return reflect.ValueOf(Convert.ConvertToFloat64(arg, 0))
	case reflect.Float32:
		return reflect.ValueOf(Convert.ConvertToFloat(arg, 0))
	case reflect.Bool:
		return reflect.ValueOf(Convert.ConvertToBool(arg, false))
	}
	if reflect.TypeOf(arg).ConvertibleTo(argType) {
		return reflect.ValueOf(arg).Convert(argType)
	}
	return reflect.ValueOf(arg)
}
