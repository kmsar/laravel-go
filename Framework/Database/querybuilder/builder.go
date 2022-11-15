package querybuilder

import (
	"errors"
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Str"
	"github.com/kmsar/laravel-go/Framework/Support/Utils/Convert"
	"strings"
)

type bindingType string
type Builder struct {
	IDatabase.QueryBuilder
	limit    int64
	offset   int64
	distinct bool
	table    string
	fields   []string
	wheres   *Wheres
	orderBy  OrderByFields
	groupBy  GroupBy
	joins    Joins
	unions   Unions
	having   *Wheres
	bindings map[bindingType][]interface{}
}

func (b *Builder) Bind(builder IDatabase.QueryBuilder) IDatabase.QueryBuilder {
	b.QueryBuilder = builder
	return b
}

func (b *Builder) Skip(offset int64) IDatabase.QueryBuilder {
	return b.Offset(offset)
}

func (b *Builder) Take(num int64) IDatabase.QueryBuilder {
	return b.Limit(num)
}

const (
	selectBinding  bindingType = "select"
	fromBinding    bindingType = "from"
	joinBinding    bindingType = "join"
	whereBinding   bindingType = "where"
	groupByBinding bindingType = "groupBy"
	havingBinding  bindingType = "having"
	orderBinding   bindingType = "order"
	unionBinding   bindingType = "union"
)

func NewQuery(table string) *Builder {
	return &Builder{
		table:    table,
		fields:   []string{"*"},
		orderBy:  OrderByFields{},
		bindings: map[bindingType][]interface{}{},
		joins:    Joins{},
		unions:   Unions{},
		groupBy:  GroupBy{},
		wheres: &Wheres{
			wheres:    map[IDatabase.WhereJoinType][]*Where{},
			subWheres: map[IDatabase.WhereJoinType][]*Wheres{},
		},
		having: &Wheres{
			wheres:    map[IDatabase.WhereJoinType][]*Where{},
			subWheres: map[IDatabase.WhereJoinType][]*Wheres{},
		},
	}
}

func FromSub(callback IDatabase.QueryProvider, as string) IDatabase.QueryBuilder {
	return NewQuery("").FromSub(callback, as)
}

func (b *Builder) getWheres() *Wheres {
	return b.wheres
}

func (b *Builder) prepareArgs(condition string, args interface{}) (raw string, bindings []interface{}) {
	if expression, isExpression := args.(Expression); isExpression {
		return string(expression), bindings
	} else if builder, isBuilder := args.(IDatabase.QueryBuilder); isBuilder {
		raw, bindings = builder.SelectSql()
		raw = fmt.Sprintf("(%s)", raw)
		return
	}
	condition = strings.ToLower(condition)
	switch condition {
	case "in", "not in", "between", "not between":
		isInGrammar := strings.Contains(condition, "in")
		joinSymbol := Str.IfString(isInGrammar, ",", " and ")
		var stringArg string
		switch arg := args.(type) {
		case string:
			stringArg = arg
		case fmt.Stringer:
			stringArg = arg.String()
		case []string:
			stringArg = strings.Join(arg, joinSymbol)
		case []int:
			stringArg = Str.JoinIntArray(arg, joinSymbol)
		case []int64:
			stringArg = Str.JoinInt64Array(arg, joinSymbol)
		case []float64:
			stringArg = Str.JoinFloat64Array(arg, joinSymbol)
		case []float32:
			stringArg = Str.JoinFloatArray(arg, joinSymbol)
		case []interface{}:
			bindings = arg
			raw = fmt.Sprintf("(%s)", strings.Join(Str.MakeSymbolArray("?", len(bindings)), ","))
			return
		default:
			panic(ParamException{errors.New("不支持的参数类型"), Support.Fields{
				"arg":       arg,
				"condition": condition,
			}})
		}
		bindings = Str.StringArray2InterfaceArray(strings.Split(stringArg, joinSymbol))
		if isInGrammar {
			raw = fmt.Sprintf("(%s)", strings.Join(Str.MakeSymbolArray("?", len(bindings)), ","))
		} else {
			raw = "? and ?"
		}
	case "is", "is not", "exists", "not exists":
		raw = Convert.ConvertToString(args, "")
	default:
		raw = "?"
		bindings = append(bindings, Convert.ConvertToString(args, ""))
	}

	return
}

