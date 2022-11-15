package commands

import (
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Logs"

	"io/ioutil"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Console/Commands"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConsole"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
	"os"
)

type runner struct {
	IConsole.Command
	app IFoundation.IApplication
}

func Runner(app IFoundation.IApplication) IConsole.Command {
	return &runner{
		Command: Commands.Base("run", "启动 goal"),
		app:     app,
	}
}

func (this *runner) Handle() interface{} {
	path := this.app.Get("path").(string)

	pidPath := path + "/goal.pid"

	_ = ioutil.WriteFile(pidPath, []byte(fmt.Sprintf("%d", os.Getpid())), os.ModePerm)

	if errors := this.app.Start(); len(errors) > 0 {
		Logs.WithField("errors", errors).Fatal("Goal started abnormally!")
	} else {
		_ = os.Remove(pidPath)
		Logs.Default().Info("goal closed")
	}
	return nil
}
