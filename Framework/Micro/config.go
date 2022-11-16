package Micro

import "go-micro.dev/v4"

type Config struct {
	micro.Options
	CustomOptions []micro.Option
}
