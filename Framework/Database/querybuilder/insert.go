package querybuilder

import (
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Str"
	"github.com/kmsar/laravel-go/Framework/Support/Utils"
	"strings"
)

func getInsertType(insertType2 ...IDatabase.InsertType) IDatabase.InsertType {
	if len(insertType2) > 0 {
		return insertType2[0]
	}
	return IDatabase.Insert
}

func (b *Builder) CreateSql(value Support.Fields, insertType2 ...IDatabase.InsertType) (sql string, bindings []interface{}) {
	if len(value) == 0 {
		return
	}
	keys := make([]string, 0)

	valuesString := fmt.Sprintf("(%s)", strings.Join(Str.MakeSymbolArray("?", len(value)), ","))
	for name, field := range value {
		bindings = append(bindings, field)
		keys = append(keys, name)
	}

	sql = fmt.Sprintf("%s into %s %s values %s", getInsertType(insertType2...), b.table, fmt.Sprintf("(%s)", strings.Join(keys, ",")), valuesString)
	return
}

func (b *Builder) InsertSql(values []Support.Fields, insertType2 ...IDatabase.InsertType) (sql string, bindings []interface{}) {
	if len(values) == 0 {
		return
	}
	fields := Utils.GetMapKeys(values[0])
	valuesString := make([]string, 0)

	for _, value := range values {
		valuesString = append(valuesString, fmt.Sprintf("(%s)", strings.Join(Str.MakeSymbolArray("?", len(value)), ",")))
		for _, field := range fields {
			bindings = append(bindings, value[field])
		}
	}

	fieldsString := fmt.Sprintf(" (%s)", strings.Join(fields, ","))

	sql = fmt.Sprintf("%s into %s%s values %s", getInsertType(insertType2...), b.table, fieldsString, strings.Join(valuesString, ","))
	return
}

func (b *Builder) InsertIgnoreSql(values []Support.Fields) (sql string, bindings []interface{}) {
	return b.InsertSql(values, IDatabase.InsertIgnore)
}

func (b *Builder) InsertReplaceSql(values []Support.Fields) (sql string, bindings []interface{}) {
	return b.InsertSql(values, IDatabase.InsertReplace)
}
