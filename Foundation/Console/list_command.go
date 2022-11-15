package Console

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Console/myco"
)

type ListCommand struct {
}

// Signature The name and signature of the console command.
func (receiver *ListCommand) Signature() string {
	return "list"
}

// Description The console command description.
func (receiver *ListCommand) Description() string {
	return "List commands"
}

// Extend The console command extend.
func (receiver *ListCommand) Extend() myco.Extend {
	return myco.Extend{}
}

// Handle Execute the console command.
func (receiver *ListCommand) Handle(ctx myco.IContext) error {
	facades.Artisan.Call("--help")

	return nil
}
