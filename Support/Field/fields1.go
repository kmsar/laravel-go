package Field

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Utils/Convert"
	"strings"
)

type InstanceGetter func(key string) interface{}

type BaseFields struct {
	Support.FieldsProvider
	Getter InstanceGetter
}

func (f *BaseFields) get(key string) interface{} {
	if f.Getter != nil {
		if value := f.Getter(key); value != nil && value != "" {
			return value
		}
	}
	return f.Fields()[key]
}

func (f *BaseFields) Only(keys ...string) Support.Fields {
	var fields = make(Support.Fields)

	for _, key := range keys {
		if value := f.get(key); value != nil {
			fields[key] = value
		}
	}

	return fields
}

func (f *BaseFields) ExceptFields(keys ...string) Support.Fields {
	var (
		results = make(Support.Fields)
		keysMap = MakeKeysMap(keys...)
	)

	for key, value := range f.Fields() {
		if _, exists := keysMap[key]; !exists {
			results[key] = value
		}
	}

	return results
}

func (f *BaseFields) OnlyExists(keys ...string) Support.Fields {
	var fields = make(Support.Fields)

	for _, key := range keys {
		if value := f.get(key); value != nil {
			fields[key] = value
		}
	}

	return fields
}

func (f *BaseFields) StringOption(key string, defaultValue string) string {
	if value := f.get(key); value != nil && value != "" {
		return Convert.ConvertToString(value, defaultValue)
	}
	return defaultValue
}

func (f *BaseFields) Int64Option(key string, defaultValue int64) int64 {
	if value := f.get(key); value != nil && value != "" {
		return Convert.ConvertToInt64(value, defaultValue)
	}
	return defaultValue
}

func (f *BaseFields) IntOption(key string, defaultValue int) int {
	if value := f.get(key); value != nil && value != "" {
		return Convert.ConvertToInt(value, defaultValue)
	}
	return defaultValue
}

func (f *BaseFields) Float64Option(key string, defaultValue float64) float64 {
	if value := f.get(key); value != nil && value != "" {
		return Convert.ConvertToFloat64(value, defaultValue)
	}
	return defaultValue
}

func (f *BaseFields) FloatOption(key string, defaultValue float32) float32 {
	if value := f.get(key); value != nil && value != "" {
		return Convert.ConvertToFloat(value, defaultValue)
	}
	return defaultValue
}

func (f *BaseFields) BoolOption(key string, defaultValue bool) bool {
	if value := f.get(key); value != nil && value != "" {
		return Convert.ConvertToBool(value, defaultValue)
	}
	return defaultValue
}

func (f *BaseFields) FieldsOption(key string, defaultValue Support.Fields) Support.Fields {
	if value := f.get(key); value != nil && value != "" {
		if fields, err := ConvertToFields(value); err == nil {
			return fields
		}
	}
	fields := Support.Fields{}
	for fieldKey, value := range f.Fields() {
		if strings.HasPrefix(fieldKey, key+".") {
			fields[strings.ReplaceAll(fieldKey, key+".", "")] = value
		}
	}
	if len(fields) > 0 {
		return fields
	}
	return defaultValue
}

func (f *BaseFields) GetString(key string) string {
	return f.StringOption(key, "")
}

func (f *BaseFields) GetInt64(key string) int64 {
	return f.Int64Option(key, 0)
}

func (f *BaseFields) GetInt(key string) int {
	return f.IntOption(key, 0)
}

func (f *BaseFields) GetFloat64(key string) float64 {
	return f.Float64Option(key, 0)
}

func (f *BaseFields) GetFloat(key string) float32 {
	return f.FloatOption(key, 0)
}

func (f *BaseFields) GetBool(key string) bool {
	return f.BoolOption(key, false)
}

func (f *BaseFields) GetFields(key string) Support.Fields {
	return f.FieldsOption(key, Support.Fields{})
}
