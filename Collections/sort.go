package Collections

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Exceptions"
	"reflect"
	"sort"
)

func (this *Collection) Len() int {
	return len(this.array)
}

func (this *Collection) Swap(i, j int) {
	this.array[i], this.array[j] = this.array[j], this.array[i]
	if len(this.mapData) > 0 {
		this.mapData[i], this.mapData[j] = this.mapData[j], this.mapData[i]
	}
}

func (this *Collection) Less(i, j int) bool {
	if this.sorter != nil {
		return this.sorter(i, j)
	}
	return i > j
}

func (this *Collection) SetSorter(sorter func(i, j int) bool) Support.Collection {
	this.sorter = sorter
	return this
}

func (this *Collection) Sort(sorter interface{}) Support.Collection {
	sorterType := reflect.TypeOf(sorter)

	if sorterType.Kind() != reflect.Func || sorterType.NumIn() != 2 || sorterType.NumOut() != 1 || sorterType.Out(0).Kind() != reflect.Bool {
		Exceptions.Throw(SortException{
			Exceptions.New(
				"：sorter is invalid",
				Support.Fields{
					"sorter": sorter,
				})})
	}
	sorterValue := reflect.ValueOf(sorter)

	newCollection := (&Collection{
		mapData: this.mapData,
		array:   this.array,
	}).SetSorter(func(i, j int) bool {
		return sorterValue.Call([]reflect.Value{
			this.argumentConvertor(sorterType.In(0), this.array[i]),
			this.argumentConvertor(sorterType.In(1), this.array[j]),
		})[0].Bool()
	})

	sort.Stable(newCollection)

	return newCollection
}
