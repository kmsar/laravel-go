package Session

import "time"

type Config struct {
	Driver string

	Encrypt bool

	Domain string

	Lifetime time.Duration

	Connection string

	Table string

	Name string

	Key string
}
