package Console

import (
	"errors"
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Console/scheduling"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Exceptions"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Str/table"

	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConsole"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
)

var CommandDontExists = errors.New("命令不存在！")

const logoText = "  ▄████  ▒█████   ▄▄▄       ██▓    \n ██▒ ▀█▒▒██▒  ██▒▒████▄    ▓██▒    \n▒██░▄▄▄░▒██░  ██▒▒██  ▀█▄  ▒██░    \n░▓█  ██▓▒██   ██░░██▄▄▄▄██ ▒██░    \n░▒▓███▀▒░ ████▓▒░ ▓█   ▓██▒░██████▒\n ░▒   ▒ ░ ▒░▒░▒░  ▒▒   ▓▒█░░ ▒░▓  ░\n  ░   ░   ░ ▒ ▒░   ▒   ▒▒ ░░ ░ ▒  ░\n░ ░   ░ ░ ░ ░ ▒    ░   ▒     ░ ░   \n      ░     ░ ░        ░  ░    ░  ░\n                                   "

type Kernel struct {
	app              IFoundation.IApplication
	commands         map[string]IConsole.CommandProvider
	schedule         IConsole.Schedule
	exceptionHandler IExeption.ExceptionHandler
}

func (this *Kernel) RegisterCommand(name string, command IConsole.CommandProvider) {
	this.commands[name] = command
}

func (this *Kernel) GetSchedule() IConsole.Schedule {
	return this.schedule
}

func (this *Kernel) Schedule(schedule IConsole.Schedule) {
}

func NewKernel(app IFoundation.IApplication, commandProviders []IConsole.CommandProvider) *Kernel {
	var commands = make(map[string]IConsole.CommandProvider)
	for _, provider := range commandProviders {
		commands[provider(app).GetName()] = provider
	}
	return &Kernel{
		app:              app,
		commands:         commands,
		schedule:         scheduling.NewSchedule(app),
		exceptionHandler: app.Get("exceptions.handler").(IExeption.ExceptionHandler),
	}
}

type CommandItem struct {
	Command     string
	Signature   string
	Description string
}

func (this Kernel) Help() {
	cmdTable := make([]CommandItem, 0)
	for _, command := range this.commands {
		cmd := command(this.app)
		cmdTable = append(cmdTable, CommandItem{
			Command:     cmd.GetName(),
			Signature:   cmd.GetSignature(),
			Description: cmd.GetDescription(),
		})
	}
	fmt.Println(logoText)
	table.Output(cmdTable)
}

func (this *Kernel) Call(cmd string, arguments IConsole.CommandArguments) interface{} {
	if cmd == "" {
		this.Help()
		return nil
	}
	for name, provider := range this.commands {
		if cmd == name {
			command := provider(this.app)
			if arguments.Exists("h") || arguments.Exists("help") {
				fmt.Println(logoText)
				fmt.Printf(" %s 命令：%s\n", command.GetName(), command.GetDescription())
				fmt.Println(command.GetHelp())
				return nil
			}
			if err := command.InjectArguments(arguments); err != nil {
				this.exceptionHandler.Handle(CommandArgumentException{
					Exceptions.WithError(err, Support.Fields{
						"command":   cmd,
						"arguments": arguments,
					}),
				})
				fmt.Println(err.Error())
				fmt.Println(command.GetHelp())
				return nil
			}
			return command.Handle()
		}
	}
	return CommandDontExists
}

func (this *Kernel) Run(input IConsole.ConsoleInput) interface{} {
	return this.Call(input.GetCommand(), input.GetArguments())
}

func (this *Kernel) Exists(name string) bool {
	return this.commands[name] != nil
}
