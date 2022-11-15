package Slices

import "reflect"

// ArrayFlip exchanges all keys with their associated values in an array
func ArrayFlip(input any) any {
	if input == nil {
		return nil
	}
	val := reflect.ValueOf(input)
	if val.Len() == 0 {
		return nil
	}
	res := make(map[any]any, val.Len())
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			res[val.Index(i).Interface()] = i
		}
		return res
	case reflect.Map:
		for _, k := range val.MapKeys() {
			res[val.MapIndex(k).Interface()] = k.Interface()
		}
		return res
	}
	return nil
}
