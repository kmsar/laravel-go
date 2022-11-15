package querybuilder

import (
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
)

func (b *Builder) WhereExists(provider IDatabase.QueryProvider, where ...IDatabase.WhereJoinType) IDatabase.QueryBuilder {
	subBuilder := provider()
	subSql := fmt.Sprintf("(%s)", subBuilder.ToSql())
	if len(where) == 0 {
		return b.addBinding(whereBinding, subBuilder.GetBindings()...).
			Where("", "exists", subSql)
	}

	return b.addBinding(whereBinding, subBuilder.GetBindings()...).
		Where("", "exists", subSql, where[0])

}

func (b *Builder) OrWhereExists(provider IDatabase.QueryProvider) IDatabase.QueryBuilder {
	return b.WhereExists(provider, IDatabase.Or)
}

func (b *Builder) WhereNotExists(provider IDatabase.QueryProvider, where ...IDatabase.WhereJoinType) IDatabase.QueryBuilder {
	subBuilder := provider()
	subSql := fmt.Sprintf("(%s)", subBuilder.ToSql())
	if len(where) == 0 {
		return b.addBinding(whereBinding, subBuilder.GetBindings()...).
			Where("", "not exists", subSql)
	}

	return b.addBinding(whereBinding, subBuilder.GetBindings()...).
		Where("", "not exists", subSql, where[0])
}

func (b *Builder) OrWhereNotExists(provider IDatabase.QueryProvider) IDatabase.QueryBuilder {
	return b.WhereNotExists(provider, IDatabase.Or)
}
