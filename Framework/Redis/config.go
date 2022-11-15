package Redis

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

type Config struct {
	Default string
	Stores  map[string]Support.Fields
}
