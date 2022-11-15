package querybuilder

import (
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
)

func (b *Builder) Select(fields ...string) IDatabase.QueryBuilder {
	b.fields = fields
	return b
}

func (b *Builder) AddSelect(fields ...string) IDatabase.QueryBuilder {
	b.fields = append(b.fields, fields...)
	return b
}

func (b *Builder) SelectSub(provider IDatabase.QueryProvider, as string) IDatabase.QueryBuilder {
	subBuilder := provider()
	b.fields = []string{fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as)}
	return b.addBinding(selectBinding, subBuilder.GetBindings()...)
}
func (b *Builder) AddSelectSub(provider IDatabase.QueryProvider, as string) IDatabase.QueryBuilder {
	subBuilder := provider()
	b.fields = append(b.fields, fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as))
	return b.addBinding(selectBinding, subBuilder.GetBindings()...)
}
