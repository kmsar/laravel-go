package IConsole

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"time"
)

type Console interface {

	// Call run  console command by name.
	Call(command string, arguments CommandArguments) interface{}

	// Run  an incoming console command.
	Run(input ConsoleInput) interface{}

	// Schedule register Task Scheduler.
	Schedule(schedule Schedule)

	// GetSchedule get task schedule.
	GetSchedule() Schedule

	Exists(name string) bool

	// RegisterCommand Register the given commands.
	RegisterCommand(name string, command CommandProvider)
}

// command provider.
type CommandProvider func(application IFoundation.IApplication) Command

type Command interface {

	// Handle  an incoming console command.
	Handle() interface{}

	// InjectArguments Inject console command parameters.
	InjectArguments(arguments CommandArguments) error

	// GetSignature Get console command signature.
	GetSignature() string

	// GetName Get console command name.
	GetName() string

	// GetDescription Get console command description.
	GetDescription() string

	// GetHelp Get console command help
	GetHelp() string
}

type ConsoleInput interface {

	// GetCommand get console command.
	GetCommand() string

	// GetArguments Get console command parameters.
	GetArguments() CommandArguments
}

type CommandArguments interface {
	Support.FieldsProvider
	Support.Getter
	Support.OptionalGetter

	// GetArg Get the specified parameters of the console command.
	GetArg(index int) string

	// GetArgs Get console command parameters.
	GetArgs() []string

	// SetOption Set console command options.
	SetOption(key string, value interface{})

	// Exists Verify if the specified key exists.
	Exists(key string) bool

	// StringArrayOption Get the specified option, return []string type.
	StringArrayOption(key string, defaultValue []string) []string

	// IntArrayOption Get the specified option,  return []int type.
	IntArrayOption(key string, defaultValue []int) []int

	// Int64ArrayOption Get the specified option, return int 64 type.
	Int64ArrayOption(key string, defaultValue []int64) []int64

	// FloatArrayOption Get the specified option, return []float32 type.
	FloatArrayOption(key string, defaultValue []float32) []float32

	// Float64ArrayOption Get the specified option, return []float64 type.
	Float64ArrayOption(key string, defaultValue []float64) []float64
}

type Schedule interface {

	// UseStore Use the specified redis link.
	UseStore(store string)

	// Call Add a new callback event to the schedule.
	Call(callback interface{}, args ...interface{}) CallbackEvent

	// Command Add a new command event to the schedule.
	Command(command Command, args ...string) CommandEvent

	// Exec Add a new command event to the schedule.
	Exec(command string, args ...string) CommandEvent

	// GetEvents Get all the events on the schedule.
	GetEvents() []ScheduleEvent
}

