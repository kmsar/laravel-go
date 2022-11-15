package inputs

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConsole"
	"os"
)

type ArgsInput struct {
	StringArrayInput
}

func NewOSArgsInput() IConsole.ConsoleInput {
	return &ArgsInput{StringArray(os.Args[1:])}
}
