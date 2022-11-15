package Mail

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

type Config struct {
	Default string
	Mailers map[string]Support.Fields
}