type ScheduleEvent interface {

	// Run run the given event.
	Run(application IFoundation.IApplication)

	// WithoutOverlapping Do not allow the event to overlap each other.
	WithoutOverlapping(expiresAt int) ScheduleEvent

	// OnOneServer Allow the event to only run on one server for each cron expression.
	OnOneServer() ScheduleEvent

	// MutexName Get the mutex name for the scheduled command.
	MutexName() string

	// SetMutexName Set the mutex name for the scheduled command.
	SetMutexName(mutexName string) ScheduleEvent

	// Skip Register a callback to further filter the schedule, skip when returning true.
	Skip(callback func() bool) ScheduleEvent

	// When Register a callback to further filter the schedule, Do not skip when returning true.
	When(callback func() bool) ScheduleEvent

	// SpliceIntoPosition Splice the given value into the given position of the expression.
	SpliceIntoPosition(position int, value string) ScheduleEvent

	// Expression Get the cron expression for the event.
	Expression() string

	// Cron The cron expression representing the event's frequency.
	Cron(expression string) ScheduleEvent

	// Timezone Set the instance's timezone from a string or object.
	Timezone(timezone string) ScheduleEvent

	// Days Set the days of the week the command should run on.
	Days(day string, days ...string) ScheduleEvent

	// Years Set which years the command should run.
	Years(years ...string) ScheduleEvent

	// Yearly schedule the event to run yearly.
	Yearly() ScheduleEvent

	// YearlyOn schedule the event to run yearly on a given month, day, and time.
	YearlyOn(month time.Month, dayOfMonth int, time string) ScheduleEvent

	// Quarterly schedule the event to run quarterly.
	Quarterly() ScheduleEvent

	// LastDayOfMonth schedule the event to run on the last day of the month.
	LastDayOfMonth(time string) ScheduleEvent

	// TwiceMonthly schedule the event to run twice monthly at a given time.
	TwiceMonthly(first, second int, time string) ScheduleEvent

	// Monthly schedule the event to run monthly.
	Monthly() ScheduleEvent

	// MonthlyOn schedule the event to run monthly on a given day and time.
	MonthlyOn(dayOfMonth int, time string) ScheduleEvent

	// WeeklyOn schedule the event to run weekly on a given day and time.
	WeeklyOn(dayOfWeek time.Weekday, time string) ScheduleEvent

	// Weekly schedule the event to run weekly.
	Weekly() ScheduleEvent

	// Sundays schedule the event to run only on sundays.
	Sundays() ScheduleEvent

	// Saturdays schedule the event to run only on saturdays.
	Saturdays() ScheduleEvent

	// Fridays schedule the event to run only on fridays.
	Fridays() ScheduleEvent

	// Thursdays schedule the event to run only on thursdays.
	Thursdays() ScheduleEvent

	// Wednesdays schedule the event to run only on wednesdays.
	Wednesdays() ScheduleEvent

	// Tuesdays schedule the event to run only on tuesdays.
	Tuesdays() ScheduleEvent

	// Mondays schedule the event to run only on mondays.
	Mondays() ScheduleEvent

	// Weekends schedule the event to run only on weekends.
	Weekends() ScheduleEvent

	// Weekdays schedule the event to run only on weekdays.
	Weekdays() ScheduleEvent

	// TwiceDailyAt schedule the event to run twice daily at a given offset.
	TwiceDailyAt(first, second, offset int) ScheduleEvent

	// TwiceDaily schedule the event to run twice daily.
	TwiceDaily(first, second int) ScheduleEvent

	// DailyAt schedule the event to run daily at a given time (10:00, 19:30, etc).
	DailyAt(time string) ScheduleEvent

	// Daily schedule the event to run daily.
	Daily() ScheduleEvent

	// EverySixHours schedule the event to run every six hours.
	EverySixHours() ScheduleEvent

	// EveryFourHours schedule the event to run every four hours.
	EveryFourHours() ScheduleEvent

	// EveryThreeHours schedule the event to run every three hours.
	EveryThreeHours() ScheduleEvent

	// EveryTwoHours schedule the event to run every two hours.
	EveryTwoHours() ScheduleEvent

	// HourlyAt schedule the event to run hourly at a given offset in the hour.
	HourlyAt(offset ...int) ScheduleEvent

	// Hourly schedule the event to run hourly.
	Hourly() ScheduleEvent

	// EveryThirtyMinutes schedule the event to run every thirty minutes.
	EveryThirtyMinutes() ScheduleEvent

	// EveryFifteenMinutes schedule the event to run every fifteen minutes.
	EveryFifteenMinutes() ScheduleEvent

	// EveryTenMinutes schedule the event to run every ten minutes.
	EveryTenMinutes() ScheduleEvent

	// EveryFiveMinutes schedule the event to run every five minutes.
	EveryFiveMinutes() ScheduleEvent

	// EveryFourMinutes schedule the event to run every four minutes.
	EveryFourMinutes() ScheduleEvent

	// EveryThreeMinutes schedule the event to run every three minutes.
	EveryThreeMinutes() ScheduleEvent

	// EveryTwoMinutes schedule the event to run every two minutes.
	EveryTwoMinutes() ScheduleEvent

	// EveryMinute schedule the event to run every minute.
	EveryMinute() ScheduleEvent

	// UnlessBetween schedule the event to not run between start and end time.
	UnlessBetween(startTime, endTime string) ScheduleEvent

	// Between schedule the event to run between start and end time.
	Between(startTime, endTime string) ScheduleEvent

	// EveryThirtySeconds schedule the event to run every 30 seconds.
	EveryThirtySeconds() ScheduleEvent

	// EveryFifteenSeconds schedule the event to run every 15 seconds.
	EveryFifteenSeconds() ScheduleEvent

	// EveryTenSeconds schedule the event to run every 10 seconds.
	EveryTenSeconds() ScheduleEvent

	// EveryFiveSeconds schedule the event to run every 5 seconds.
	EveryFiveSeconds() ScheduleEvent

	// EveryFourSeconds schedule the event to run every 4 seconds.
	EveryFourSeconds() ScheduleEvent

	// EveryThreeSeconds schedule the event to run every 3 seconds.
	EveryThreeSeconds() ScheduleEvent

	// EveryTwoSeconds schedule the event to run every 2 seconds.
	EveryTwoSeconds() ScheduleEvent

	// EverySecond schedule the event to run every second.
	EverySecond() ScheduleEvent
}

type CallbackEvent interface {
	ScheduleEvent

	// Description Set the human-friendly description of the event.
	Description(description string) CallbackEvent
}
type CommandEvent interface {
	ScheduleEvent
}
