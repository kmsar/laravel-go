package drivers

import (
	"github.com/jmoiron/sqlx"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
	"github.com/kmsar/laravel-go/Framework/Contracts/IEvent"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Logs"
	"github.com/kmsar/laravel-go/Framework/Support/Field"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	*Base
}

func SqliteConnector(config Support.Fields, events IEvent.EventDispatcher) IDatabase.DBConnection {
	db, err := sqlx.Connect("sqlite3", Field.GetStringField(config, "database"))

	if err != nil {
		Logs.WithError(err).WithField("config", config).Fatal("sqlite 数据库初始化失败")
	}

	return &Sqlite{NewDriver(db, events)}
}
