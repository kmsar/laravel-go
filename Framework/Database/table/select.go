package table

import (
	"database/sql"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Exceptions"
)

func (this *Table) fetch(query string, bindings ...interface{}) Support.Collection {
	rows, err := this.getExecutor().Query(query, bindings...)
	if err != nil && err != sql.ErrNoRows {
		panic(SelectException{Exceptions.WithError(err, Support.Fields{"sql": query, "bindings": bindings})})
	}

	if this.class == nil {
		return rows
	}

	return rows.Map(func(fields Support.Fields) interface{} {
		return this.class.NewByTag(fields, "db")
	})
}

func (this *Table) Get() Support.Collection {
	queryStatement, bindings := this.SelectSql()

	return this.fetch(queryStatement, bindings...)
}

func (this *Table) SelectForUpdate() Support.Collection {
	queryStatement, bindings := this.SelectForUpdateSql()

	return this.fetch(queryStatement, bindings...)
}

func (this *Table) Find(key interface{}) interface{} {
	return this.Where(this.primaryKey, key).First()
}

func (this *Table) First() interface{} {
	return this.Take(1).Get().First()
}

func (this *Table) FirstOrFail() interface{} {
	if result := this.First(); result != nil {
		return result
	}
	panic(NotFoundException{Exceptions.WithError(sql.ErrNoRows, nil)})
}
