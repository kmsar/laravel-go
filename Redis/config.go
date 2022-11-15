package Redis

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

type Config struct {
	Default string
	Stores  map[string]Support.Fields
}
