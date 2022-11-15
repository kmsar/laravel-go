package Collections

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Exceptions"
	"github.com/kmsar/laravel-go/Framework/Support/Field"
	"github.com/kmsar/laravel-go/Framework/Support/Utils"
	"github.com/kmsar/laravel-go/Framework/Support/Utils/Convert"
	"reflect"
)

type array []interface{}

type Collection struct {
	array
	mapData []Support.Fields
	sorter  func(i, j int) bool
}

func New(data interface{}) (*Collection, error) {
	dataValue := reflect.ValueOf(data)

	if dataValue.Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
	}

	switch dataValue.Kind() {
	case reflect.Array, reflect.Slice:
		collect := &Collection{mapData: nil, array: make([]interface{}, 0)}
		isMapData := true

		Utils.EachSlice(dataValue, func(_ int, value reflect.Value) {
			switch value.Kind() {
			case reflect.Map, reflect.Struct:
				collect.array = append(collect.array, value.Interface())
			default:
				isMapData = false
				collect.array = append(collect.array, value.Interface())
			}
		})

		if isMapData {
			collect.mapData = make([]Support.Fields, 0)
			for _, item := range collect.array {
				fields, _ := Field.ConvertToFields(item)
				collect.mapData = append(collect.mapData, fields)
			}
		}

		return collect, nil
	}

	return nil, Exception{
		Exception: Exceptions.New("Exceptions"+Utils.GetTypeKey(reflect.TypeOf(data)), Support.Fields{
			"data": data,
		}),
	}
}

func MustNew(data interface{}) Support.Collection {
	c, err := New(data)
	Exceptions.Throw(err)
	return c
}

func FromFieldsSlice(data []Support.Fields) Support.Collection {
	collection := &Collection{mapData: data}
	for _, fields := range data {
		collection.array = append(collection.array, fields)
	}
	return collection
}

func (this *Collection) Index(index int) interface{} {
	if this.Count() > index {
		return this.array[index]
	}
	return nil
}

func (this *Collection) IsEmpty() bool {
	return len(this.array) == 0
}
func (this *Collection) Each(handler interface{}) Support.Collection {
	return this.Map(handler)
}

func (this *Collection) Map(handler interface{}) Support.Collection {

	switch handle := handler.(type) {
	case func(fields Support.Fields):
		for _, fields := range this.mapData {
			handle(fields)
		}
	case func(fields Support.Fields, index int): // 聚合函数
		for index, fields := range this.mapData {
			handle(fields, index)
		}
	case func(fields Support.Fields) Support.Fields:
		for index, fields := range this.mapData {
			this.mapData[index] = handle(fields)
		}
	case func(fields Support.Fields) bool: // 用于 where
		results := make([]interface{}, 0)
		for _, fields := range this.mapData {
			results = append(results, handle(fields))
		}
		return &Collection{mapData: nil, array: results}
	case func(fields interface{}):
		for _, data := range this.array {
			handle(data)
		}
	case func(fields interface{}) interface{}:
		for index, data := range this.array {
			this.array[index] = handle(data)
		}
	default:
		handlerType := reflect.TypeOf(handler)
		handlerValue := reflect.ValueOf(handler)
		argsGetter := func(index int, arg interface{}) []reflect.Value {
			return []reflect.Value{}
		}
		numOut := handlerType.NumOut()

		switch handlerType.NumIn() {
		case 1:
			argsGetter = func(_ int, arg interface{}) []reflect.Value {
				return []reflect.Value{this.argumentConvertor(handlerType.In(0), arg)}
			}
		case 2:
			argsGetter = func(index int, arg interface{}) []reflect.Value {
				return []reflect.Value{
					this.argumentConvertor(handlerType.In(0), arg),
					reflect.ValueOf(index),
				}
			}
		case 3:
			array := reflect.ValueOf(this.array)
			argsGetter = func(index int, arg interface{}) []reflect.Value {
				return []reflect.Value{
					this.argumentConvertor(handlerType.In(0), arg),
					reflect.ValueOf(index),
					array,
				}
			}
		}

		if numOut > 0 {
			mapLen := len(this.mapData)
			newCollection := &Collection{mapData: make([]Support.Fields, mapLen), array: make([]interface{}, 0)}
			for index, data := range this.array {
				result := handlerValue.Call(argsGetter(index, data))[0].Interface()
				if mapLen > 0 {
					newCollection.mapData[index], _ = Field.ConvertToFields(result)
				}
				newCollection.array = append(newCollection.array, result)
			}
			return newCollection
		} else {
			for index, data := range this.array {
				handlerValue.Call(argsGetter(index, data))
			}
		}

	}

	return this
}

func (this *Collection) argumentConvertor(argType reflect.Type, arg interface{}) reflect.Value {
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
