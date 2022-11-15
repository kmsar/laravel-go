package Slices

import "reflect"

// ArrayFilter filters elements of an array using a callback function
func ArrayFilter(input any, callback func(any) bool) any {
	if input == nil {
		return nil
	}
	val := reflect.ValueOf(input)
	if val.Len() == 0 {
		return nil
	}
	if callback == nil {
		callback = func(v any) bool {
			return v != nil
		}
	}
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		var res []any
		for i := 0; i < val.Len(); i++ {
			v := val.Index(i).Interface()
			if callback(v) {
				res = append(res, v)
			}
		}
		return res
	case reflect.Map:
		res := make(map[any]any)
		for _, k := range val.MapKeys() {
			v := val.MapIndex(k).Interface()
			if callback(v) {
				res[k.Interface()] = v
			}
		}
		return res
	}

	return input
}
