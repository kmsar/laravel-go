package Auth

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

type Defaults struct {
	Guard string
	User  string
}

type Config struct {
	Defaults Defaults
	Guards   map[string]Support.Fields
	Users    map[string]Support.Fields
}
