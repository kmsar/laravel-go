package querybuilder

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IContainer"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

func (b *Builder) SetExecutor(executor IDatabase.SqlExecutor) IDatabase.QueryBuilder {
	return b.QueryBuilder.SetExecutor(executor)
}

func (b *Builder) Create(fields Support.Fields) interface{} {
	return b.QueryBuilder.Create(fields)
}

func (b *Builder) Insert(values ...Support.Fields) bool {
	return b.QueryBuilder.Insert(values...)
}

func (b *Builder) Delete() int64 {
	return b.QueryBuilder.Delete()
}

func (b *Builder) Update(fields Support.Fields) int64 {
	return b.QueryBuilder.Update(fields)
}

func (b *Builder) Get() Support.Collection {
	return b.QueryBuilder.Get()
}

func (b *Builder) SelectForUpdate() Support.Collection {
	return b.QueryBuilder.SelectForUpdate()
}

func (b *Builder) Find(key interface{}) interface{} {
	return b.QueryBuilder.Find(key)
}

func (b *Builder) First() interface{} {
	return b.QueryBuilder.First()
}

func (b *Builder) InsertGetId(values ...Support.Fields) int64 {
	return b.QueryBuilder.InsertGetId(values...)
}

func (b *Builder) InsertOrIgnore(values ...Support.Fields) int64 {
	return b.QueryBuilder.InsertOrIgnore(values...)
}

func (b *Builder) InsertOrReplace(values ...Support.Fields) int64 {
	return b.QueryBuilder.InsertOrReplace(values...)
}

func (b *Builder) FirstOrCreate(values ...Support.Fields) interface{} {
	return b.QueryBuilder.FirstOrCreate(values...)
}

func (b *Builder) UpdateOrInsert(attributes Support.Fields, values ...Support.Fields) bool {
	return b.QueryBuilder.UpdateOrInsert(attributes, values...)
}

func (b *Builder) UpdateOrCreate(attributes Support.Fields, values ...Support.Fields) interface{} {
	return b.QueryBuilder.UpdateOrCreate(attributes, values...)
}

func (b *Builder) FirstOrFail() interface{} {
	return b.QueryBuilder.FirstOrFail()
}

func (b *Builder) Count(columns ...string) int64 {
	return b.QueryBuilder.Count(columns...)
}

func (b *Builder) Avg(column string, as ...string) int64 {
	return b.QueryBuilder.Avg(column, as...)
}

func (b *Builder) Sum(column string, as ...string) int64 {
	return b.QueryBuilder.Sum(column, as...)
}

func (b *Builder) Max(column string, as ...string) int64 {
	return b.QueryBuilder.Max(column, as...)
}

func (b *Builder) Min(column string, as ...string) int64 {
	return b.QueryBuilder.Min(column, as...)
}

func (b *Builder) SimplePaginate(perPage int64, current ...int64) Support.Collection {
	return b.WithPagination(perPage, current...).Get()
}

func (b *Builder) FirstOr(provider IContainer.InstanceProvider) interface{} {
	if result := b.First(); result != nil {
		return result
	}
	return provider()
}

func (b *Builder) FirstWhere(column string, args ...interface{}) interface{} {
	return b.Where(column, args...).First()
}

func (b *Builder) Paginate(perPage int64, current ...int64) (Support.Collection, int64) {
	return b.SimplePaginate(perPage, current...), b.Count()
}
