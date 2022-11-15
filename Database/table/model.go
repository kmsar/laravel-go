package table

import "github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"

type BaseModel struct {
	class      Support.Class
	table      string
	connection string
	primaryKey string
}

func Model(class Support.Class, table string, connection ...string) *Table {
	return FromModel(NewModel(class, table, connection...))
}

func NewModel(class Support.Class, table string, connection ...string) BaseModel {
	conn := ""
	if len(connection) > 0 {
		conn = connection[0]
	}
	return BaseModel{class: class, table: table, connection: conn}
}

func (model BaseModel) GetClass() Support.Class {
	return model.class
}

func (model BaseModel) GetPrimaryKey() string {
	if model.primaryKey == "" {
		return "id"
	}
	return model.primaryKey
}

func (model BaseModel) SetPrimaryKey(key string) {
	model.primaryKey = key
}

func (model BaseModel) GetTable() string {
	return model.table
}

func (model BaseModel) GetConnection() string {
	return model.connection
}
