package Collections

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Utils"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Utils/Convert"
)

func (this *Collection) Filter(filter interface{}) Support.Collection {
	results := make([]interface{}, 0)
	newFields := make([]Support.Fields, 0)
	for index, data := range this.Map(filter).ToInterfaceArray() {
		if Convert.ConvertToBool(data, false) {
			if fields := this.mapData[index]; fields != nil {
				newFields = append(newFields, fields)
			}
			results = append(results, this.array[index])
		}
	}
	return &Collection{
		mapData: newFields,
		array:   results,
	}
}

func (this *Collection) Skip(filter interface{}) Support.Collection {
	results := make([]interface{}, 0)
	newFields := make([]Support.Fields, 0)
	for index, data := range this.Map(filter).ToInterfaceArray() {
		if !Convert.ConvertToBool(data, false) {
			if fields := this.mapData[index]; fields != nil {
				newFields = append(newFields, fields)
			}
			results = append(results, this.array[index])
		}
	}
	return &Collection{
		mapData: newFields,
		array:   results,
	}
}

// Where 根据条件过滤数据，支持 =,>,>=,<,<=,in,not in 等条件判断
func (this *Collection) Where(field string, args ...interface{}) Support.Collection {
	results := make([]interface{}, 0)
	var (
		arg      interface{}
		operator = "="
	)
	switch len(args) {
	case 1:
		arg = args[0]
	case 2:
		operator = args[0].(string)
		arg = args[1]
	}
	newFields := make([]Support.Fields, 0)
	for index, data := range this.Map(func(fields Support.Fields) bool {
		return Utils.Compare(fields[field], operator, arg)
	}).ToInterfaceArray() {
		if Convert.ConvertToBool(data, false) {
			if fields := this.mapData[index]; fields != nil {
				newFields = append(newFields, fields)
			}
			results = append(results, this.array[index])
		}
	}
	return &Collection{
		mapData: newFields,
		array:   results,
	}
}

func (this *Collection) WhereLt(field string, arg interface{}) Support.Collection {
	return this.Where(field, "lt", arg)
}
func (this *Collection) WhereLte(field string, arg interface{}) Support.Collection {
	return this.Where(field, "lte", arg)
}
func (this *Collection) WhereGt(field string, arg interface{}) Support.Collection {
	return this.Where(field, "gt", arg)
}
func (this *Collection) WhereGte(field string, arg interface{}) Support.Collection {
	return this.Where(field, "gte", arg)
}
func (this *Collection) WhereIn(field string, arg interface{}) Support.Collection {
	return this.Where(field, "in", arg)
}
func (this *Collection) WhereNotIn(field string, arg interface{}) Support.Collection {
	return this.Where(field, "not in", arg)
}
