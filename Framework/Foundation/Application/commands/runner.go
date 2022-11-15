package commands

import (
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Logs"

	"github.com/kmsar/laravel-go/Framework/Console/Commands"
	"github.com/kmsar/laravel-go/Framework/Contracts/IConsole"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"io/ioutil"
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
