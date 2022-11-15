package migrations

import (
	"errors"
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Collections"
	"github.com/kmsar/laravel-go/Framework/Console/Commands"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Logs"

	//"github.com/goal-web/collection"
	//"github.com/goal-web/contracts"
	//"github.com/goal-web/supports/commands"
	//"github.com/goal-web/supports/logs"
	"github.com/kmsar/laravel-go/Framework/Contracts/IConsole"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Contracts/IRedis"
	"github.com/kmsar/laravel-go/Framework/Database/table"
)

func Migrate(app IFoundation.IApplication) IConsole.Command {
	return &migrate{
		production: app.IsProduction(),
		redis:      app.Get("redis").(IRedis.RedisConnection),
		db:         app.Get("db.factory").(IDatabase.DBFactory),
		table:      app.Get("migrations.table").(string),
		migrations: app.Get("migrations").(IDatabase.Migrations),
		Command:    Commands.Base("migrate {--force:", "Run the database migrations"),
	}
}

type migrate struct {
	Commands.Command
	production bool
	table      string
	redis      IRedis.RedisConnection
	db         IDatabase.DBFactory
	migrations IDatabase.Migrations
}

var MustForceErr = errors.New("use the force option in production")

func (this *migrate) Handle() interface{} {
	if this.production && !this.GetBool("force") {
		Logs.WithError(MustForceErr).Error("refresh.Handle: ")
		return MustForceErr
	}

	var (
		raw           = getMigrations(this.db.Connection(), this.table)
		executedNum   = 0
		migratedItems = raw.Pluck("migration")
	)

	batch := raw.Max("batch") + 1

	migrations := Collections.MustNew(this.migrations).Sort(func(migrate IDatabase.Migrate, next IDatabase.Migrate) bool {
		return migrate.CreatedAt.Before(next.CreatedAt)
	}).ToInterfaceArray()

	for _, item := range migrations {
		migration := item.(IDatabase.Migrate)
		if _, exists := migratedItems[migration.Name]; !exists {
			conn := this.db.Connection(migration.Connection)
			Logs.Default().Info(fmt.Sprintf("migrate.Handle: %s migrating", migration.Name))
			if err := migration.Up(conn); err != nil {
				Logs.Default().WithError(err).Error(fmt.Sprintf("migrate.Handle: %s failed to execute", migration.Name))
				panic(err)
			}
			Logs.Default().Info(fmt.Sprintf("migrate.Handle: %s migrated", migration.Name))
			executedNum++
			table.WithConnection(this.table, conn).Insert(Support.Fields{
				"batch":     batch,
				"migration": migration.Name,
			})
		}
	}

	if executedNum == 0 {
		Logs.Default().Info("migrate.Handle: No migration was performed")
	}

	return nil
}
