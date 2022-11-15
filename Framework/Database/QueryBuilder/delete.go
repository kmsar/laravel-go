package querybuilder

import (
	"fmt"
)

func (b *Builder) DeleteSql() (sql string, bindings []interface{}) {
	sql = fmt.Sprintf("delete from %s", b.table)

	if !b.wheres.IsEmpty() {
		sql = fmt.Sprintf("%s where %s", sql, b.wheres.String())
	}
	bindings = b.GetBindings()
	return
}
