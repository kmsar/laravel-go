package IEnv

import "github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"

type Env interface {
	//Get(key string, def ...string) string
	//Set(key, value string) error
	//List() map[string]string
	//Write(toEnvPath string) error
	//GetString(key string) string

	Support.Getter
	Support.OptionalGetter
	Support.FieldsProvider
	// Load load configuration.
	Load() Support.Fields
}
