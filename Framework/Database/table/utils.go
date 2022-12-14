package table

import (
	"database/sql"

	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

func ParseRows(rows *sql.Rows) ([]Support.Fields, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	columnLength := len(columns)
	cache := make([]interface{}, columnLength)
	for index, _ := range cache {
		var a interface{}
		cache[index] = &a
	}
	var results []Support.Fields
	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{})
		}
		results = append(results, item)
	}
	_ = rows.Close()
	return results, nil
}
