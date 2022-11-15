package drivers

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
	"github.com/kmsar/laravel-go/Framework/Contracts/IEvent"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Logs"
	"github.com/kmsar/laravel-go/Framework/Support/Field"
	_ "github.com/lib/pq"
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
