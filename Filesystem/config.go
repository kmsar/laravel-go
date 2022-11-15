package Filesystem

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

type Config struct {
	Default string

	Disks map[string]Support.Fields
}
