package Convert

import (
	"fmt"
	"strconv"
)

func ConvertToInt64(rawValue interface{}, defaultValue int64) int64 {
	switch value := rawValue.(type) {
	case int64:
		return value
	case int:
		return int64(value)
	case uint:
		return int64(value)
	case uint32:
		return int64(value)
	case uint8:
		return int64(value)
	case uint16:
		return int64(value)
	case uint64:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case float64:
		return int64(value)
	case float32:
		return int64(value)
	case bool:
		if value {
			return 1
		} else {
			return 0
		}
	case []rune:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return i64
	case []byte:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return i64
	case string:
		i64, _ := strconv.ParseInt(value, 10, 64)
		return i64
	}

	return defaultValue
}

func ConvertToInt(rawValue interface{}, defaultValue int) int {
	switch value := rawValue.(type) {
	case int64:
		return int(value)
	case int:
		return value
	case uint:
		return int(value)
	case uint32:
		return int(value)
	case uint8:
		return int(value)
	case uint16:
		return int(value)
	case uint64:
		return int(value)
	case int8:
		return int(value)
	case int16:
		return int(value)
	case int32:
		return int(value)
	case float64:
		return int(value)
	case float32:
		return int(value)
	case bool:
		if value {
			return 1
		} else {
			return 0
		}
	case []rune:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return int(i64)
	case []byte:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return int(i64)
	case string:
		i64, _ := strconv.ParseInt(value, 10, 32)
		return int(i64)
	}

	return defaultValue
}

func ConvertToFloat64(rawValue interface{}, defaultValue float64) float64 {
	switch value := rawValue.(type) {
	case float64:
		return value
	case int64:
		return float64(value)
	case uint:
		return float64(value)
	case uint32:
		return float64(value)
	case uint8:
		return float64(value)
	case uint16:
		return float64(value)
	case uint64:
		return float64(value)
	case int:
		return float64(value)
	case int8:
		return float64(value)
	case int16:
		return float64(value)
	case int32:
		return float64(value)
	case float32:
		return float64(value)
	case bool:
		if value {
			return 1
		} else {
			return 0
		}
	case []rune:
		f64, _ := strconv.ParseFloat(string(value), 64)
		return f64
	case []byte:
		f64, _ := strconv.ParseFloat(string(value), 64)
		return f64
	case string:
		f64, _ := strconv.ParseFloat(value, 64)
		return f64
	}

	return defaultValue
}

func ConvertToFloat(rawValue interface{}, defaultValue float32) float32 {
	switch value := rawValue.(type) {
	case float64:
		return float32(value)
	case uint:
		return float32(value)
	case uint32:
		return float32(value)
	case uint8:
		return float32(value)
	case uint16:
		return float32(value)
	case uint64:
		return float32(value)
	case int64:
		return float32(value)
	case int:
		return float32(value)
	case int8:
		return float32(value)
	case int16:
		return float32(value)
	case int32:
		return float32(value)
	case float32:
		return value
	case bool:
		if value {
			return 1
		} else {
			return 0
		}
	case string:
		f64, _ := strconv.ParseFloat(value, 32)
		return float32(f64)
	case []rune:
		f64, _ := strconv.ParseFloat(string(value), 32)
		return float32(f64)
	case []byte:
		f64, _ := strconv.ParseFloat(string(value), 32)
		return float32(f64)
	}

	return defaultValue
}

func ConvertToBool(rawValue interface{}, defaultValue bool) bool {
	switch value := rawValue.(type) {
	case bool:
		return value
	case string:
		switch value {
		case "false", "(false)", "0", "":
			return false
		case "true", "(true)", "1":
			return true
		}
	case float64, int, int64, int8, float32:
		return ConvertToInt64(value, 0) > 0 || defaultValue
	}

	return defaultValue
}

func ConvertToString(rawValue interface{}, defaultValue string) string {
	switch value := rawValue.(type) {
	case bool:
		if value {
			return "true"
		} else {
			return "false"
		} //Str.IfString(value, "true", "false")
	case string:
		return value
	case []byte:
		return string(value)
	case []rune:
		return string(value)
	case fmt.Stringer:
		return value.String()
	case int, int64, int8, int32, int16, uint16, uint, uint8, uint32, uint64:
		return fmt.Sprintf("%d", value)
	case float32, float64:
		return fmt.Sprintf("%f", value)
		//case interface{}:
		//	return fmt.Sprintf("%v", value)
	}

	return defaultValue
}
