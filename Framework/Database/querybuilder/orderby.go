package querybuilder

import (
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
	"strings"
)

const RandomOrder IDatabase.OrderType = "RANDOM()"
const RandOrder IDatabase.OrderType = "RAND()"

type OrderBy struct {
	field          string
	fieldOrderType IDatabase.OrderType
}

type OrderByFields []OrderBy

func (this OrderByFields) IsEmpty() bool {
	return len(this) == 0
}

func (this OrderByFields) String() string {
	if this.IsEmpty() {
		return ""
	}

	columns := make([]string, 0)

	for _, orderBy := range this {
		if orderBy.field == "" {
			columns = append(columns, fmt.Sprintf("%s", orderBy.fieldOrderType))
		} else {
			columns = append(columns, fmt.Sprintf("%s %s", orderBy.field, orderBy.fieldOrderType))
		}
	}

	return strings.Join(columns, ",")
}

func (b *Builder) OrderBy(field string, columnOrderType ...IDatabase.OrderType) IDatabase.QueryBuilder {
	if len(columnOrderType) > 0 {
		b.orderBy = append(b.orderBy, OrderBy{
			field:          field,
			fieldOrderType: columnOrderType[0],
		})
	} else {
		b.orderBy = append(b.orderBy, OrderBy{
			field:          field,
			fieldOrderType: IDatabase.Asc,
		})
	}

	return b
}

func (b *Builder) OrderByDesc(field string) IDatabase.QueryBuilder {
	b.orderBy = append(b.orderBy, OrderBy{
		field:          field,
		fieldOrderType: IDatabase.Desc,
	})
	return b
}

func (b *Builder) InRandomOrder(orderFunc ...IDatabase.OrderType) IDatabase.QueryBuilder {
	fn := RandomOrder
	if len(orderFunc) > 0 {
		fn = orderFunc[0]
	}

	b.orderBy = append(b.orderBy, OrderBy{
		field:          "",
		fieldOrderType: fn,
	})
	return b
}
