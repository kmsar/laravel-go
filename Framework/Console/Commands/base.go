package Commands

import (
	"errors"
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IConsole"
)

type Command struct {
	IConsole.CommandArguments
	Signature   string
	Description string
	Name        string
	Help        string
	args        []Arg
}

//func (c *Command) Handle() interface{} {
//	//TODO implement me
//	panic("implement me")
//}

func Base(signature, description string) Command {
	name, args := ParseSignature(signature)
	return Command{
		Signature:   signature,
		Description: description,
		Name:        name,
		Help:        args.Help(),
		args:        args,
	}
}

func (c *Command) InjectArguments(arguments IConsole.CommandArguments) error {
	argIndex := 0
	for _, arg := range c.args {
		switch arg.Type {
		case RequiredArg:
			argValue := arguments.GetArg(argIndex)
			if argValue == "" {
				if c.Exists(arg.Name) {
					arguments.SetOption(arg.Name, arguments.Fields()[arg.Name])
				} else {
					return errors.New(fmt.Sprintf("Missing required parameter：%s - %s", arg.Name, arg.Description))
				}
			} else {
				arguments.SetOption(arg.Name, argValue)
			}
			argIndex++
		case OptionalArg:
			argValue := arguments.GetArg(argIndex)
			if argValue == "" {
				arguments.SetOption(arg.Name, arg.Default)
			} else {
				arguments.SetOption(arg.Name, argValue)
			}
			argIndex++
		case Option:
			if !arguments.Exists(arg.Name) && arg.Default != nil {
				arguments.SetOption(arg.Name, arg.Default)
			}
		}
	}

	c.CommandArguments = arguments
	return nil
}

func (c *Command) GetSignature() string {
	return c.Signature
}
func (c *Command) GetDescription() string {
	return c.Description
}
func (c *Command) GetName() string {
	return c.Name
}
func (c *Command) GetHelp() string {
	return c.Help
}
