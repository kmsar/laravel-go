package querybuilder

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"strings"
)

type GroupBy []string

func (b GroupBy) IsEmpty() bool {
	return len(b) == 0
}

func (b GroupBy) String() string {
	if b.IsEmpty() {
		return ""
	}

	return strings.Join(b, ",")
}

func (b *Builder) GroupBy(columns ...string) IDatabase.QueryBuilder {
	b.groupBy = append(b.groupBy, columns...)

	return b
}

func (b *Builder) Having(field string, args ...interface{}) IDatabase.QueryBuilder {
	var (
		arg       interface{}
		condition = "="
		whereType = IDatabase.And
	)
	switch len(args) {
	case 1:
		arg = args[0]
	case 2:
		condition = args[0].(string)
		arg = args[1]
	case 3:
		condition = args[0].(string)
		arg = args[1]
		whereType = args[2].(IDatabase.WhereJoinType)
	}

	raw, bindings := b.prepareArgs(condition, arg)

	b.having.wheres[whereType] = append(b.having.wheres[whereType], &Where{
		field:     field,
		condition: condition,
		arg:       raw,
	})

	return b.addBinding(havingBinding, bindings...)
}

func (b *Builder) OrHaving(field string, args ...interface{}) IDatabase.QueryBuilder {
	var (
		arg       interface{}
		condition = "="
	)
	switch len(args) {
	case 1:
		arg = args[0]
	case 2:
		condition = args[0].(string)
		arg = args[1]
	default:
		condition = args[0].(string)
		arg = args[1]
	}
	raw, bindings := b.prepareArgs(condition, arg)

	b.having.wheres[IDatabase.Or] = append(b.having.wheres[IDatabase.Or], &Where{
		field:     field,
		condition: condition,
		arg:       raw,
	})
	return b.addBinding(havingBinding, bindings...)
}
