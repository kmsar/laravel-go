package Mail

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

type Config struct {
	Default string
	Mailers map[string]Support.Fields
}
