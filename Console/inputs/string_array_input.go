package inputs

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Console/arguments"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConsole"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"strings"
)

type StringArrayInput struct {
	ArgsArray []string
}

func StringArray(argsArray []string) StringArrayInput {
	return StringArrayInput{argsArray}
}

func (this *StringArrayInput) GetCommand() string {
	if len(this.ArgsArray) > 0 {
		return this.ArgsArray[0]
	}
	return ""
}

func (this *StringArrayInput) GetArguments() IConsole.CommandArguments {
	if len(this.ArgsArray) > 0 {
		args := make([]string, 0)
		options := Support.Fields{}

		for _, arg := range this.ArgsArray[1:] {
			if strings.HasPrefix(arg, "--") {
				if argArr := strings.Split(strings.ReplaceAll(arg, "--", ""), "="); len(argArr) > 1 {
					options[argArr[0]] = argArr[1]
				} else {
					options[argArr[0]] = true
				}
			} else if strings.HasPrefix(arg, "-") {
				if argArr := strings.Split(strings.ReplaceAll(arg, "-", ""), "="); len(argArr) > 1 {
					options[argArr[0]] = argArr[1]
				} else {
					options[argArr[0]] = true
				}
			} else {
				args = append(args, arg)
			}
		}

		return arguments.NewArguments(args, options)
	}
	return arguments.NewArguments([]string{}, nil)
}
