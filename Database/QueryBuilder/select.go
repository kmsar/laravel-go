package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func (b *Builder) Select(fields ...string) contracts.QueryBuilder {
	b.fields = fields
	return b
}

func (b *Builder) AddSelect(fields ...string) contracts.QueryBuilder {
	b.fields = append(b.fields, fields...)
	return b
}

func (b *Builder) SelectSub(provider contracts.QueryProvider, as string) contracts.QueryBuilder {
	subBuilder := provider()
	b.fields = []string{fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as)}
	return b.addBinding(selectBinding, subBuilder.GetBindings()...)
}
func (b *Builder) AddSelectSub(provider contracts.QueryProvider, as string) contracts.QueryBuilder {
	subBuilder := provider()
	b.fields = append(b.fields, fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as))
	return b.addBinding(selectBinding, subBuilder.GetBindings()...)
}
