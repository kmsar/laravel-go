package Config

import (
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IEnv"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Field"
	"github.com/kmsar/laravel-go/Framework/Support/Str"
	"io/ioutil"
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
	}

	for _, file := range files {
		tempFields, _ := LoadEnv(file, Str.StringOr(this.Sep, "="))
		if tempFields["env"] != nil { // 加载成功并且设置了 env
			newFields := make(Support.Fields)
			envValue := tempFields["env"].(string)
			for key, field := range tempFields {
				if key != "env" {
					newFields[fmt.Sprintf("%s:%s", envValue, key)] = field
				}
			}
			tempFields = newFields
		}
		Field.MergeFields(fields, tempFields)
	}

	return fields
}

func LoadEnv(envPath, sep string) (Support.Fields, error) {
	envBytes, err := ioutil.ReadFile(envPath)
	if err != nil {
		return nil, err
	}

	fields := make(Support.Fields)
	for _, line := range strings.Split(string(envBytes), "\n") {
		if strings.HasPrefix(line, "#") { // 跳过注释
			continue
		}
		values := strings.Split(line, sep)
		if len(values) > 1 {
			fields[values[0]] = strings.Trim(strings.ReplaceAll(strings.Join(values[1:], sep), `"`, ""), "\r\t\v\x00")
		}
	}

	return fields, nil
}
