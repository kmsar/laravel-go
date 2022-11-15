package querybuilder

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
)

func (b *Builder) WhereIsNull(field string, whereType ...IDatabase.WhereJoinType) IDatabase.QueryBuilder {
	if len(whereType) == 0 {
		return b.Where(field, "is", "null")
	}
	return b.Where(field, "is", "null", whereType[0])
}

func (b *Builder) WhereNotNull(field string, whereType ...IDatabase.WhereJoinType) IDatabase.QueryBuilder {
	if len(whereType) == 0 {
		return b.Where(field, "is not", "null")
	}
	return b.Where(field, "is not", "null", whereType[0])
}

func (b *Builder) OrWhereIsNull(field string) IDatabase.QueryBuilder {
	return b.OrWhere(field, "is", "null")
}

func (b *Builder) OrWhereNotNull(field string) IDatabase.QueryBuilder {
	return b.OrWhere(field, "is not", "null")
}
