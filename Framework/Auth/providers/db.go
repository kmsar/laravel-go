package providers

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IAuth"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Database/table"
)

type DB struct {
	model IDatabase.Model
}

func DBDriver(config Support.Fields) IAuth.UserProvider {
	return &DB{model: config["model"].(IDatabase.Model)}
}

func (db *DB) RetrieveById(identifier string) IAuth.Authenticatable {
	if user := table.FromModel(db.model).Where(db.model.GetPrimaryKey(), identifier).First(); user != nil {
		return user.(IAuth.Authenticatable)
	}
	return nil
}
