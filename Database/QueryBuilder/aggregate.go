package querybuilder

import (
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
)

func (b *Builder) WithCount(fields ...string) IDatabase.QueryBuilder {
	if len(fields) == 0 {
		return b.Select("count(*)")
	}
	return b.Select(fmt.Sprintf("count(%s) as %s_count", fields[0], fields[0]))
}

func (b *Builder) WithAvg(field string, as ...string) IDatabase.QueryBuilder {
	if len(as) == 0 {
		return b.Select(fmt.Sprintf("avg(%s) as %s_avg", field, field))
	}
	return b.Select(fmt.Sprintf("avg(%s) as %s", field, as[0]))
}

func (b *Builder) WithSum(field string, as ...string) IDatabase.QueryBuilder {
	if len(as) == 0 {
		return b.Select(fmt.Sprintf("sum(%s) as %s_sum", field, field))
	}
	return b.Select(fmt.Sprintf("sum(%s) as %s", field, as[0]))
}

func (b *Builder) WithMax(field string, as ...string) IDatabase.QueryBuilder {
	if len(as) == 0 {
		return b.Select(fmt.Sprintf("max(%s) as %s_max", field, field))
	}
	return b.Select(fmt.Sprintf("max(%s) as %s", field, as[0]))
}

func (b *Builder) WithMin(field string, as ...string) IDatabase.QueryBuilder {
	if len(as) == 0 {
		return b.Select(fmt.Sprintf("min(%s) as %s_min", field, field))
	}
	return b.Select(fmt.Sprintf("min(%s) as %s", field, as[0]))
}
