package migrations

import (
	"fmt"

	"github.com/kmsar/laravel-go/Framework/Collections"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
	"github.com/kmsar/laravel-go/Framework/Logs"

	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Database/table"
)

func Transaction(sql string) IDatabase.MigrateHandler {
	return func(db IDatabase.DBConnection) error {
		return db.Transaction(func(executor IDatabase.SqlExecutor) error {
			_, err := executor.Exec(sql)
			return err
		})
	}
}

func Exec(sql string) IDatabase.MigrateHandler {
	return func(db IDatabase.DBConnection) error {
		_, err := db.Exec(sql)
		return err
	}
}

func getMigrations(db IDatabase.DBConnection, tableName string) Support.Collection {
	query := table.WithConnection(tableName, db)
	ddl := fmt.Sprintf("create table %s\n(\n    id        int unsigned auto_increment\n        primary key,\n    migration varchar(191) not null,\n    batch     int          not null\n)\n", tableName)
	_, err := db.Exec(ddl)

	if err == nil {
		Logs.Default().Info("migrations.getMigrations: Migration table has been generated")
		return Collections.MustNew([]Support.Fields{})
	}

	return query.Get()
}
