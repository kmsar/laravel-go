package Database

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

type Config struct {
	Default     string
	Connections map[string]Support.Fields
	Migrations  string
}
