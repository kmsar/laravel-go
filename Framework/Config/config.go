package Config

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IConfig"
	"github.com/kmsar/laravel-go/Framework/Contracts/IEnv"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Field"
	"github.com/kmsar/laravel-go/Framework/Support/Utils/Convert"
	"strings"
	"sync"
)

func NewConfig(env IEnv.Env, providers map[string]IConfig.ConfigProvider) IConfig.Config {
	return &config{
		writeMutex: sync.RWMutex{},
		providers:  providers,
		Env:        env,
		fields:     make(Support.Fields),
		configs:    make(map[string]IConfig.Config, 0),
	}
}

func WithFields(fields Support.Fields) IConfig.Config {
	return &config{
		fields:  fields,
		configs: make(map[string]IConfig.Config, 0),
	}
}

type config struct {
	writeMutex sync.RWMutex
	fields     Support.Fields
	configs    map[string]IConfig.Config
	providers  map[string]IConfig.ConfigProvider
	IEnv.Env
}

func (this *config) Fields() Support.Fields {
	return this.fields
}

func (this *config) Load(provider Support.FieldsProvider) {
	Field.MergeFields(this.fields, provider.Fields())
}

func (this *config) Reload() {
	for name, provider := range this.providers {
		this.Set(name, provider(this.Env))
	}
}

func (this *config) Merge(key string, config IConfig.Config) {
	this.fields[key] = config.Fields()
	this.configs[key] = config
}

func (this *config) Set(key string, value interface{}) {
	this.writeMutex.Lock()
	this.fields[key] = value
	this.writeMutex.Unlock()
}

func (this *config) Get(key string, defaultValue ...interface{}) interface{} {
	this.writeMutex.RLock()
	defer this.writeMutex.RUnlock()

	// 环境变量优先级最高
	if this.Env != nil {
		if envValue := this.Env.GetString(key); envValue != "" {
			return envValue
		}
	}

	if field, existsField := this.fields[key]; existsField {
		return field
	}

	// 尝试获取 fields
	var (
		fields = Support.Fields{}
		prefix = key + "."
	)

	for fieldKey, fieldValue := range this.fields {
		if strings.HasPrefix(fieldKey, prefix) {
			fields[strings.Replace(fieldKey, prefix, "", 1)] = fieldValue
		}
	}

	if len(fields) > 0 {
		return fields
	}

	var keys = strings.Split(key, ".")

	if len(keys) > 1 {
		if subConfig, existsSubConfig := this.configs[keys[0]]; existsSubConfig {
			return subConfig.Get(strings.Join(keys[1:], "."), defaultValue...)
		}
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return nil
}

func (this *config) GetConfig(key string) IConfig.Config {
	return this.configs[key]
}

func (this *config) GetFields(key string) Support.Fields {
	if field, isTypeRight := this.Get(key).(Support.Fields); isTypeRight {
		return field
	}

	return nil
}

func (this *config) GetString(key string) string {
	if field, isTypeRight := this.Get(key).(string); isTypeRight {
		return field
	}

	return ""
}

func (this *config) GetInt(key string) int {
	if field := this.Get(key); field != nil {
		value := Convert.ConvertToInt(field, 0)
		if value != 0 { // 缓存转换结果
			this.Set(key, value)
		}
		return value
	}

	return 0
}
func (this *config) GetInt64(key string) int64 {
	if field := this.Get(key); field != nil {
		value := Convert.ConvertToInt64(field, 0)
		if value != 0 { // 缓存转换结果
			this.Set(key, value)
		}
		return value
	}

	return 0
}

func (this *config) Unset(key string) {
	delete(this.fields, key)
	delete(this.configs, key)
}

func (this *config) GetFloat(key string) float32 {
	if field := this.Get(key); field != nil {
		value := Convert.ConvertToFloat(field, 0)
		if value != 0 { // 缓存转换结果
			this.Set(key, value)
		}
		return value
	}

	return 0
}
func (this *config) GetFloat64(key string) float64 {
	if field := this.Get(key); field != nil {
		value := Convert.ConvertToFloat64(field, 0)
		if value != 0 { // 缓存转换结果
			this.Set(key, value)
		}
		return value
	}

	return 0
}

func (this *config) GetBool(key string) bool {
	if field := this.Get(key); field != nil {
		result := Convert.ConvertToBool(field, false)
		this.Set(key, result)
		return result
	}

	return false
}
