package table

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Exceptions"
	"github.com/kmsar/laravel-go/Framework/Support/Field"
)

func (this *Table) UpdateOrInsert(attributes Support.Fields, values ...Support.Fields) bool {
	this.WhereFields(attributes)
	sql, bindings := this.UpdateSql(attributes)
	result, err := this.getExecutor().Exec(sql, bindings...)
	if err != nil {
		panic(UpdateException{Exceptions.WithError(err, Support.Fields{
			"attributes": attributes,
			"values":     values,
		})})
	}
	num, _ := result.RowsAffected()
	if num > 0 {
		return true
	}
	if len(values) > 0 {
		Field.MergeFields(attributes, values[0])
	}
	return this.Insert(attributes)
}

func (this *Table) UpdateOrCreate(attributes Support.Fields, values ...Support.Fields) interface{} {
	this.WhereFields(attributes)
	sql, bindings := this.UpdateSql(attributes)
	result, err := this.getExecutor().Exec(sql, bindings...)
	if err != nil {
		panic(UpdateException{Exceptions.WithError(err, Support.Fields{
			"attributes": attributes,
			"values":     values,
		})})
	}
	num, _ := result.RowsAffected()
	if num > 0 {
		return true
	}
	if len(values) > 0 {
		Field.MergeFields(attributes, values[0])
	}
	return this.Insert(attributes)
}

func (this *Table) Update(fields Support.Fields) int64 {
	sql, bindings := this.UpdateSql(fields)
	result, err := this.getExecutor().Exec(sql, bindings...)
	if err != nil {
		panic(UpdateException{Exceptions.WithError(err, fields)})
	}
	num, err := result.RowsAffected()
	if err != nil {
		panic(UpdateException{Exceptions.WithError(err, fields)})
	}
	return num
}
