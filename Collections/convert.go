package Collections

import (
	"encoding/json"
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Logs"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Utils/Convert"
	"strconv"
	"strings"
)

func (this *Collection) ToIntArray() (results []int) {
	for _, data := range this.array {
		results = append(results, Convert.ConvertToInt(data, 0))
	}
	return
}

func (this *Collection) ToInt64Array() (results []int64) {
	for _, data := range this.array {
		results = append(results, Convert.ConvertToInt64(data, 0))
	}
	return
}
func (this *Collection) ToInterfaceArray() []interface{} {
	return this.array
}

func (this *Collection) ToFloatArray() (results []float32) {
	for _, data := range this.array {
		results = append(results, Convert.ConvertToFloat(data, 0))
	}
	return
}

func (this *Collection) ToFloat64Array() (results []float64) {
	for _, data := range this.array {
		results = append(results, Convert.ConvertToFloat64(data, 0))
	}
	return
}

func (this *Collection) ToBoolArray() (results []bool) {
	for _, data := range this.array {
		results = append(results, Convert.ConvertToBool(data, false))
	}
	return
}

func (this *Collection) ToStringArray() (results []string) {
	for _, data := range this.array {
		results = append(results, Convert.ConvertToString(data, ""))
	}
	return
}

func (this *Collection) ToFields() Support.Fields {
	fields := Support.Fields{}
	for index, data := range this.mapData {
		fields[strconv.Itoa(index)] = data
	}
	return fields
}

func (this *Collection) ToArrayFields() []Support.Fields {
	return this.mapData
}

func (this *Collection) ToJson() string {
	results := make([]string, 0)
	this.Map(func(data interface{}) {
		if jsonify, isJson := data.(Support.Json); isJson {
			results = append(results, jsonify.ToJson())
			return
		}
		jsonStr, err := json.Marshal(data)
		if err != nil {
			Logs.WithError(err).WithFields(this.ToFields()).Fatal("json err")
		}
		results = append(results, string(jsonStr))
	})

	return fmt.Sprintf("[%s]", strings.Join(results, ","))
}

func (this *Collection) String() string {
	return this.ToJson()
}
