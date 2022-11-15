package Field

import (
	"errors"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Str"
	"github.com/kmsar/laravel-go/Framework/Support/Utils"
	"github.com/kmsar/laravel-go/Framework/Support/Utils/Convert"
	"reflect"
	"strings"
)

func OnlyFields(fields Support.Fields, keys ...string) Support.Fields {
	var results = make(Support.Fields)

	for _, key := range keys {
		results[key] = fields[key]
	}

	return results
}

func MakeKeysMap(keys ...string) Support.Fields {
	var keysMap = Support.Fields{}
	for _, key := range keys {
		keysMap[key] = 1
	}

	return keysMap
}

func ExceptFields(fields Support.Fields, keys ...string) Support.Fields {
	var (
		results = make(Support.Fields)
		keysMap = MakeKeysMap(keys...)
	)

	for key, value := range fields {
		if _, exists := keysMap[key]; !exists {
			results[key] = value
		}
	}

	return results
}

func OnlyExistsFields(fields Support.Fields, keys ...string) Support.Fields {
	var results = make(Support.Fields)

	for _, key := range keys {
		if value := fields[key]; value != nil {
			results[key] = value
		}
	}

	return results
}

// MergeFields 合并两个 contracts.Fields
func MergeFields(fields Support.Fields, finalFields Support.Fields) {
	for key, value := range finalFields {
		fields[key] = value
	}
}

// GetStringField 获取 Fields 中的字符串，会尝试转换类型
func GetStringField(fields Support.Fields, key string, defaultValues ...string) string {
	if value, existsString := fields[key]; existsString {
		if str, isString := value.(string); isString {
			return str
		}
	}
	return Str.StringOr(defaultValues...)
}

// GetSubField 获取下级 Fields ，如果没有的话，匹配同前缀的放到下级 Fields 中
func GetSubField(fields Support.Fields, key string, defaultValues ...Support.Fields) Support.Fields {

	if subField, isField := fields[key].(Support.Fields); isField {
		return subField
	}

	if len(defaultValues) > 0 {
		return defaultValues[0]
	}

	subField := make(Support.Fields)
	prefix := key + "."

	for fieldKey, fieldValue := range fields {
		if strings.HasPrefix(fieldKey, prefix) {
			subField[strings.ReplaceAll(fieldKey, prefix, "")] = fieldValue
		}
	}

	if len(subField) > 0 {
		fields[key] = subField
	}

	return subField
}

// GetInt64Field 获取 Fields 中的 int64，会尝试转换类型
func GetInt64Field(fields Support.Fields, key string, defaultValues ...int64) int64 {
	var defaultValue int64 = 0
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if value, existsValue := fields[key]; existsValue {
		if intValue, isInt := value.(int64); isInt {
			return intValue
		}
		return Convert.ConvertToInt64(value, defaultValue)
	} else {
		return defaultValue
	}
}

// GetIntField 获取 Fields 中的 int，会尝试转换类型
func GetIntField(fields Support.Fields, key string, defaultValues ...int) int {
	var defaultValue = 0
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if value, existsValue := fields[key]; existsValue {
		if intValue, isInt := value.(int); isInt {
			return intValue
		}
		return int(Convert.ConvertToInt64(value, int64(defaultValue)))
	} else {
		return defaultValue
	}
}

// GetFloatField 获取 Fields 中的 float32，会尝试转换类型
func GetFloatField(fields Support.Fields, key string, defaultValues ...float32) float32 {
	var defaultValue float32 = 0
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if value, existsValue := fields[key]; existsValue {
		if intValue, isInt := value.(float32); isInt {
			return intValue
		}
		return Convert.ConvertToFloat(value, defaultValue)
	} else {
		return defaultValue
	}
}

// GetFloat64Field 获取 Fields 中的 float64，会尝试转换类型
func GetFloat64Field(fields Support.Fields, key string, defaultValues ...float64) float64 {
	var defaultValue float64 = 0
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if value, existsValue := fields[key]; existsValue {
		if intValue, isInt := value.(float64); isInt {
			return intValue
		}
		return Convert.ConvertToFloat64(value, defaultValue)
	} else {
		return defaultValue
	}
}

// GetBoolField 获取 Fields 中的 bool，会尝试转换类型
func GetBoolField(fields Support.Fields, key string, defaultValues ...bool) bool {
	var defaultValue = false
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if fieldValue, existsValue := fields[key]; existsValue {
		return Convert.ConvertToBool(fieldValue, defaultValue)
	}
	return defaultValue
}

// ConvertToFields 尝试把一个变量转换成 Fields 类型
func ConvertToFields(anyValue interface{}) (Support.Fields, error) {
	fields := Support.Fields{}
	switch paramValue := anyValue.(type) {
	case Support.Fields:
		fields = paramValue
	case map[string]interface{}:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]int:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]int8:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]int16:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]int32:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]int64:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]uint:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]uint8:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]uint16:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]uint32:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]uint64:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]float64:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]float32:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]string:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]bool:
		for key, value := range paramValue {
			fields[key] = value
		}
	default:
		paramType := reflect.ValueOf(anyValue)

		switch paramType.Kind() {
		case reflect.Struct: // 结构体
			Utils.EachStructField(paramType, anyValue, func(field reflect.StructField, value reflect.Value) {
				if field.IsExported() {
					fields[Str.SnakeString(field.Name)] = value.Interface()
				} else {
					fields[Str.SnakeString(field.Name)] = nil
				}
			})
		case reflect.Map: // 自定义的 map
			for _, key := range paramType.MapKeys() {
				name := key.String()
				fields[name] = paramType.MapIndex(key).Interface()
			}
		default:
			return nil, errors.New("不支持转 Support.Fields 的类型： " + paramType.String())
		}
	}
	return fields, nil
}
