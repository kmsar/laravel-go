package Auth

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
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
