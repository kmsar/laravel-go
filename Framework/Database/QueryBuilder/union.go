package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

type Unions map[contracts.UnionJoinType][]contracts.QueryBuilder

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

func (b *Builder) Union(builder contracts.QueryBuilder, unionType ...contracts.UnionJoinType) contracts.QueryBuilder {
	if builder != nil {
		if len(unionType) > 0 {
			b.unions[unionType[0]] = append(b.unions[unionType[0]], builder)
		} else {
			b.unions[contracts.Union] = append(b.unions[contracts.Union], builder)
		}
	}

	return b.addBinding(unionBinding, builder.GetBindings()...)
}

func (b *Builder) UnionAll(builder contracts.QueryBuilder) contracts.QueryBuilder {
	return b.Union(builder, contracts.UnionAll)
}

func (b *Builder) UnionByProvider(builder contracts.QueryProvider, unionType ...contracts.UnionJoinType) contracts.QueryBuilder {
	return b.Union(builder(), unionType...)
}

func (b *Builder) UnionAllByProvider(builder contracts.QueryProvider) contracts.QueryBuilder {
	return b.Union(builder(), contracts.UnionAll)
}
