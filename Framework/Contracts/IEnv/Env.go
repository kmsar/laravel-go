package IEnv

import "github.com/kmsar/laravel-go/Framework/Contracts/Support"

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
