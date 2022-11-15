package drivers

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEvent"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Logs"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Field"
)

type PostgresSQL struct {
	*Base
}

func PostgresSqlConnector(config Support.Fields, events IEvent.EventDispatcher) IDatabase.DBConnection {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		Field.GetStringField(config, "host"),
		Field.GetStringField(config, "port"),
		Field.GetStringField(config, "username"),
		Field.GetStringField(config, "password"),
		Field.GetStringField(config, "database"),
		Field.GetStringField(config, "sslmode"),
	))
	if err != nil {
		Logs.WithError(err).WithField("config", config).Fatal("postgreSql ")
	}
	db.SetMaxOpenConns(Field.GetIntField(config, "max_connections"))
	db.SetMaxIdleConns(Field.GetIntField(config, "max_idles"))

	return &PostgresSQL{WithWrapper(db, events, dollarNParamBindWrapper)}
}
