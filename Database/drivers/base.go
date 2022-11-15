package drivers

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEvent"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Database/exceptions"
	"strings"

	"github.com/laravel-go-version/v2/pkg/Illuminate/Database/support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Database/tx"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Exceptions"
)

type Base struct {
	support.Executor
	db     *sqlx.DB
	events IEvent.EventDispatcher
}

func NewDriver(db *sqlx.DB, dispatcher IEvent.EventDispatcher) *Base {
	return &Base{
		db:       db,
		Executor: support.NewExecutor(db, dispatcher, nil),
		events:   dispatcher,
	}
}

func WithWrapper(db *sqlx.DB, dispatcher IEvent.EventDispatcher, wrapper func(string) string) *Base {
	return &Base{
		db:       db,
		Executor: support.NewExecutor(db, dispatcher, wrapper),
		events:   dispatcher,
	}
}

func (base *Base) Begin() (IDatabase.DBTx, error) {
	sqlxTx, err := base.db.Beginx()
	if err != nil {
		return nil, err
	}
	return tx.New(sqlxTx, base.events), err
}

func (base *Base) Transaction(fn func(tx IDatabase.SqlExecutor) error) (err error) {
	sqlxTx, err := base.Begin()
	if err != nil {
		return exceptions.BeginException{Exception: Exceptions.WithError(err, nil)}
	}

	defer func() { // 处理 panic 情况
		if recoverErr := recover(); recoverErr != nil {
			rollbackErr := sqlxTx.Rollback()
			err = recoverErr.(error)
			if rollbackErr != nil {
				err = exceptions.RollbackException{Exception: Exceptions.WithPrevious(rollbackErr, nil, Exceptions.WithError(err, nil))}
			} else {
				err = exceptions.TransactionException{Exception: Exceptions.WithError(err, nil)}
			}
		}
	}()

	err = fn(sqlxTx)

	if err != nil {
		rollbackErr := sqlxTx.Rollback()
		if rollbackErr != nil {
			return exceptions.RollbackException{Exception: Exceptions.WithPrevious(rollbackErr, nil, Exceptions.WithError(err, nil))}
		}
		return exceptions.TransactionException{Exception: Exceptions.WithError(err, nil)}
	}

	return sqlxTx.Commit()
}

func dollarNParamBindWrapper(sql string) (result string) {
	var (
		parts    = strings.Split(sql, "?")
		partsLen = len(parts)
	)
	if partsLen == 1 {
		return sql
	}
	result = parts[0]
	for i := 1; i < partsLen; i++ {
		result = fmt.Sprintf("%s$%d%s", result, i, parts[i])
	}
	return
}
