package inputs

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IConsole"
	"os"
)

type ArgsInput struct {
	StringArrayInput
}

func NewOSArgsInput() IConsole.ConsoleInput {
	return &ArgsInput{StringArray(os.Args[1:])}
}
