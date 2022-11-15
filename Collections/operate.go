package Collections

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Field"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Random"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Utils"
)

func (this *Collection) Pluck(key string) Support.Fields {
	fields := Support.Fields{}

	for index, data := range this.mapData {
		var name, ok = data[key].(string)
		if _, exists := fields[name]; ok && !exists {
			fields[name] = this.array[index]
		}
	}

	return fields
}

func (this *Collection) Only(keys ...string) Support.Collection {
	arrayFields := make([]Support.Fields, 0)
	rawResults := make([]interface{}, 0)

	for index, data := range this.mapData {
		fields := Support.Fields{}
		for key, value := range data {
			if Utils.IsIn(key, keys) {
				fields[key] = value
			}
		}
		arrayFields = append(arrayFields, fields)
		rawResults = append(rawResults, this.array[index])
	}

	return &Collection{mapData: arrayFields, array: rawResults}
}

func (this *Collection) First(keys ...string) interface{} {
	if this.Count() == 0 {
		return nil
	}
	if len(keys) == 0 {
		return this.array[0]
	}
	return this.mapData[0][keys[0]]
}

func (this *Collection) Last(keys ...string) interface{} {
	if this.Count() == 0 {
		return nil
	}
	if len(keys) == 0 {
		return this.array[len(this.array)-1]
	}
	return this.mapData[len(this.array)-1][keys[0]]
}

func (this *Collection) Prepend(items ...interface{}) Support.Collection {
	newCollection := &Collection{}
	newCollection.array = append(items, this.array...)
	if len(this.mapData) > 0 {
		newMaps := make([]Support.Fields, 0)
		for _, item := range items {
			fields, _ := Field.ConvertToFields(item)
			newMaps = append(newMaps, fields)
		}
		newCollection.mapData = append(newMaps, this.mapData...)
	}
	return newCollection
}

func (this *Collection) Push(items ...interface{}) Support.Collection {
	newCollection := &Collection{}
	newCollection.array = append(this.array, items...)
	if len(this.mapData) > 0 {
		newMaps := make([]Support.Fields, 0)
		for _, item := range items {
			fields, _ := Field.ConvertToFields(item)
			newMaps = append(newMaps, fields)
		}
		newCollection.mapData = append(this.mapData, newMaps...)
	}
	return newCollection
}

func (this *Collection) Pull(defaultValue ...interface{}) interface{} {
	if result := this.Last(); result != nil {
		this.array = this.array[:this.Count()-1]
		if len(this.mapData) > 0 {
			this.mapData = this.mapData[:this.Count()-1]
		}
		return result
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return nil
}

func (this *Collection) Shift(defaultValue ...interface{}) interface{} {
	if result := this.First(); result != nil {
		this.array = this.array[1:]
		if len(this.mapData) > 0 {
			this.mapData = this.mapData[1:]
		}
		return result
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return nil
}

func (this *Collection) Offset(index int, item interface{}) Support.Collection {
	if this.Count() > index {
		this.array[index] = item
		if len(this.mapData) > 0 {
			fields, _ := Field.ConvertToFields(item)
			this.mapData[index] = fields
		}
		return this
	}
	return this.Push(item)
}

func (this *Collection) Put(index int, item interface{}) Support.Collection {
	if this.Count() > index {
		return (&Collection{array: append(this.array), mapData: append(this.mapData)}).Offset(index, item)
	}
	return this.Push(item)
}

func (this *Collection) Merge(collections ...Support.Collection) Support.Collection {
	newCollection := &Collection{array: append(this.array), mapData: append(this.mapData)}

	for _, collection := range collections {
		newCollection.mapData = append(newCollection.mapData, collection.ToArrayFields()...)
		newCollection.array = append(newCollection.array, collection.ToInterfaceArray()...)
	}

	return newCollection
}

func (this *Collection) Reverse() Support.Collection {
	newCollection := &Collection{array: append(this.array), mapData: append(this.mapData)}
	for from, to := 0, len(newCollection.array)-1; from < to; from, to = from+1, to-1 {
		newCollection.array[from], newCollection.array[to] = newCollection.array[to], newCollection.array[from]
		if len(this.mapData) > 0 {
			newCollection.mapData[from], newCollection.mapData[to] = newCollection.mapData[to], newCollection.mapData[from]
		}
	}
	return newCollection
}

func (this *Collection) Chunk(size int, handler func(collection Support.Collection, page int) error) (err error) {
	total := this.Count()
	page := 1
	for err == nil && (page-1)*size <= total {
		offset := (page - 1) * size
		endIndex := size + offset
		if endIndex > total {
			endIndex = total
		}
		newCollection := &Collection{array: this.array[offset:endIndex]}
		if len(this.mapData) > 0 {
			newCollection.mapData = this.mapData[offset:endIndex]
		}

		err = handler(newCollection, page)
		page++
	}

	return
}

func (this *Collection) Random(size ...uint) Support.Collection {
	num := 1
	if len(size) > 0 {
		num = int(size[0])
	}
	newCollection := &Collection{}
	if this.Count() >= num {
		for _, index := range Random.RandIntArray(0, this.Count()-1, num) {
			newCollection.array = append(newCollection.array, this.array[index])
			if len(this.mapData) > 0 {
				newCollection.mapData = append(newCollection.mapData, this.mapData[index])
			}
		}
	}
	return newCollection
}
