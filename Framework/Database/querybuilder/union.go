package querybuilder

import (
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
)

type Unions map[IDatabase.UnionJoinType][]IDatabase.QueryBuilder

func (this Unions) IsEmpty() bool {
	return len(this) == 0
}

func (this Unions) String() (result string) {
	if this.IsEmpty() {
		return
	}
	for unionType, builders := range this {
		for _, builder := range builders {
			result = fmt.Sprintf("%s %s (%s)", result, unionType, builder.ToSql())
		}
	}

	return
}

func (b *Builder) Union(builder IDatabase.QueryBuilder, unionType ...IDatabase.UnionJoinType) IDatabase.QueryBuilder {
	if builder != nil {
		if len(unionType) > 0 {
			b.unions[unionType[0]] = append(b.unions[unionType[0]], builder)
		} else {
			b.unions[IDatabase.Union] = append(b.unions[IDatabase.Union], builder)
		}
	}

	return b.addBinding(unionBinding, builder.GetBindings()...)
}

func (b *Builder) UnionAll(builder IDatabase.QueryBuilder) IDatabase.QueryBuilder {
	return b.Union(builder, IDatabase.UnionAll)
}

func (b *Builder) UnionByProvider(builder IDatabase.QueryProvider, unionType ...IDatabase.UnionJoinType) IDatabase.QueryBuilder {
	return b.Union(builder(), unionType...)
}

func (b *Builder) UnionAllByProvider(builder IDatabase.QueryProvider) IDatabase.QueryBuilder {
	return b.Union(builder(), IDatabase.UnionAll)
}
