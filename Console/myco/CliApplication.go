package myco

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Console/color"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

type CliApplication struct {
	instance *cli.App
}

func NewCli() Artisan {
	instance := cli.NewApp()
	instance.Name = Support.Name
	instance.Usage = Support.Version
	instance.UsageText = "artisan [global options] command [options] [arguments...]"

	return &CliApplication{instance}
}

func (c *CliApplication) Register(commands []ICommand) {
	for _, item := range commands {
		item := item
		cliCommand := cli.Command{
			Name:  item.Signature(),
			Usage: item.Description(),
			Action: func(ctx *cli.Context) error {
				return item.Handle(&CliContext{ctx})
			},
		}

		cliCommand.Category = item.Extend().Category
		cliCommand.Flags = flagsToCliFlags(item.Extend().Flags)
		c.instance.Commands = append(c.instance.Commands, &cliCommand)
	}
}

// Call Run an Artisan console command by name.
func (c *CliApplication) Call(command string) {
	c.Run(append([]string{os.Args[0], "artisan"}, strings.Split(command, " ")...), false)
}

// CallAndExit Run an Artisan console command by name and exit.
func (c *CliApplication) CallAndExit(command string) {
	c.Run(append([]string{os.Args[0], "artisan"}, strings.Split(command, " ")...), true)
}

// Run a command. Args come from os.Args.
func (c *CliApplication) Run(args []string, exitIfArtisan bool) {
	if len(args) >= 2 {
		if args[1] == "artisan" {
			if len(args) == 2 {
				args = append(args, "--help")
			}

			if args[2] != "-V" && args[2] != "--version" {
				cliArgs := append([]string{args[0]}, args[2:]...)
				if err := c.instance.Run(cliArgs); err != nil {
					panic(err.Error())
				}
			}

			printResult(args[2])

			if exitIfArtisan {
				os.Exit(0)
			}
		}
	}
}

func flagsToCliFlags(flags []Flag) []cli.Flag {
	var cliFlags []cli.Flag
	for _, flag := range flags {
		cliFlags = append(cliFlags, &cli.StringFlag{
			Name:     flag.Name,
			Aliases:  flag.Aliases,
			Usage:    flag.Usage,
			Required: flag.Required,
			Value:    flag.Value,
		})
	}

	return cliFlags
}

func printResult(command string) {
	switch command {
	case "make:command":
		color.Greenln("Console command created successfully")
	case "-V", "--version":
		color.Greenln(Support.Name + Support.Version)
	}
}
