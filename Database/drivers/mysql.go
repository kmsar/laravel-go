package drivers

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEvent"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Logs"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Field"
)

type Mysql struct {
	*Base
}

func MysqlConnector(config Support.Fields, events IEvent.EventDispatcher) IDatabase.DBConnection {
	dsn := Field.GetStringField(config, "unix_socket")
	if dsn == "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
			Field.GetStringField(config, "username"),
			Field.GetStringField(config, "password"),
			Field.GetStringField(config, "host"),
			Field.GetStringField(config, "port"),
			Field.GetStringField(config, "database"),
			Field.GetStringField(config, "charset"),
		)
	}
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		Logs.WithError(err).WithField("config", config).Fatal("mysql connection error")
	}
	db.SetMaxOpenConns(Field.GetIntField(config, "max_connections"))
	db.SetMaxIdleConns(Field.GetIntField(config, "max_idles"))

	return &Mysql{NewDriver(db, events)}
}
