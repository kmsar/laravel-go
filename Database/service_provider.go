package Database

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConfig"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConsole"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEvent"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Database/drivers"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Database/migrations"
)

type ServiceProvider struct {
	migrations IDatabase.Migrations
}

func Service(migrations IDatabase.Migrations) IFoundation.ServiceProvider {
	return &ServiceProvider{migrations: migrations}
}

func (this *ServiceProvider) Register(application IFoundation.IApplication) {
	application.Instance("migrations", this.migrations)
	application.NamedSingleton("migrations.table", func(config IConfig.Config) string {
		return config.Get("database").(Config).Migrations
	})
	application.NamedSingleton("db.factory", func(config IConfig.Config, events IEvent.EventDispatcher) IDatabase.DBFactory {
		return &Factory{
			events:      events,
			config:      config,
			dbConfig:    config.Get("database").(Config),
			connections: make(map[string]IDatabase.DBConnection),
			drivers: map[string]IDatabase.DBConnector{
				"mysql":    drivers.MysqlConnector,
				"postgres": drivers.PostgresSqlConnector,
				"sqlite":   drivers.SqliteConnector,
				//"clickhouse": drivers.ClickHouseConnector,
			},
		}
	})
	application.NamedSingleton("db", func(config IConfig.Config, factory IDatabase.DBFactory) IDatabase.DBConnection {
		return factory.Connection()
	})

	application.Call(func(console IConsole.Console) {
		console.RegisterCommand("migrate", migrations.Migrate)
		console.RegisterCommand("migrate:rollback", migrations.Rollback)
		console.RegisterCommand("migrate:refresh", migrations.Refresh)
		console.RegisterCommand("migrate:reset", migrations.Reset)
		console.RegisterCommand("migrate:status", migrations.Status)
	})
}

func (this *ServiceProvider) Boot(application IFoundation.IApplication) {

}
func (this *ServiceProvider) Start() error {
	return nil
}

func (this *ServiceProvider) Stop() {
}
