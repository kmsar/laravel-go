package providers

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IAuth"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Database/table"
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
