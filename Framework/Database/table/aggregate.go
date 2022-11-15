package table

import (
	"database/sql"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Exceptions"
)

func (this *Table) Count(columns ...string) int64 {
	queryStatement, bindings := this.WithCount(columns...).SelectSql()
	var num int64
	err := this.getExecutor().Get(&num, queryStatement, bindings...)
	if err != nil && err != sql.ErrNoRows {
		Exceptions.Throw(SelectException{Exceptions.WithError(err, Support.Fields{
			"columns":  columns,
			"sql":      queryStatement,
			"bindings": bindings,
		})})
	}
	return num
}

func (this *Table) Avg(column string, as ...string) int64 {
	queryStatement, bindings := this.WithAvg(column, as...).SelectSql()
	var num int64
	err := this.getExecutor().Get(&num, queryStatement, bindings...)
	if err != nil && err != sql.ErrNoRows {
		Exceptions.Throw(SelectException{Exceptions.WithError(err, Support.Fields{
			"column":   column,
			"sql":      queryStatement,
			"bindings": bindings,
		})})
	}
	return num
}

func (this *Table) Sum(column string, as ...string) int64 {
	queryStatement, bindings := this.WithSum(column, as...).SelectSql()
	var num int64
	err := this.getExecutor().Get(&num, queryStatement, bindings...)
	if err != nil && err != sql.ErrNoRows {
		Exceptions.Throw(SelectException{Exceptions.WithError(err, Support.Fields{
			"column":   column,
			"sql":      queryStatement,
			"bindings": bindings,
		})})
	}
	return num
}

func (this *Table) Max(column string, as ...string) int64 {
	queryStatement, bindings := this.WithMax(column, as...).SelectSql()
	var num int64
	err := this.getExecutor().Get(&num, queryStatement, bindings...)
	if err != nil && err != sql.ErrNoRows {
		Exceptions.Throw(SelectException{Exceptions.WithError(err, Support.Fields{
			"column":   column,
			"sql":      queryStatement,
			"bindings": bindings,
		})})
	}
	return num
}

func (this *Table) Min(column string, as ...string) int64 {
	queryStatement, bindings := this.WithMin(column, as...).SelectSql()
	var num int64
	err := this.getExecutor().Get(&num, queryStatement, bindings...)
	if err != nil && err != sql.ErrNoRows {
		Exceptions.Throw(SelectException{Exceptions.WithError(err, Support.Fields{
			"column":   column,
			"sql":      queryStatement,
			"bindings": bindings,
		})})
	}
	return num
}
