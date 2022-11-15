package querybuilder

import (
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

type Where struct {
	field     string
	condition string
	arg       string
}

func (this *Where) String() string {
	if this == nil {
		return ""
	}
	return fmt.Sprintf("%s %s %s", this.field, this.condition, this.arg)
}

type Wheres struct {
	subWheres map[IDatabase.WhereJoinType][]*Wheres
	wheres    map[IDatabase.WhereJoinType][]*Where
}

func (this *Wheres) IsEmpty() bool {
	return len(this.subWheres) == 0 && len(this.wheres) == 0
}

func (this Wheres) getSubStringers(whereType IDatabase.WhereJoinType) []fmt.Stringer {
	stringers := make([]fmt.Stringer, 0)
	for _, where := range this.subWheres[whereType] {
		stringers = append(stringers, where)
	}
	return stringers
}

func (this Wheres) getStringers(whereType IDatabase.WhereJoinType) []fmt.Stringer {
	stringers := make([]fmt.Stringer, 0)
	for _, where := range this.wheres[whereType] {
		stringers = append(stringers, where)
	}
	return stringers
}

func (this *Wheres) getSubWheres(whereType IDatabase.WhereJoinType) string {
	return JoinSubStringerArray(this.getSubStringers(whereType), string(whereType))
}

func (this *Wheres) getWheres(whereType IDatabase.WhereJoinType) string {
	return JoinStringerArray(this.getStringers(whereType), string(whereType))
}

func (this *Wheres) String() (result string) {
	if this == nil || this.IsEmpty() {
		return ""
	}

	result = this.getSubWheres(IDatabase.And)
	andWheres := this.getWheres(IDatabase.And)

	if result != "" {
		if andWheres != "" {
			result = fmt.Sprintf("%s and %s", result, andWheres)
		}
	} else {
		result = andWheres
	}

	orSubWheres := this.getSubWheres(IDatabase.Or)
	if result == "" {
		result = orSubWheres
	} else if orSubWheres != "" {
		result = fmt.Sprintf("%s or %s", result, orSubWheres)
	}

	orWheres := this.getWheres(IDatabase.Or)
	if result == "" {
		result = orWheres
	} else if orWheres != "" {
		result = fmt.Sprintf("%s or %s", result, orWheres)
	}

	return
}

func (b *Builder) WhereFunc(callback IDatabase.QueryFunc, whereType ...IDatabase.WhereJoinType) IDatabase.QueryBuilder {
	subBuilder := &Builder{
		wheres: &Wheres{
			wheres:    map[IDatabase.WhereJoinType][]*Where{},
			subWheres: map[IDatabase.WhereJoinType][]*Wheres{},
		},
		bindings: map[bindingType][]interface{}{},
	}
	callback(subBuilder)
	if len(whereType) == 0 {
		b.wheres.subWheres[IDatabase.And] = append(b.wheres.subWheres[IDatabase.And], subBuilder.getWheres())
	} else {
		b.wheres.subWheres[whereType[0]] = append(b.wheres.subWheres[whereType[0]], subBuilder.getWheres())
	}
	return b.addBinding(whereBinding, subBuilder.GetBindings()...)
}

func (b *Builder) WhereFields(fields Support.Fields) IDatabase.QueryBuilder {
	for column, value := range fields {
		b.Where(column, value)
	}
	return b
}

func (b *Builder) OrWhereFunc(callback IDatabase.QueryFunc) IDatabase.QueryBuilder {
	return b.WhereFunc(callback, IDatabase.Or)
}

func (b *Builder) Where(field string, args ...interface{}) IDatabase.QueryBuilder {
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

	b.wheres.wheres[whereType] = append(b.wheres.wheres[whereType], &Where{
		field:     field,
		condition: condition,
		arg:       raw,
	})

	return b.addBinding(whereBinding, bindings...)
}

func (b *Builder) OrWhere(field string, args ...interface{}) IDatabase.QueryBuilder {
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

	b.wheres.wheres[IDatabase.Or] = append(b.wheres.wheres[IDatabase.Or], &Where{
		field:     field,
		condition: condition,
		arg:       raw,
	})
	return b.addBinding(whereBinding, bindings...)
}

func JoinStringerArray(arr []fmt.Stringer, sep string) (result string) {
	for index, stringer := range arr {
		if index == 0 {
			result = stringer.String()
		} else {
			result = fmt.Sprintf("%s %s %s", result, sep, stringer.String())
		}
	}

	return
}

func JoinSubStringerArray(arr []fmt.Stringer, sep string) (result string) {
	for index, stringer := range arr {
		if index == 0 {
			result = fmt.Sprintf("(%s)", stringer.String())
		} else {
			result = fmt.Sprintf("%s %s (%s)", result, sep, stringer.String())
		}
	}

	return
}
