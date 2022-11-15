package migrations

import (
	"fmt"

	"github.com/kmsar/laravel-go/Framework/Collections"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Logs"

	"github.com/kmsar/laravel-go/Framework/Console/Commands"
	"github.com/kmsar/laravel-go/Framework/Contracts/IConsole"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Contracts/IRedis"
	"github.com/kmsar/laravel-go/Framework/Database/table"
)

func Refresh(app IFoundation.IApplication) IConsole.Command {
	return &refresh{
		production: app.IsProduction(),
		redis:      app.Get("redis").(IRedis.RedisConnection),
		db:         app.Get("db.factory").(IDatabase.DBFactory),
		table:      app.Get("migrations.table").(string),
		migrations: app.Get("migrations").(IDatabase.Migrations),
		Command:    Commands.Base("migrate:refresh {--force:是否在生产环境强制执行}", "Run the database migrations"),
	}
}

type refresh struct {
	Commands.Command
	production bool
	table      string
	redis      IRedis.RedisConnection
	db         IDatabase.DBFactory
	migrations IDatabase.Migrations
}

func (this *refresh) Handle() interface{} {
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
	}

	var (
		migrations = Collections.MustNew(this.migrations).Sort(func(migrate IDatabase.Migrate, next IDatabase.Migrate) bool {
			return migrate.CreatedAt.Before(next.CreatedAt)
		}).ToInterfaceArray()
		executedNum = 0
	)

	for _, item := range migrations {
		migration := item.(IDatabase.Migrate)
		conn := this.db.Connection(migration.Connection)
		Logs.Default().Info(fmt.Sprintf("migrate.Handle: %s migrating", migration.Name))
		if err := migration.Up(conn); err != nil {
			Logs.Default().WithError(err).Error(fmt.Sprintf("migrate.Handle: %s failed to execute", migration.Name))
			return err
		}
		Logs.Default().Info(fmt.Sprintf("migrate.Handle: %s migrated", migration.Name))
		executedNum++
		table.WithConnection(this.table, conn).Insert(Support.Fields{
			"batch":     1,
			"migration": migration.Name,
		})
	}

	if executedNum == 0 {
		Logs.Default().Info("migrate.Handle: No migration was performed")
	}

	return nil
}
