package tx

import (
	"github.com/jmoiron/sqlx"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
	"github.com/kmsar/laravel-go/Framework/Contracts/IEvent"
	"github.com/kmsar/laravel-go/Framework/Database/support"
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
