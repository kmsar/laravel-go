package querybuilder

import "github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"

// WhereBetween args 参数可以是整数、浮点数、字符串、interface{} 等类型的数组，或者用` and `隔开的字符串，或者在源码中了解更多 https://github.com/goal-web/querybuilder/blob/78bcc832604bfcdb68579e3dd1441796a16994cf/builder.go#L74
func (b *Builder) WhereBetween(field string, args interface{}, whereType ...IDatabase.WhereJoinType) IDatabase.QueryBuilder {
	if len(whereType) > 0 {
		return b.Where(field, "between", args, whereType[0])
	}

	return b.Where(field, "between", args)
}

func (b *Builder) OrWhereBetween(field string, args interface{}) IDatabase.QueryBuilder {
	return b.OrWhere(field, "between", args)
}

func (b *Builder) WhereNotBetween(field string, args interface{}, whereType ...IDatabase.WhereJoinType) IDatabase.QueryBuilder {
	if len(whereType) > 0 {
		return b.Where(field, "not between", args, whereType[0])
	}

	return b.Where(field, "not between", args)
}

func (b *Builder) OrWhereNotBetween(field string, args interface{}) IDatabase.QueryBuilder {
	return b.OrWhere(field, "not between", args)
}
