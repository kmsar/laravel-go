package Utils

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"time"
)

func DefaultString(values []string, defaultValue ...string) string {
	if len(values) > 0 {
		return values[0]
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return ""
}

func DefaultInt(values []int, defaultValue ...int) int {
	if len(values) > 0 {
		return values[0]
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return 0
}

func DefaultDuration(values []time.Duration, defaultValue ...time.Duration) time.Duration {
	if len(values) > 0 {
		return values[0]
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return 0
}

func DefaultTime(values []time.Time, defaultValue ...time.Time) time.Time {
	if len(values) > 0 {
		return values[0]
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return time.Time{}
}

func DefaultInt64(values []int64, defaultValue ...int64) int64 {
	if len(values) > 0 {
		return values[0]
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return 0
}

func DefaultUint64(values []uint64, defaultValue ...uint64) uint64 {
	if len(values) > 0 {
		return values[0]
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return 0
}

func DefaultUint(values []uint, defaultValue ...uint) uint {
	if len(values) > 0 {
		return values[0]
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return 0
}

func DefaultException(values []IExeption.Exception, defaultValue ...IExeption.Exception) IExeption.Exception {
	if len(values) > 0 {
		return values[0]
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return nil
}

func DefaultFloat(values []float32, defaultValue ...float32) float32 {
	if len(values) > 0 {
		return values[0]
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return 0
}

func DefaultFloat64(values []float64, defaultValue ...float64) float64 {
	if len(values) > 0 {
		return values[0]
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return 0
}

func DefaultBool(values []bool, defaultValue ...bool) bool {
	if len(values) > 0 {
		return values[0]
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return false
}

func DefaultInterface(values []interface{}, defaultValue ...interface{}) interface{} {
	if len(values) > 0 {
		return values[0]
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return nil
}

func DefaultError(values []error, defaultValue ...error) error {
	if len(values) > 0 {
		return values[0]
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return nil
}

func DefaultFields(values []Support.Fields, defaultValue ...Support.Fields) Support.Fields {
	if len(values) > 0 {
		return values[0]
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return nil
}
