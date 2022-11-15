package querybuilder

import "github.com/goal-web/contracts"

func (b *Builder) WhereIsNull(field string, whereType ...contracts.WhereJoinType) contracts.QueryBuilder {
	if len(whereType) == 0 {
		return b.Where(field, "is", "null")
	}
	return b.Where(field, "is", "null", whereType[0])
}

func (b *Builder) WhereNotNull(field string, whereType ...contracts.WhereJoinType) contracts.QueryBuilder {
	if len(whereType) == 0 {
		return b.Where(field, "is not", "null")
	}
	return b.Where(field, "is not", "null", whereType[0])
}

func (b *Builder) OrWhereIsNull(field string) contracts.QueryBuilder {
	return b.OrWhere(field, "is", "null")
}

func (b *Builder) OrWhereNotNull(field string) contracts.QueryBuilder {
	return b.OrWhere(field, "is not", "null")
}