func (b *Builder) addBinding(bindType bindingType, bindings ...interface{}) IDatabase.QueryBuilder {
	b.bindings[bindType] = append(b.bindings[bindType], bindings...)
	return b
}

func (b *Builder) GetBindings() (results []interface{}) {
	for _, binding := range []bindingType{
		selectBinding, fromBinding, joinBinding,
		whereBinding, groupByBinding, havingBinding, orderBinding, unionBinding,
	} {
		results = append(results, b.bindings[binding]...)
	}
	return
}

func (b *Builder) Distinct() IDatabase.QueryBuilder {
	b.distinct = true
	return b
}

func (b *Builder) From(table string, as ...string) IDatabase.QueryBuilder {
	if len(as) == 0 {
		b.table = table
	} else {
		b.table = fmt.Sprintf("%s as %s", table, as[0])
	}
	return b
}

func (b *Builder) Offset(offset int64) IDatabase.QueryBuilder {
	b.offset = offset
	return b
}

func (b *Builder) Limit(num int64) IDatabase.QueryBuilder {
	b.limit = num
	return b
}

func (b *Builder) WithPagination(perPage int64, current ...int64) IDatabase.QueryBuilder {
	b.limit = perPage
	if len(current) > 0 {
		b.offset = perPage * (current[0] - 1)
	}
	return b
}

func (b *Builder) FromMany(tables ...string) IDatabase.QueryBuilder {
	if len(tables) > 0 {
		b.table = strings.Join(tables, ",")
	}
	return b
}

func (b *Builder) FromSub(provider IDatabase.QueryProvider, as string) IDatabase.QueryBuilder {
	subBuilder := provider()
	b.table = fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as)
	return b.addBinding(fromBinding, subBuilder.GetBindings()...)
}

func (b *Builder) When(condition bool, callback IDatabase.QueryCallback, elseCallback ...IDatabase.QueryCallback) IDatabase.QueryBuilder {
	if condition {
		return callback(b)
	} else if len(elseCallback) > 0 {
		return elseCallback[0](b)
	}
	return b
}

func (b *Builder) getSelect() string {
	if b.distinct {
		return "distinct " + strings.Join(b.fields, ",")
	}
	return strings.Join(b.fields, ",")
}

func (b *Builder) ToSql() string {
	sql := fmt.Sprintf("select %s from %s", b.getSelect(), b.table)

	if !b.joins.IsEmpty() {
		sql = fmt.Sprintf("%s %s", sql, b.joins.String())
	}

	if !b.wheres.IsEmpty() {
		sql = fmt.Sprintf("%s where %s", sql, b.wheres.String())
	}

	if !b.groupBy.IsEmpty() {
		sql = fmt.Sprintf("%s group by %s", sql, b.groupBy.String())

		if !b.having.IsEmpty() {
			sql = fmt.Sprintf("%s having %s", sql, b.having.String())
		}
	}

	if !b.orderBy.IsEmpty() {
		sql = fmt.Sprintf("%s order by %s", sql, b.orderBy.String())
	}

	if b.limit > 0 {
		sql = fmt.Sprintf("%s limit %d", sql, b.limit)
	}
	if b.offset > 0 {
		sql = fmt.Sprintf("%s offset %d", sql, b.offset)
	}

	if !b.unions.IsEmpty() {
		sql = fmt.Sprintf("(%s) %s", sql, b.unions.String())
	}

	return sql
}

func (b *Builder) SelectSql() (string, []interface{}) {
	return b.ToSql(), b.GetBindings()
}

func (b *Builder) SelectForUpdateSql() (string, []interface{}) {
	return b.ToSql() + " for update", b.GetBindings()
}
