package Database

import (
	"errors"
	"github.com/kmsar/laravel-go/Framework/Support/Field"

	"github.com/kmsar/laravel-go/Framework/Contracts/IConfig"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
	"github.com/kmsar/laravel-go/Framework/Contracts/IEvent"
)

type Factory struct {
	events      IEvent.EventDispatcher
	config      IConfig.Config
	connections map[string]IDatabase.DBConnection
	drivers     map[string]IDatabase.DBConnector
	dbConfig    Config
}

func (this *Factory) Connection(name ...string) IDatabase.DBConnection {
	connection := this.dbConfig.Default
	if len(name) > 0 && name[0] != "" {
		connection = name[0]
	}
	if conn, existsConnection := this.connections[connection]; existsConnection {
		return conn
	}

	this.connections[connection] = this.make(connection)

	return this.connections[connection]
}

func (this *Factory) Extend(name string, driver IDatabase.DBConnector) {
	this.drivers[name] = driver
}

func (this *Factory) make(name string) IDatabase.DBConnection {
	config := this.config.Get("database").(Config)

	if connectionConfig, existsConnection := config.Connections[name]; existsConnection {
		driverName := Field.GetStringField(connectionConfig, "driver")
		if driver, existsDriver := this.drivers[driverName]; existsDriver {
			return driver(connectionConfig, this.events)
		}

		panic(DBConnectionException{
			error:  errors.New("driver error：" + driverName),
			Code:   DbDriverDontExist,
			fields: connectionConfig,
		})
	}

	panic(DBConnectionException{
		error:      errors.New("db connection exception：" + name),
		Code:       DbConnectionDontExist,
		Connection: name,
	})
}
