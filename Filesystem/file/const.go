package file

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFilesystem"
)

const (
	VISIBLE IFilesystem.FileVisibility = iota
	INVISIBLE
)
