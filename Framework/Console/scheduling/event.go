package scheduling

import (
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IConsole"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Support/Carbon/carbon"
	"strconv"
	"strings"
	"time"
)

func NewEvent(mutex *Mutex, callback interface{}, timezone string) *Event {
	return &Event{
		callback:           callback,
		mutex:              mutex,
		filters:            make([]filter, 0),
		rejects:            make([]filter, 0),
		beforeCallbacks:    make([]func(), 0),
		afterCallbacks:     make([]func(), 0),
		withoutOverlapping: false,
		onOneServer:        false,
		timezone:           timezone,
		expression:         "0 * * * * * *",
		mutexName:          "",
		expiresAt:          0,
	}
}

type filter func() bool

type Event struct {
	callback interface{}

	mutex           *Mutex
	filters         []filter
	rejects         []filter
	beforeCallbacks []func()
	afterCallbacks  []func()

	withoutOverlapping bool
	onOneServer        bool

	timezone   string
	expression string
	mutexName  string
	expiresAt  time.Duration
}

func (this *Event) Years(years ...string) IConsole.ScheduleEvent {
	if len(years) > 0 {
		return this.SpliceIntoPosition(6, strings.Join(years, ","))
	}
	return this
}

func (this *Event) Expression() string {
	return this.expression
}

func (this *Event) EveryThirtySeconds() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(0, "0,30")
}

func (this *Event) EveryFifteenSeconds() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(0, "*/15")
}

func (this *Event) EveryTenSeconds() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(0, "*/10")
}

func (this *Event) EveryFiveSeconds() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(0, "*/5")
}

func (this *Event) EveryFourSeconds() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(0, "*/4")
}

func (this *Event) EveryThreeSeconds() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(0, "*/3")
}

func (this *Event) EveryTwoSeconds() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(0, "*/2")
}

func (this *Event) EverySecond() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(0, "*")
}

func (this *Event) WithoutOverlapping(expiresAt int) IConsole.ScheduleEvent {
	this.expiresAt = time.Duration(expiresAt) * time.Second
	this.withoutOverlapping = true
	return this.Skip(func() bool {
		return this.mutex.Exists(this)
	})
}

func (this *Event) Run(application IFoundation.IApplication) {
	if !this.FiltersPass() {
		return
	}
	defer this.removeMutex()
	if this.withoutOverlapping && !this.mutex.Create(this) {
		return
	}
	application.Call(this.callback)
	return
}

func (this *Event) removeMutex() {
	if this.withoutOverlapping {
		this.mutex.Forget(this)
	}
}
func (this *Event) OnOneServer() IConsole.ScheduleEvent {
	this.onOneServer = true
	return this
}

func (this *Event) Timezone(timezone string) IConsole.ScheduleEvent {
	this.timezone = timezone
	return this
}

func (this *Event) Days(day string, days ...string) IConsole.ScheduleEvent {
	days = append([]string{day}, days...)
	return this.SpliceIntoPosition(5, strings.Join(days, ","))
}

func (this *Event) YearlyOn(month time.Month, dayOfMonth int, timeStr string) IConsole.ScheduleEvent {
	this.DailyAt(timeStr)

	return this.SpliceIntoPosition(3, strconv.Itoa(dayOfMonth)).
		SpliceIntoPosition(4, strconv.Itoa(int(month)))
}

func (this *Event) Yearly() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "0").
		SpliceIntoPosition(3, "1").
		SpliceIntoPosition(4, "1")
}

func (this *Event) Quarterly() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "0").
		SpliceIntoPosition(3, "1").
		SpliceIntoPosition(4, "1-12/3")
}

func (this *Event) LastDayOfMonth(timeStr string) IConsole.ScheduleEvent {
	this.DailyAt(timeStr)

	return this.When(func() bool {
		return carbon.Now(this.timezone).Day() == carbon.Now(this.timezone).EndOfMonth().Day()
	})
}

func (this *Event) TwiceMonthly(first, second int, timeStr string) IConsole.ScheduleEvent {
	this.DailyAt(timeStr)
	return this.SpliceIntoPosition(3, fmt.Sprintf("%d,%d", first, second))
}

func (this *Event) MonthlyOn(dayOfMonth int, timeStr string) IConsole.ScheduleEvent {
	this.DailyAt(timeStr)
	return this.SpliceIntoPosition(3, strconv.Itoa(dayOfMonth))
}

func (this *Event) Monthly() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "0").
		SpliceIntoPosition(3, "1")
}

func (this *Event) WeeklyOn(dayOfWeek time.Weekday, timeStr string) IConsole.ScheduleEvent {
	return this.DailyAt(timeStr).Days(strconv.Itoa(int(dayOfWeek)))
}

func (this *Event) Weekly() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "0").
		SpliceIntoPosition(5, "0")
}

