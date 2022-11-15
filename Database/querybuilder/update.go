package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
	"strings"
)

type Expression string

func (b *Builder) UpdateSql(value contracts.Fields) (sql string, bindings []interface{}) {
	if len(value) == 0 {
		return
	}
	valuesString := make([]string, 0)
	for name, field := range value {
		if expression, isExpression := field.(Expression); isExpression {
			valuesString = append(valuesString, fmt.Sprintf("%s = %s", name, expression))
		} else {
			valuesString = append(valuesString, fmt.Sprintf("%s = ?", name))
			bindings = append(bindings, field)
		}
	}

	sql = fmt.Sprintf("update %s set %s", b.table, strings.Join(valuesString, ","))

	if !b.wheres.IsEmpty() {
		sql = fmt.Sprintf("%s where %s", sql, b.wheres.String())
	}

	bindings = append(bindings, b.bindings[whereBinding]...)

	return
}
