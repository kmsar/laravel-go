package Console

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
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
