package Filesystem

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

type Config struct {
	Default string

	Disks map[string]Support.Fields
}
