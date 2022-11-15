package Console

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Console/myco"
)

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Boot() {
	receiver.registerCommands()
}

func (receiver *ServiceProvider) Register() {
	app := myco.NewCli()
	app.Register([]myco.ICommand{
		&ListCommand{},
		&KeyGenerateCommand{},
		&MakeCommand{},
	})
	//facades.Artisan = app.Init()
}

func (receiver *ServiceProvider) registerCommands() {
	//facades.Artisan.Register([]console2.Command{
	//	&console.ListCommand{},
	//	&console.KeyGenerateCommand{},
	//	&console.MakeCommand{},
	//})
}
