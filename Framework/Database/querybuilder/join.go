package querybuilder

import (
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
)

type Join struct {
	table      string
	join       IDatabase.JoinType
	conditions *Wheres
}

func (this Join) String() (result string) {
	result = fmt.Sprintf("%s join %s", this.join, this.table)
	if this.conditions.IsEmpty() {
		return
	}
	result = fmt.Sprintf("%s on (%s)", result, this.conditions.String())
	return
}

type Joins []Join

func (this Joins) IsEmpty() bool {
	return len(this) == 0
}

func (this Joins) String() (result string) {
	if this.IsEmpty() {
		return
	}

	for index, join := range this {
		if index == 0 {
			result = join.String()
		} else {
			result = fmt.Sprintf("%s %s", result, join.String())
		}
	}

	return
}

func (b *Builder) Join(table string, first, condition, second string, joins ...IDatabase.JoinType) IDatabase.QueryBuilder {
	join := IDatabase.InnerJoin
	if len(joins) > 0 {
		join = joins[0]
	}
	b.joins = append(b.joins, Join{table, join, &Wheres{wheres: map[IDatabase.WhereJoinType][]*Where{
		IDatabase.And: {&Where{
			field:     first,
			condition: condition,
			arg:       second,
		}},
	}}})

	return b
}

func (b *Builder) JoinSub(provider IDatabase.QueryProvider, as, first, condition, second string, joins ...IDatabase.JoinType) IDatabase.QueryBuilder {
	join := IDatabase.InnerJoin
	if len(joins) > 0 {
		join = joins[0]
	}
	subBuilder := provider()
	b.joins = append(b.joins, Join{fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as), join, &Wheres{wheres: map[IDatabase.WhereJoinType][]*Where{
		IDatabase.And: {&Where{
			field:     first,
			condition: condition,
			arg:       second,
		}},
	}}})

	return b.addBinding(joinBinding, subBuilder.GetBindings()...)
}

func (b *Builder) FullJoin(table string, first, condition, second string) IDatabase.QueryBuilder {
	return b.Join(table, first, condition, second, IDatabase.FullJoin)
}
func (b *Builder) FullOutJoin(table string, first, condition, second string) IDatabase.QueryBuilder {
	return b.Join(table, first, condition, second, IDatabase.FullOutJoin)
}

func (b *Builder) LeftJoin(table string, first, condition, second string) IDatabase.QueryBuilder {
	return b.Join(table, first, condition, second, IDatabase.LeftJoin)
}

func (b *Builder) RightJoin(table string, first, condition, second string) IDatabase.QueryBuilder {
	return b.Join(table, first, condition, second, IDatabase.RightJoin)
}
