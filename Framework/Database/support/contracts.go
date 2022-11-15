package support

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

type SqlxExecutor interface {
	DriverName() string
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	Get(dest interface{}, query string, args ...interface{}) (err error)
	Select(dest interface{}, query string, args ...interface{}) (err error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type Executor interface {
	DriverName() string
	Query(query string, args ...interface{}) (Support.Collection, error)
	Get(dest interface{}, query string, args ...interface{}) (err error)
	Select(dest interface{}, query string, args ...interface{}) (err error)
	Exec(query string, args ...interface{}) (IDatabase.Result, error)
}
