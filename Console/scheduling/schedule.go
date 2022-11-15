package scheduling

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Console/inputs"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConfig"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConsole"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IRedis"
	application "github.com/laravel-go-version/v2/pkg/Illuminate/Foundation/Application"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Logs"
	"os/exec"
)

type Schedule struct {
	store    string
	timezone string
	mutex    *Mutex
	app      IFoundation.IApplication

	events []IConsole.ScheduleEvent
}

func (this *Schedule) GetEvents() []IConsole.ScheduleEvent {
	return this.events
}

func (this *Schedule) UseStore(store string) {
	this.store = store
}

func NewSchedule(app IFoundation.IApplication) IConsole.Schedule {
	var (
		appConfig = app.Get("config").(IConfig.Config).Get("app").(application.Config)
		redis, _  = app.Get("redis.factory").(IRedis.RedisFactory)
	)
	return &Schedule{
		timezone: appConfig.Timezone,
		mutex: &Mutex{
			redis: redis,
			store: "cache",
		},
		app:    app,
		events: make([]IConsole.ScheduleEvent, 0),
	}
}

func (this *Schedule) Call(callback interface{}, args ...interface{}) IConsole.CallbackEvent {
	event := NewCallbackEvent(this.mutex, func() {
		this.app.Call(callback, args...)
	}, this.timezone)
	this.events = append(this.events, event)
	return event
}

func (this *Schedule) Command(command IConsole.Command, args ...string) IConsole.CommandEvent {
	args = append([]string{command.GetName()}, args...)
	input := inputs.StringArray(args)
	err := command.InjectArguments(input.GetArguments())
	if err != nil {
		Logs.WithError(err).WithField("args", args).Debug("Schedule.Command: arguments invalid")
		panic(err) // 因为这个阶段框架还没正式运行，所以 panic
	}
	event := NewCommandEvent(command.GetName(), this.mutex, func(console IConsole.Console) {
		command.Handle()
	}, this.timezone)
	this.events = append(this.events, event)
	return event
}

func (this *Schedule) Exec(command string, args ...string) IConsole.CommandEvent {
	var event = NewCommandEvent(command, this.mutex, func(console IConsole.Console) {
		if console.Exists(command) {
			args = append([]string{command}, args...)
			input := inputs.StringArray(args)
			console.Run(&input)
		} else {
			if err := exec.Command(command, args...).Run(); err != nil {
				Logs.WithError(err).
					WithField("command", command).
					WithField("args", args).
					Debug("Schedule.Exec: failed")
			}
		}

	}, this.timezone)
	this.events = append(this.events, event)
	return event
}
