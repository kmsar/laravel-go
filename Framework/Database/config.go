package Database

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

type Config struct {
	Default     string
	Connections map[string]Support.Fields
	Migrations  string
}
