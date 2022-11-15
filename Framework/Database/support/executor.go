package support

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
	"github.com/kmsar/laravel-go/Framework/Contracts/IEvent"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Database/events"
	"github.com/kmsar/laravel-go/Framework/Support/Exceptions"
	"time"
)

type BaseExecutor struct {
	executor SqlxExecutor
	events   IEvent.EventDispatcher
	wrapper  func(sql string) string
}

func (this *BaseExecutor) DriverName() string {
	return this.executor.DriverName()
}

func (this *BaseExecutor) getStatement(sql string) string {
	if this.wrapper != nil {
		return this.wrapper(sql)
	}
	return sql
}

func NewExecutor(executor SqlxExecutor, dispatcher IEvent.EventDispatcher, wrapper func(sql string) string) Executor {
	return &BaseExecutor{
		executor: executor,
		events:   dispatcher,
		wrapper:  wrapper,
	}
}

func (this *BaseExecutor) Query(query string, args ...interface{}) (results Support.Collection, err error) {
	query = this.getStatement(query)
	var timeConsuming time.Duration
	defer func() {
		if err == nil {
			err = Exceptions.ResolveException(recover())
		}
		if this.events != nil {
			this.events.Dispatch(&events.QueryExecuted{Sql: query, Bindings: args, Time: timeConsuming, Error: err})
		}
	}()
	var startAt = time.Now()
	rows, err := this.executor.Queryx(query, args...)
	timeConsuming = time.Now().Sub(startAt)
	if err != nil {
		return nil, err
	}

	return ParseRowsToCollection(rows)
}

func (this *BaseExecutor) Get(dest interface{}, query string, args ...interface{}) (err error) {
	query = this.getStatement(query)
	var startAt = time.Now()
	defer func() {
		if err == nil {
			err = Exceptions.ResolveException(recover())
		}

		if this.events != nil {
			this.events.Dispatch(&events.QueryExecuted{Sql: query, Bindings: args, Time: time.Now().Sub(startAt), Error: err})
		}
	}()
	return this.executor.Get(dest, query, args...)
}

func (this *BaseExecutor) Select(dest interface{}, query string, args ...interface{}) (err error) {
	query = this.getStatement(query)
	var startAt = time.Now()
	defer func() {
		if err == nil {
			err = Exceptions.ResolveException(recover())
		}
		if this.events != nil {
			this.events.Dispatch(&events.QueryExecuted{Sql: query, Bindings: args, Time: time.Now().Sub(startAt), Error: err})
		}
	}()
	return this.executor.Get(dest, query, args...)
}

func (this *BaseExecutor) Exec(query string, args ...interface{}) (result IDatabase.Result, err error) {
	query = this.getStatement(query)
	var startAt = time.Now()
	defer func() {
		if err == nil {
			err = Exceptions.ResolveException(recover())
		}
		if this.events != nil {
			this.events.Dispatch(&events.QueryExecuted{Sql: query, Bindings: args, Time: time.Now().Sub(startAt), Error: err})
		}
	}()
	return this.executor.Exec(query, args...)
}
