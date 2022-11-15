package Env

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEnv"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Field"
	"os"
	"path/filepath"
	"strings"
)

type envProvider struct {
	Field.BaseFields
	Paths  []string
	Sep    string
	fields Support.Fields
}

func NewEnv(paths []string, sep string) IEnv.Env {
	provider := &envProvider{
		BaseFields: Field.BaseFields{Getter: func(key string) interface{} {
			return os.Getenv(key)
		}},
		Paths:  paths,
		Sep:    sep,
		fields: nil,
	}

	provider.BaseFields.FieldsProvider = provider
	return provider
}

func (this *envProvider) Fields() Support.Fields {
	if this.fields != nil {
		return this.fields
	}

	this.fields = this.Load()

	return this.fields
}

func (this *envProvider) Load() Support.Fields {
	var (
		files  []string
		fields = make(Support.Fields)
	)
	for _, path := range this.Paths {
		tmpFiles, _ := filepath.Glob(path + "/*.env")
		files = append(files, tmpFiles...)
		err := Load(path)
		if err != nil {
			panic(err)
		}
	}

	for _, e := range os.Environ() {
		newFields := make(Support.Fields)
		pair := strings.SplitN(e, "=", 2)
		newFields[pair[0]] = pair[1]
		Field.MergeFields(fields, newFields)
	}

	//for _, file := range files {
	//	tempFields, _ := utils.LoadEnv(file, utils.StringOr(this.Sep, "="))
	//	if tempFields["env"] != nil { // Loaded successfully and set up env
	//		newFields := make(Support.Fields)
	//		envValue := tempFields["env"].(string)
	//		for key, field := range tempFields {
	//			if key != "env" {
	//				newFields[fmt.Sprintf("%s:%s", envValue, key)] = field
	//			}
	//		}
	//		tempFields = newFields
	//	}
	//	utils.MergeFields(fields, tempFields)
	//}

	return fields
}

//
//
//
//type Env struct {
//	fields Support.Fields
//}
//
//func (e *Env) GetInt64(key string) int64 {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *Env) GetInt(key string) int {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *Env) GetFloat64(key string) float64 {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *Env) GetFloat(key string) float32 {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *Env) GetBool(key string) bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *Env) GetFields(key string) Support.Fields {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *Env) Int64Option(key string, defaultValue int64) int64 {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *Env) IntOption(key string, defaultValue int) int {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *Env) Float64Option(key string, defaultValue float64) float64 {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *Env) FloatOption(key string, defaultValue float32) float32 {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *Env) BoolOption(key string, defaultValue bool) bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *Env) FieldsOption(key string, defaultValue Support.Fields) Support.Fields {
//	val, ok := os.LookupEnv(key)
//	if ok {
//		return Support.Fields[]
//	}
//	return defaultValue[0]
//}
//
//func (e *Env) Fields() Support.Fields {
//	return e.fields
//}
//
//func (e *Env) Load() Support.Fields {
//	e.fields = make(Support.Fields)
//
//	for _, en := range os.Environ() {
//		pair := strings.SplitN(en, "=", 2)
//		e.fields[pair[0]] = pair[1]
//	}
//
//	return e.fields
//}
//
//func (e *Env) GetString(key string) string {
//	val, ok := os.LookupEnv(key)
//	if ok {
//		return val
//	}
//	return ""
//}
//
//// New New(".env-sample")
//func New(envFullPaths ...string) IEnv.Env {
//	res := &Env{}
//	for _, path := range envFullPaths {
//		err := Load(path)
//		if err != nil {
//			panic(err)
//		}
//	}
//	return res
//}
//
//func (e *Env) load(key string, def ...string) string {
//	val, ok := os.LookupEnv(key)
//	if ok {
//		return val
//	}
//	return def[0]
//}
//
//func (e *Env) Get(key string, def ...string) string {
//	val, ok := os.LookupEnv(key)
//	if ok {
//		return val
//	}
//	return def[0]
//}
//func (e *Env) StringOption(key string, def string) string {
//	val, ok := os.LookupEnv(key)
//	if ok {
//		return val
//	}
//	return def
//}
//
//func (e *Env) Set(key, value string) error {
//	return os.Setenv(key, value)
//}
//
//func (e *Env) List() map[string]string {
//	var res map[string]string
//	res = make(map[string]string)
//	for _, e := range os.Environ() {
//		pair := strings.SplitN(e, "=", 2)
//		res[pair[0]] = pair[1]
//	}
//	return res
//}
//
//func (e *Env) Write(toEnvPath string) error {
//	return Write(e.List(), toEnvPath)
//}
