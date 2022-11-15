package Console

import (
	"github.com/gorhill/cronexpr"
	"github.com/kmsar/laravel-go/Framework/Console/inputs"
	"github.com/kmsar/laravel-go/Framework/Contracts/IConfig"
	"github.com/kmsar/laravel-go/Framework/Contracts/IConsole"
	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Contracts/IRedis"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	application "github.com/kmsar/laravel-go/Framework/Foundation/Application"
	"github.com/kmsar/laravel-go/Framework/Logs"
	"github.com/kmsar/laravel-go/Framework/Support/Carbon/carbon"
	"github.com/kmsar/laravel-go/Framework/Support/Exceptions"
	"github.com/kmsar/laravel-go/Framework/Support/Utils"
	"reflect"
	"time"
)

type Provider func(application IFoundation.IApplication) IConsole.Console

type ServiceProvider struct {
	ConsoleProvider Provider

	stopChan         chan bool
	serverIdChan     chan bool
	app              IFoundation.IApplication
	execRecords      map[int]time.Time
	exceptionHandler IExeption.ExceptionHandler
}

func (this *ServiceProvider) Register(application IFoundation.IApplication) {
	this.app = application
	this.exceptionHandler = application.Get("exceptions.handler").(IExeption.ExceptionHandler)

	application.NamedSingleton("console", func() IConsole.Console {
		console := this.ConsoleProvider(application)
		console.Schedule(console.GetSchedule())
		return console
	})
	application.NamedSingleton("scheduling", func(console IConsole.Console) IConsole.Schedule {
		return console.GetSchedule()
	})
	application.NamedSingleton("console.input", func() IConsole.ConsoleInput {
		return inputs.NewOSArgsInput()
	})
}

func (this *ServiceProvider) runScheduleEvents(events []IConsole.ScheduleEvent) {
	if len(events) > 0 {
		// 并发执行所有事件
		now := time.Now()
		for index, event := range events {
			lastExecTime := this.execRecords[index]
			nextTime := carbon.Time2Carbon(cronexpr.MustParse(event.Expression()).Next(lastExecTime))
			nowCarbon := carbon.Time2Carbon(now)
			if nextTime.DiffInSeconds(nowCarbon) == 0 {
				this.execRecords[index] = now
				go (func(event IConsole.ScheduleEvent) {
					defer func() {
						if err := recover(); err != nil {
							this.exceptionHandler.Handle(ScheduleEventException{
								Exception: Exceptions.WithRecover(err, Support.Fields{
									"expression": event.Expression(),
									"mutex_name": event.MutexName(),
									"one_server": event.OnOneServer(),
									"event":      Utils.GetTypeKey(reflect.TypeOf(event)),
								}),
							})
						}
					}()
					event.Run(this.app)
				})(event)
			} else if nextTime.Lt(nowCarbon) {
				this.execRecords[index] = now
			}
		}
	}
}

func (this *ServiceProvider) Start() error {
	this.execRecords = make(map[int]time.Time)
	go this.maintainServerId()
	this.app.Call(func(schedule IConsole.Schedule) {
		if len(schedule.GetEvents()) > 0 {
			this.stopChan = Utils.SetInterval(1, func() {
				this.runScheduleEvents(schedule.GetEvents())
			}, func() {
				Logs.Default().Info("the goal scheduling is closed")
			})
		}
	})
	return nil
}

func (this *ServiceProvider) Stop() {
	if this.stopChan != nil {
		this.stopChan <- true
	}
	if this.serverIdChan != nil {
		this.serverIdChan <- true
	}
}

// maintainServerId 维护服务实例ID
func (this *ServiceProvider) maintainServerId() {
	this.app.Call(func(redis IRedis.RedisConnection, config IConfig.Config, handler IExeption.ExceptionHandler) {
		appConfig := config.Get("app").(application.Config)
		this.serverIdChan = Utils.SetInterval(1, func() {
			// 维持当前服务心跳
			_, _ = redis.Set("goal.server."+appConfig.ServerId, time.Now().String(), time.Second*2)
		}, func() {
			_, _ = redis.Del("goal.server." + appConfig.ServerId)
		})
	})
}
