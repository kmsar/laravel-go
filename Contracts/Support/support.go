package Support

import (
	"reflect"
	"sort"
)

type Context interface {

	// Get retrieves data from the context.
	Get(key string) interface{}

	// Set saves data in the context.
	Set(key string, val interface{})
}

type Fields map[string]interface{}

type Interface interface {
	reflect.Type
	GetType() reflect.Type
	IsSubClass(class interface{}) bool
	ClassName() string
}

type Class interface {
	Interface
	New(fields Fields) interface{}
	NewByTag(fields Fields, tag string) interface{}
}

type FieldsProvider interface {
	Fields() Fields
}

type Json interface {
	ToJson() string
}

type Getter interface {
	GetString(key string) string
	GetInt64(key string) int64
	GetInt(key string) int
	GetFloat64(key string) float64
	GetFloat(key string) float32
	GetBool(key string) bool
	GetFields(key string) Fields
}

type OptionalGetter interface {
	StringOption(key string, defaultValue string) string
	Int64Option(key string, defaultValue int64) int64
	IntOption(key string, defaultValue int) int
	Float64Option(key string, defaultValue float64) float64
	FloatOption(key string, defaultValue float32) float32
	BoolOption(key string, defaultValue bool) bool
	FieldsOption(key string, defaultValue Fields) Fields
}

type Collection interface {
	Json
	// sort

	sort.Interface
	Sort(sorter interface{}) Collection
	IsEmpty() bool

	// filter

	Map(filter interface{}) Collection
	Filter(filter interface{}) Collection
	Skip(filter interface{}) Collection
	Where(field string, args ...interface{}) Collection
	WhereLt(field string, arg interface{}) Collection
	WhereLte(field string, arg interface{}) Collection
	WhereGt(field string, arg interface{}) Collection
	WhereGte(field string, arg interface{}) Collection
	WhereIn(field string, arg interface{}) Collection
	WhereNotIn(field string, arg interface{}) Collection

	// keys、values

	// Pluck []map、[]struct
	Pluck(key string) Fields
	// Only []map、[]struct
	Only(keys ...string) Collection

	// First
	First(keys ...string) interface{}
	// Last
	Last(keys ...string) interface{}

	// union、merge...

	// Prepend
	Prepend(item ...interface{}) Collection
	// Push
	Push(items ...interface{}) Collection
	// Pull
	Pull(defaultValue ...interface{}) interface{}
	// Shift
	Shift(defaultValue ...interface{}) interface{}
	// Put
	Put(index int, item interface{}) Collection
	// Offset
	Offset(index int, item interface{}) Collection
	// Merge
	Merge(collections ...Collection) Collection
	// Reverse
	Reverse() Collection
	// Chunk
	Chunk(size int, handler func(collection Collection, page int) error) error
	// Random
	Random(size ...uint) Collection

	// aggregate

	Sum(key ...string) float64
	Max(key ...string) float64
	Min(key ...string) float64
	Avg(key ...string) float64
	Count() int

	// convert

	ToIntArray() (results []int)
	ToInt64Array() (results []int64)
	ToInterfaceArray() []interface{}
	ToFloat64Array() (results []float64)
	ToFloatArray() (results []float32)
	ToBoolArray() (results []bool)
	ToStringArray() (results []string)
	ToFields() Fields
	ToArrayFields() []Fields
}
