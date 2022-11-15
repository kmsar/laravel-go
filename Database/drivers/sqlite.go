package drivers

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEvent"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Logs"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Field"
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
