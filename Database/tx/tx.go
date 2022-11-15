package tx

import (
	"github.com/jmoiron/sqlx"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEvent"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Database/support"
)

type Tx struct {
	tx *sqlx.Tx
	support.Executor
	events IEvent.EventDispatcher
}

func (t *Tx) Commit() error {
	return t.tx.Commit()
}

func (t Tx) Rollback() error {
	return t.tx.Rollback()
}

func New(tx *sqlx.Tx, events IEvent.EventDispatcher) IDatabase.DBTx {
	return &Tx{
		tx:       tx,
		Executor: support.NewExecutor(tx, events, nil),
		events:   events,
	}
}
