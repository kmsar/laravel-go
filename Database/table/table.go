package table

import (
	//"github.com/goal-web/application"
	//"github.com/goal-web/contracts"
	//"github.com/goal-web/querybuilder"
	//"github.com/goal-web/supports/exceptions"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Database/querybuilder"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Exceptions"
	"next-doc/app/application"
)

type Table struct {
	IDatabase.QueryBuilder
	executor IDatabase.SqlExecutor

	table      string
	primaryKey string
	class      Support.Class
}

func getTable(name string) *Table {
	builder := querybuilder.NewQuery(name)
	instance := &Table{
		QueryBuilder: builder,
		primaryKey:   "id",
		table:        name,
	}
	builder.Bind(instance)
	return instance
}

func Query(name string) *Table {
	return getTable(name).SetConnection(application.App().Get("db").(IDatabase.DBConnection))
}

func FromModel(model IDatabase.Model) *Table {
	return WithConnection(model.GetTable(), model.GetConnection()).SetClass(model.GetClass()).SetPrimaryKey(model.GetPrimaryKey())
}

func WithConnection(name string, connection interface{}) *Table {
	if connection == "" || connection == nil {
		return Query(name)
	}
	return getTable(name).SetConnection(connection)
}

func WithTX(name string, tx IDatabase.DBTx) IDatabase.QueryBuilder {
	return getTable(name).SetExecutor(tx)
}

// SetConnection 参数要么是 contracts.DBConnection 要么是 string
func (this *Table) SetConnection(connection interface{}) *Table {
	if conn, ok := connection.(IDatabase.DBConnection); ok {
		this.executor = conn
	} else {
		this.executor = application.App(nil).Get("db.factory").(IDatabase.DBFactory).Connection(connection.(string))
	}
	return this
}

func (this *Table) SetClass(class Support.Class) *Table {
	this.class = class
	return this
}

func (this *Table) SetPrimaryKey(name string) *Table {
	this.primaryKey = name
	return this
}

func (this *Table) getExecutor() IDatabase.SqlExecutor {
	return this.executor
}

// SetExecutor 参数必须是 contracts.DBTx 实例
func (this *Table) SetExecutor(executor IDatabase.SqlExecutor) IDatabase.QueryBuilder {
	this.executor = executor
	return this
}

func (this *Table) Delete() int64 {
	sql, bindings := this.DeleteSql()
	result, err := this.getExecutor().Exec(sql, bindings...)
	if err != nil {
		panic(DeleteException{Exceptions.WithError(err, Support.Fields{"sql": sql, "bindings": bindings})})
	}
	num, err := result.RowsAffected()
	if err != nil {
		panic(DeleteException{Exceptions.WithError(err, Support.Fields{"sql": sql, "bindings": bindings})})
	}
	return num
}
