package migrations

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Console/Commands"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConsole"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IRedis"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Str/table"
)

func Status(app IFoundation.IApplication) IConsole.Command {
	return &status{
		production: app.IsProduction(),
		redis:      app.Get("redis").(IRedis.RedisConnection),
		db:         app.Get("db.factory").(IDatabase.DBFactory),
		table:      app.Get("migrations.table").(string),
		migrations: app.Get("migrations").(IDatabase.Migrations),
		Command:    Commands.Base("migrate:status", "Run the database migrations"),
	}
}

type status struct {
	Commands.Command
	production bool
	table      string
	redis      IRedis.RedisConnection
	db         IDatabase.DBFactory
	migrations IDatabase.Migrations
}

type MigrationStatus struct {
	Ran       string
	Migration string
	Batch     interface{}
}

func (this *status) Handle() interface{} {

	var (
		migrated = getMigrations(this.db.Connection(), this.table).Pluck("migration")
		data     = make([]MigrationStatus, 0)
	)

	for _, migration := range this.migrations {
		if migratedItem, exists := migrated[migration.Name].(Support.Fields); exists {
			data = append(data, MigrationStatus{
				Ran:       "Yes",
				Migration: migration.Name,
				Batch:     migratedItem["batch"],
			})
		} else {
			data = append(data, MigrationStatus{
				Ran:       "No",
				Migration: migration.Name,
				Batch:     0,
			})
		}
	}

	table.Output(data)

	return nil
}
