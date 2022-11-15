package querybuilder

import "github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"

func (b *Builder) WhereIn(field string, args interface{}, joinType ...IDatabase.WhereJoinType) IDatabase.QueryBuilder {
	if len(joinType) == 0 {
		return b.Where(field, "in", args)
	}
	return b.Where(field, "in", args, joinType[0])
}

func (b *Builder) OrWhereIn(field string, args interface{}) IDatabase.QueryBuilder {
	return b.OrWhere(field, "in", args)
}

func (b *Builder) WhereNotIn(field string, args interface{}, joinType ...IDatabase.WhereJoinType) IDatabase.QueryBuilder {
	if len(joinType) == 0 {
		return b.Where(field, "not in", args)
	}
	return b.Where(field, "not in", args)
}

func (b *Builder) OrWhereNotIn(field string, args interface{}) IDatabase.QueryBuilder {
	return b.OrWhere(field, "not in", args)
}
