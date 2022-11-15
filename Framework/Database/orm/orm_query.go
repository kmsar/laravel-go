package gorose

import (
	"github.com/gohouse/t"
	"math"
	"reflect"
	"strings"
)

func (dba *Orm) Select() error {
	switch dba.GetIBinder().GetBindType() {
	case OBJECT_STRUCT, OBJECT_MAP, OBJECT_MAP_T:
		dba.Limit(1)
	}

	sqlStr, args, err := dba.BuildSql()
	if err != nil {
		return err
	}

	_, err = dba.GetISession().Query(sqlStr, args...)
	return err
}

func (dba *Orm) First() (result Data, err error) {
	dba.GetIBinder().SetBindType(OBJECT_STRING)
	err = dba.Limit(1).Select()
	if err != nil {
		return
	}
	res := dba.GetISession().GetBindAll()
	if len(res) > 0 {
		result = res[0]
	}
	return
}

func (dba *Orm) Get() (result []Data, err error) {
	dba.GetIBinder().SetBindType(OBJECT_STRING)
	tabname := dba.GetISession().GetIBinder().GetBindName()
	prefix := dba.GetISession().GetIBinder().GetBindPrefix()
	tabname2 := strings.TrimPrefix(tabname, prefix)
	dba.ResetTable()
	dba.Table(tabname2)
	err = dba.Select()
	result = dba.GetISession().GetBindAll()
	return
}

func (dba *Orm) Count(args ...string) (int64, error) {
	fields := "*"
	if len(args) > 0 {
		fields = args[0]
	}
	count, err := dba._unionBuild("count", fields)
	if count == nil {
		return 0, err
	}
	return t.New(count).Int64(), err
}

func (dba *Orm) Sum(sum string) (interface{}, error) {
	return dba._unionBuild("sum", sum)
}

func (dba *Orm) Avg(avg string) (interface{}, error) {
	return dba._unionBuild("avg", avg)
}

func (dba *Orm) Max(max string) (interface{}, error) {
	return dba._unionBuild("max", max)
}

func (dba *Orm) Min(min string) (interface{}, error) {
	return dba._unionBuild("min", min)
}

func (dba *Orm) _unionBuild(union, field string) (interface{}, error) {
	fields := union + "(" + field + ") as " + union
	dba.fields = []string{fields}

	res, err := dba.First()
	if r, ok := res[union]; ok {
		return r, err
	}
	return 0, err
}

func (dba *Orm) Pluck(field string, fieldKey ...string) (v interface{}, err error) {
	var resMap = make(map[interface{}]interface{}, 0)
	var resSlice = make([]interface{}, 0)

	res, err := dba.Get()

	if err != nil {
		return
	}

	if len(res) > 0 {
		for _, val := range res {
			if len(fieldKey) > 0 {
				resMap[val[fieldKey[0]]] = val[field]
			} else {
				resSlice = append(resSlice, val[field])
			}
		}
	}
	if len(fieldKey) > 0 {
		v = resMap
	} else {
		v = resSlice
	}
	return
}

func (dba *Orm) Pluck_bak(field string, fieldKey ...string) (v interface{}, err error) {
	var binder = dba.GetISession().GetIBinder()
	var resMap = make(map[interface{}]interface{}, 0)
	var resSlice = make([]interface{}, 0)

	err = dba.Select()
	if err != nil {
		return
	}

	switch binder.GetBindType() {
	case OBJECT_MAP, OBJECT_MAP_T, OBJECT_STRUCT:
		var key, val interface{}
		if len(fieldKey) > 0 {
			key, err = dba.Value(fieldKey[0])
			if err != nil {
				return
			}
			val, err = dba.Value(field)
			if err != nil {
				return
			}
			resMap[key] = val
		} else {
			v, err = dba.Value(field)
			if err != nil {
				return
			}
		}
	case OBJECT_MAP_SLICE, OBJECT_MAP_SLICE_T:
		for _, item := range t.New(binder.GetBindResultSlice().Interface()).Slice() {
			val := item.MapInterfaceT()
			if len(fieldKey) > 0 {
				resMap[val[fieldKey[0]].Interface()] = val[field].Interface()
			} else {
				resSlice = append(resSlice, val[field].Interface())
			}
		}
	case OBJECT_STRUCT_SLICE:
		var brs = binder.GetBindResultSlice()
		for i := 0; i < brs.Len(); i++ {
			val := reflect.Indirect(brs.Index(i))
			if len(fieldKey) > 0 {
				mapkey := dba._valueFromStruct(val, fieldKey[0])
				mapVal := dba._valueFromStruct(val, field)
				resMap[mapkey] = mapVal
			} else {
				resSlice = append(resSlice, dba._valueFromStruct(val, field))
			}
		}
	case OBJECT_STRING:
		res := dba.GetISession().GetBindAll()
		if len(res) > 0 {
			for _, val := range res {
				if len(fieldKey) > 0 {
					resMap[val[fieldKey[0]]] = val[field]
				} else {
					resSlice = append(resSlice, val[field])
				}
			}
		}
	}
	if len(fieldKey) > 0 {
		v = resMap
	} else {
		v = resSlice
	}
	return
}

