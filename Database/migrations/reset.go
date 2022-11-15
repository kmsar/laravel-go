package migrations

import (
	"fmt"

	"github.com/laravel-go-version/v2/pkg/Illuminate/Collections"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Logs"

	"github.com/laravel-go-version/v2/pkg/Illuminate/Console/Commands"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IRedis"

	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConsole"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Database/table"
)

func Reset(app IFoundation.IApplication) IConsole.Command {
	return &reset{
		production: app.IsProduction(),
		redis:      app.Get("redis").(IRedis.RedisConnection),
		db:         app.Get("db.factory").(IDatabase.DBFactory),
		table:      app.Get("migrations.table").(string),
		migrations: app.Get("migrations").(IDatabase.Migrations),
		Command:    Commands.Base("migrate:reset {--force:是否在生产环境强制执行}", "Run the database migrations"),
	}
}

type reset struct {
	Commands.Command
	production bool
	table      string
	redis      IRedis.RedisConnection
	db         IDatabase.DBFactory
	migrations IDatabase.Migrations
}

func (this *reset) Handle() interface{} {
	if this.production && !this.GetBool("force") {
		Logs.WithError(MustForceErr).Error("refresh.Handle: ")
		return MustForceErr
	}

	// rollback all migrations
	if raw := getMigrations(this.db.Connection(), this.table); raw.Len() > 0 {
		var migrations = Collections.MustNew(this.migrations).Pluck("name")

		raw.Map(func(item Support.Fields) {
			migration, ok := migrations[item["migration"].(string)].(IDatabase.Migrate)
			if ok {
				conn := this.db.Connection(migration.Connection)
				Logs.Default().Info(fmt.Sprintf("rollback.Handle: %s rollbacking", migration.Name))
				if err := migration.Down(conn); err != nil {
					Logs.WithError(err).Error(fmt.Sprintf("rollback.Handle: %s failed to rollback", migration.Name))
					panic(err)
				}
				Logs.Default().Info(fmt.Sprintf("rollback.Handle: %s rollbacked", migration.Name))
				table.WithConnection(this.table, conn).Where("id", item["id"]).Delete()
			} else {
				Logs.Default().Warn(fmt.Sprintf("rollback.Handle: migration %s is not exists", migration.Name))
			}
		})
	} else {
		Logs.Default().Info("rollback.Handle: No migrations need to be rolled back")
	}

	return nil
}
