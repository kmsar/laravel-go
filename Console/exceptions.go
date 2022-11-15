package Console

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
)

type CommandArgumentException struct {
	IExeption.Exception
}

type CommandDontExistsException struct {
	IExeption.Exception
}

type ScheduleEventException struct {
	IExeption.Exception
}