func (dba *Orm) Value(field string) (v interface{}, err error) {
	res, err := dba.First()
	if v, ok := res[field]; ok {
		return v, err
	}
	return
}

func (dba *Orm) Value_bak(field string) (v interface{}, err error) {
	dba.Limit(1)
	err = dba.Select()
	if err != nil {
		return
	}
	var binder = dba.GetISession().GetIBinder()
	switch binder.GetBindType() {
	case OBJECT_MAP, OBJECT_MAP_SLICE, OBJECT_MAP_SLICE_T, OBJECT_MAP_T:
		v = reflect.ValueOf(binder.GetBindResult()).MapIndex(reflect.ValueOf(field)).Interface()
	case OBJECT_STRUCT, OBJECT_STRUCT_SLICE:
		bindResult := reflect.Indirect(reflect.ValueOf(binder.GetBindResult()))
		v = dba._valueFromStruct(bindResult, field)
	case OBJECT_STRING:
		res := dba.GetISession().GetBindAll()
		if len(res) > 0 {
			v = res[0][field]
		}
	}
	return
}
func (dba *Orm) _valueFromStruct(bindResult reflect.Value, field string) (v interface{}) {
	ostype := bindResult.Type()
	for i := 0; i < ostype.NumField(); i++ {
		tag := ostype.Field(i).Tag.Get(TAGNAME)
		if tag == field || ostype.Field(i).Name == field {
			v = bindResult.FieldByName(ostype.Field(i).Name).Interface()
		}
	}
	return
}

func (dba *Orm) Chunk(limit int, callback func([]Data) error) (err error) {
	var page = 1
	var tabname = dba.GetISession().GetIBinder().GetBindName()
	prefix := dba.GetISession().GetIBinder().GetBindPrefix()
	tabname2 := strings.TrimPrefix(tabname, prefix)

	result, err := dba.Table(tabname2).Limit(limit).Page(page).Get()
	if err != nil {
		return
	}
	for len(result) > 0 {
		if err = callback(result); err != nil {
			break
		}
		page++

		dba.ClearBindValues()
		result, _ = dba.Page(page).Get()
	}
	return
}

func (dba *Orm) ChunkStruct(limit int, callback func() error) (err error) {
	var page = 0

	err = dba.Limit(limit).Offset(page * limit).Select()
	if err != nil {
		return
	}
	switch dba.GetIBinder().GetBindType() {
	case OBJECT_STRUCT, OBJECT_MAP, OBJECT_MAP_T:
		var ibinder = dba.GetIBinder()
		var result = ibinder.GetBindResult()
		for result != nil {
			if err = callback(); err != nil {
				break
			}
			page++

			var rfRes = reflect.ValueOf(result)
			rfRes.Set(reflect.Zero(rfRes.Type()))

			dba.ClearBindValues()
			_ = dba.Table(ibinder.GetBindOrigin()).Offset(page * limit).Select()
			result = dba.GetIBinder().GetBindResultSlice()
		}
	case OBJECT_STRUCT_SLICE, OBJECT_MAP_SLICE, OBJECT_MAP_SLICE_T:
		var ibinder = dba.GetIBinder()
		var result = ibinder.GetBindResultSlice()
		for result.Interface() != nil {
			if err = callback(); err != nil {
				break
			}
			page++

			result.Set(result.Slice(0, 0))

			dba.ClearBindValues()
			_ = dba.Table(ibinder.GetBindOrigin()).Offset(page * limit).Select()
			result = dba.GetIBinder().GetBindResultSlice()
		}
	}
	return
}

func (dba *Orm) Loop(limit int, callback func([]Data) error) (err error) {
	var page = 0
	var tabname = dba.GetISession().GetIBinder().GetBindName()
	prefix := dba.GetISession().GetIBinder().GetBindPrefix()
	tabname2 := strings.TrimPrefix(tabname, prefix)

	result, err := dba.Table(tabname2).Limit(limit).Get()
	if err != nil {
		return
	}
	for len(result) > 0 {
		if err = callback(result); err != nil {
			break
		}
		page++

		dba.ClearBindValues()
		result, _ = dba.Get()
	}
	return
}

func (dba *Orm) Paginate(page ...int) (res Data, err error) {
	if len(page) > 0 {
		dba.Page(page[0])
	}
	var limit = dba.GetLimit()
	if limit == 0 {
		limit = 15
	}
	var offset = dba.GetOffset()
	var currentPage = int(math.Ceil(float64(offset+1) / float64(limit)))

	resData, err := dba.Get()
	if err != nil {
		return
	}

	dba.offset = 0
	count, err := dba.Count()
	var lastPage = int(math.Ceil(float64(count) / float64(limit)))
	var nextPage = currentPage + 1
	var prevPage = currentPage - 1
	res = Data{
		"total":          count,
		"per_page":       limit,
		"current_page":   currentPage,
		"last_page":      lastPage,
		"first_page_url": 1,
		"last_page_url":  lastPage,
		"next_page_url":  If(nextPage > lastPage, nil, nextPage),
		"prev_page_url":  If(prevPage < 1, nil, prevPage),

		"data": resData,
	}

	return
}
