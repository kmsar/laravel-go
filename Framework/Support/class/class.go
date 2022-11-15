package class

import (
	"errors"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Utils"
	"reflect"
	"sync"
)

type Class struct {
	reflect.Type

	fields sync.Map
}

func (this *Class) NewByTag(data Support.Fields, tag string) interface{} {
	object := reflect.New(this.Type).Elem()

	if data != nil {
		jsonFields := this.getFields("json")
		targetFields := this.getFields(tag)
		for key, value := range data {
			if field, ok := targetFields[key]; ok && field.IsExported() {
				object.FieldByIndex(field.Index).Set(Utils.ConvertToValue(field.Type, value))
			} else if field, ok = jsonFields[key]; ok && field.IsExported() {
				object.FieldByIndex(field.Index).Set(Utils.ConvertToValue(field.Type, value))
			}
		}
	}

	return object.Interface()
}

func Make(arg interface{}) Support.Class {
	argType := reflect.TypeOf(arg)
	if argType.Kind() == reflect.Ptr {
		argType = argType.Elem()
	}
	class := &Class{Type: argType}
	if argType.Kind() != reflect.Struct {
		panic(TypeException{
			errors.New("只支持 struct 类型"),
			map[string]interface{}{
				"class": class.ClassName(),
			},
		})
	}
	return class
}

func (this *Class) New(data Support.Fields) interface{} {
	return this.NewByTag(data, "json")
}

func (this *Class) getFields(tag string) map[string]reflect.StructField {
	data, exists := this.fields.Load(tag)

	if !exists {
		var fields = map[string]reflect.StructField{}
		for i := 0; i < this.Type.NumField(); i++ {
			field := this.Type.Field(i)
			tags := Utils.ParseStructTag(field.Tag)
			if target := tags[tag]; target != nil && len(target) > 0 {
				fields[target[0]] = field
			} else {
				fields[field.Name] = field
			}
		}

		this.fields.Store(tag, fields)
		return fields
	}

	return data.(map[string]reflect.StructField)
}

func (this *Class) ClassName() string {
	return Utils.GetTypeKey(this)
}

func (this *Class) GetType() reflect.Type {
	return this.Type
}

func (this *Class) IsSubClass(class interface{}) bool {
	if value, ok := class.(reflect.Type); ok {
		return value.ConvertibleTo(this.Type)
	}

	return reflect.TypeOf(class).ConvertibleTo(this.Type)
}

func (this *Class) Implements(class reflect.Type) bool {
	switch value := class.(type) {
	case *Interface:
		return this.Type.Implements(value.Type)
	case *Class:
		return this.Type.Implements(value.Type)
	}

	return this.Type.Implements(class)
}
