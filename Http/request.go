package Http

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Http/echo"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Logs"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Field"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Validation"
	"strings"
)

type Request struct {
	Field.BaseFields
	echo.Context
	fields Support.Fields
}

func NewRequest(ctx echo.Context) IHttp.IHttpRequest {
	var request = &Request{
		Context:    ctx,
		fields:     nil,
		BaseFields: Field.BaseFields{},
	}

	request.BaseFields.FieldsProvider = request
	request.BaseFields.Getter = request.get

	return request
}

func (this *Request) Get(key string) (value interface{}) {
	if value = this.Context.Get(key); value != nil && value != "" {
		return value
	}
	if value = this.Context.QueryParam(key); value != nil && value != "" {
		return value
	}
	if value = this.Context.FormValue(key); value != nil && value != "" {
		return value
	}
	if value = this.Context.Param(key); value != nil && value != "" {
		return value
	}
	if file, err := this.Context.FormFile(key); err == nil && file != nil {
		return file
	}
	return
}

func (this *Request) get(key string) (value interface{}) {
	if value = this.Get(key); value != nil && value != "" {
		return value
	}
	return this.Fields()[key]
}

func (this *Request) Validate(v interface{}) error {
	if err := this.Bind(v); err != nil {
		return err
	}

	return Validation.Struct(v)
}

func (this *Request) Fields() Support.Fields {
	if this.fields != nil {
		return this.fields
	}
	var data = make(Support.Fields)
	if strings.Contains(this.Request().Header.Get("Content-Type"), "json") {
		var bindErr = this.Context.Bind(&data)
		if bindErr != nil {
			Logs.WithError(bindErr).Debug("http.Request.Fields: bind fields failed")
		}
	}

	for key, query := range this.QueryParams() {
		if len(query) == 1 {
			data[key] = query[0]
		} else {
			data[key] = query
		}
	}
	for _, paramName := range this.ParamNames() {
		data[paramName] = this.Param(paramName)
	}
	if form, existsForm := this.FormParams(); existsForm == nil {
		for key, values := range form {
			if len(values) == 1 {
				data[key] = values[0]
			} else {
				data[key] = values
			}
		}
	}
	if multiForm, existsForm := this.MultipartForm(); existsForm == nil {
		for key, values := range multiForm.Value {
			if len(values) == 1 {
				data[key] = values[0]
			} else {
				data[key] = values
			}
		}
		for key, values := range multiForm.File {
			if len(values) == 1 {
				data[key] = values[0]
			} else {
				data[key] = values
			}
		}
	}

	this.fields = data

	return data
}
