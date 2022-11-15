package Slices

import (
	"reflect"
	"sort"
)

// ArrayKeys returns all the keys or a subset of the keys of an array
func ArrayKeys(input any) any {
	if input == nil {
		return nil
	}
	val := reflect.ValueOf(input)
	if val.Len() == 0 {
		return nil
	}
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		var res []int
		for i := 0; i < val.Len(); i++ {
			res = append(res, i)
		}
		return res
	case reflect.Map:
		var res []string
		for _, k := range val.MapKeys() {
			res = append(res, k.String())
		}
		sort.SliceStable(res, func(i, j int) bool {
			return res[i] < res[j]
		})
		return res
	}
	return nil
}
