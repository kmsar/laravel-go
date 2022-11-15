package table

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Logs"
	"github.com/kmsar/laravel-go/Framework/Support/Exceptions"
	"github.com/kmsar/laravel-go/Framework/Support/Field"
)

func (this *Table) Create(fields Support.Fields) interface{} {
	sql, bindings := this.CreateSql(fields)
	result, err := this.getExecutor().Exec(sql, bindings...)
	if err != nil {
		panic(CreateException{Exceptions.WithError(err, fields)})
	}
	var id, lastIdErr = result.LastInsertId()
	if lastIdErr != nil {
		Logs.WithError(lastIdErr).Debug("Table.Create: get last insert id failed")
	}

	if _, existsPrimaryKey := fields[this.primaryKey]; !existsPrimaryKey {
		fields[this.primaryKey] = id
	}

	if this.class != nil {
		return this.class.New(fields)
	}

	return fields
}

func (this *Table) Insert(values ...Support.Fields) bool {
	sql, bindings := this.InsertSql(values)
	result, err := this.getExecutor().Exec(sql, bindings...)

	if err != nil {
		panic(InsertException{Exceptions.WithError(err, Support.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		panic(InsertException{Exceptions.WithError(err, Support.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	return rowsAffected > 0
}

func (this *Table) InsertGetId(values ...Support.Fields) int64 {
	sql, bindings := this.InsertSql(values)
	result, err := this.getExecutor().Exec(sql, bindings...)

	if err != nil {
		panic(InsertException{Exceptions.WithError(err, Support.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	id, err := result.LastInsertId()

	if err != nil {
		panic(InsertException{Exceptions.WithError(err, Support.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	return id
}

func (this *Table) InsertOrIgnore(values ...Support.Fields) int64 {
	sql, bindings := this.InsertIgnoreSql(values)
	result, err := this.getExecutor().Exec(sql, bindings...)

	if err != nil {
		panic(InsertException{Exceptions.WithError(err, Support.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		panic(InsertException{Exceptions.WithError(err, Support.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	return rowsAffected
}

func (this *Table) InsertOrReplace(values ...Support.Fields) int64 {
	sql, bindings := this.InsertReplaceSql(values)
	result, err := this.getExecutor().Exec(sql, bindings...)

	if err != nil {
		panic(InsertException{Exceptions.WithError(err, Support.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		panic(InsertException{Exceptions.WithError(err, Support.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	return rowsAffected
}

func (this *Table) FirstOrCreate(values ...Support.Fields) interface{} {
	var attributes Support.Fields
	argsLen := len(values)
	if argsLen > 0 {
		for field, value := range values[0] {
			attributes[field] = value
			this.Where(field, value)
		}
		if result := this.First(); result != nil {
			return result
		} else if argsLen > 1 {
			Field.MergeFields(attributes, values[1])
		}
		return this.Create(attributes)
	}

	return nil
}