func (this *Event) Sundays() IConsole.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d", time.Sunday))
}

func (this *Event) Saturdays() IConsole.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d", time.Saturday))
}

func (this *Event) Fridays() IConsole.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d", time.Friday))
}

func (this *Event) Thursdays() IConsole.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d", time.Thursday))
}

func (this *Event) Wednesdays() IConsole.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d", time.Wednesday))
}

func (this *Event) Tuesdays() IConsole.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d", time.Tuesday))
}

func (this *Event) Mondays() IConsole.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d", time.Monday))
}

func (this *Event) Weekends() IConsole.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d,%d", time.Saturday, time.Sunday))
}

func (this *Event) Weekdays() IConsole.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d-%d", time.Monday, time.Friday))
}

func (this *Event) TwiceDailyAt(first, second, offset int) IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, strconv.Itoa(offset)).
		SpliceIntoPosition(2, fmt.Sprintf("%d,%d", first, second))
}

func (this *Event) TwiceDaily(first, second int) IConsole.ScheduleEvent {
	return this.TwiceDailyAt(first, second, 0)
}

func (this *Event) DailyAt(timeStr string) IConsole.ScheduleEvent {
	segments := strings.Split(timeStr, ":")
	this.SpliceIntoPosition(2, segments[0])

	if len(segments) == 2 {
		return this.SpliceIntoPosition(1, segments[1])
	} else {
		return this.SpliceIntoPosition(1, "0")
	}
}

func (this *Event) Daily() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "0")
}

func (this *Event) EverySixHours() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "*/6")
}

func (this *Event) EveryFourHours() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "*/4")
}

func (this *Event) EveryThreeHours() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "*/3")
}

func (this *Event) EveryTwoHours() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "*/2")
}

func (this *Event) HourlyAt(offset ...int) IConsole.ScheduleEvent {
	offsetStrings := make([]string, 0)
	for _, offsetInt := range offset {
		offsetStrings = append(offsetStrings, strconv.Itoa(offsetInt))
	}
	return this.SpliceIntoPosition(1, strings.Join(offsetStrings, ","))
}

func (this *Event) Hourly() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0")
}

func (this *Event) EveryThirtyMinutes() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0,30")
}

func (this *Event) EveryFifteenMinutes() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "*/15")
}

func (this *Event) EveryTenMinutes() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "*/10")
}

func (this *Event) EveryFiveMinutes() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "*/5")
}

func (this *Event) EveryFourMinutes() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "*/4")
}

func (this *Event) EveryThreeMinutes() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "*/3")
}

func (this *Event) EveryTwoMinutes() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "*/2")
}

func (this *Event) EveryMinute() IConsole.ScheduleEvent {
	return this.SpliceIntoPosition(1, "*")
}

func (this *Event) FiltersPass() bool {
	for _, filter := range this.filters {
		if !filter() {
			return false
		}
	}
	for _, reject := range this.rejects {
		if reject() {
			return false
		}
	}
	return true
}
func (this *Event) When(filter func() bool) IConsole.ScheduleEvent {
	this.filters = append(this.filters, filter)
	return this
}
func (this *Event) Skip(reject func() bool) IConsole.ScheduleEvent {
	this.rejects = append(this.rejects, reject)
	return this
}

func (this *Event) Cron(expression string) IConsole.ScheduleEvent {
	this.expression = expression
	return this
}

func (this *Event) Between(startTime, endTimeStr string) IConsole.ScheduleEvent {
	return this.When(this.inTimeInterval(startTime, endTimeStr))
}

func (this *Event) UnlessBetween(startTime, endTimeStr string) IConsole.ScheduleEvent {
	return this.Skip(this.inTimeInterval(startTime, endTimeStr))
}

func (this *Event) inTimeInterval(startTime, endTimeStr string) func() bool {
	var (
		startAt = carbon.Now().ParseByFormat(startTime, "H:i", this.timezone)
		endAt   = carbon.Now().ParseByFormat(endTimeStr, "H:i", this.timezone)
	)

	if endAt.Lt(startAt) {
		if startAt.Gt(carbon.Now(this.timezone).SetYear(0000).SetMonth(1).SetDay(1)) {
			startAt.SubDay()
		} else {
			endAt.AddDay()
		}
	}

	return func() bool {
		now := carbon.Now(this.timezone).SetYear(0000).SetMonth(1).SetDay(1)
		return now.Between(startAt, endAt)
	}
}

func (this *Event) MutexName() string {
	return this.mutexName
}

func (this *Event) SetMutexName(mutexName string) IConsole.ScheduleEvent {
	this.mutexName = mutexName
	return this
}

func (this *Event) SpliceIntoPosition(position int, value string) IConsole.ScheduleEvent {
	segments := strings.Split(this.expression, " ")
	segments[position] = value
	return this.Cron(strings.Join(segments, " "))
}
