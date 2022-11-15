package IDatabase

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IContainer"
	"github.com/kmsar/laravel-go/Framework/Contracts/IEvent"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"time"
)

// DBConnector Get a database connection instance.
type DBConnector func(config Support.Fields, dispatcher IEvent.EventDispatcher) DBConnection

type Result interface {

	// LastInsertId Get the last insert ID.
	LastInsertId() (int64, error)

	// RowsAffected Get the number of affected rows.
	RowsAffected() (int64, error)
}

type DBFactory interface {

	// Connection Get the specified database connection instance.
	Connection(key ...string) DBConnection

	// Extend Register an extension database connection resolver.
	Extend(name string, driver DBConnector)
}

type DBTx interface {
	SqlExecutor

	// Commit commit the active database transaction.
	Commit() error

	// Rollback rollback the active database transaction.
	Rollback() error
}

type SqlExecutor interface {

	// Query Execute a new query against the connection.
	Query(query string, args ...interface{}) (Support.Collection, error)

	// Get Execute the query as a "select" statement.
	Get(dest interface{}, query string, args ...interface{}) error

	// Select Run a select statement against the database.
	Select(dest interface{}, query string, args ...interface{}) error

	// Exec Execute an SQL statement.
	Exec(query string, args ...interface{}) (Result, error)
}

type DBConnection interface {
	SqlExecutor

	// Begin Start a new database transaction.
	Begin() (DBTx, error)

	// Transaction Execute a Closure within a transaction.
	Transaction(func(executor SqlExecutor) error) error

	// DriverName Get the driver name.
	DriverName() string
}

// QueryCallback query callback，for building subqueries.
type QueryCallback func(QueryBuilder) QueryBuilder

// QueryProvider query provider.
type QueryProvider func() QueryBuilder

// QueryFunc Used to construct sub-where expressions.
type QueryFunc func(QueryBuilder)

type WhereJoinType string
type UnionJoinType string

const (
	Union    UnionJoinType = "union"
	UnionAll UnionJoinType = "union all"
)

type OrderType string

const (
	Desc OrderType = "desc"
	Asc  OrderType = "asc"
)

type JoinType string

const (
	LeftJoin    JoinType = "left"
	RightJoin   JoinType = "right"
	InnerJoin   JoinType = "inner"
	FullOutJoin JoinType = "full outer"
	FullJoin    JoinType = "full"
)

type InsertType string

const (
	Insert        InsertType = "insert"
	InsertIgnore  InsertType = "insert ignore"
	InsertReplace InsertType = "replace"
)

const (
	And WhereJoinType = "and"
	Or  WhereJoinType = "or"
)

type QueryBuilder interface {

	// Select Set the columns to be selected.
	Select(columns ...string) QueryBuilder

	// AddSelect Append the columns to be selected.
	AddSelect(columns ...string) QueryBuilder

	// SelectSub Add a subselect expression to the query.
	SelectSub(provider QueryProvider, as string) QueryBuilder

	// AddSelectSub Append a subselect expression to the query.
	AddSelectSub(provider QueryProvider, as string) QueryBuilder

	// WithCount Add subselect queries to count the relations.
	WithCount(columns ...string) QueryBuilder

	// WithAvg Add subselect queries to include the average of the relation's column.
	WithAvg(column string, as ...string) QueryBuilder

	// WithSum Add subselect queries to include the sum of the relation's column.
	WithSum(column string, as ...string) QueryBuilder

	// WithMax Add subselect queries to include the max of the relation's column.
	WithMax(column string, as ...string) QueryBuilder

	// WithMin Add subselect queries to include the min of the relation's column.
	WithMin(column string, as ...string) QueryBuilder

	// Count Retrieve the "count" result of the query.
	Count(columns ...string) int64

	// Avg Retrieve the average of the values of a given column.
	Avg(column string, as ...string) int64

	// Sum Retrieve the sum of the values of a given column.
	Sum(column string, as ...string) int64

	// Max Retrieve the maximum value of a given column.
	Max(column string, as ...string) int64

	// Min Retrieve the minimum value of a given column.
	Min(column string, as ...string) int64

	// Distinct Force the query to only return distinct results.
	Distinct() QueryBuilder

	// From Set the table which the query is targeting.
	From(table string, as ...string) QueryBuilder

	// FromMany Set the table that many queries are against.
	FromMany(tables ...string) QueryBuilder

	// FromSub Makes "from" fetch from a subquery.
	FromSub(provider QueryProvider, as string) QueryBuilder

	// Join Add a join clause to the query.
	Join(table string, first, condition, second string, joins ...JoinType) QueryBuilder

	// JoinSub Add a subquery join clause to the query.
	JoinSub(provider QueryProvider, as, first, condition, second string, joins ...JoinType) QueryBuilder

	// FullJoin Add a full join to the query, associate the two tables, and query all their records.
	FullJoin(table string, first, condition, second string) QueryBuilder

	// FullOutJoin Add a full outer join to the query
	FullOutJoin(table string, first, condition, second string) QueryBuilder

	// LeftJoin Add a left join to the query.
	LeftJoin(table string, first, condition, second string) QueryBuilder

	// RightJoin Add a right join to the query.
	RightJoin(table string, first, condition, second string) QueryBuilder

	// Where Add a basic where clause to the query.
	Where(column string, args ...interface{}) QueryBuilder

	// WhereFields Add an array of where clauses to the query.
	WhereFields(fields Support.Fields) QueryBuilder

	// OrWhere Add an "or where" clause to the query.
	OrWhere(column string, args ...interface{}) QueryBuilder

	// WhereFunc Add a nested where statement to the query.
	WhereFunc(callback QueryFunc, whereType ...WhereJoinType) QueryBuilder

	// OrWhereFunc Add a nested "or where" statement to the query
	OrWhereFunc(callback QueryFunc) QueryBuilder

	// WhereIn Add a "where in" clause to the query.
	WhereIn(column string, args interface{}, whereType ...WhereJoinType) QueryBuilder

	// OrWhereIn Add an "or where in" clause to the query.
	OrWhereIn(column string, args interface{}) QueryBuilder

	// WhereNotIn Add a "where not in" clause to the query.
	WhereNotIn(column string, args interface{}, whereType ...WhereJoinType) QueryBuilder

	// OrWhereNotIn Add an "or where not in" clause to the query.
	OrWhereNotIn(column string, args interface{}) QueryBuilder

	// WhereBetween Add a where between statement to the query.
	WhereBetween(column string, args interface{}, whereType ...WhereJoinType) QueryBuilder

	// OrWhereBetween Add an or where between statement to the query.
	OrWhereBetween(column string, args interface{}) QueryBuilder

	// WhereNotBetween Add a where not between statement to the query.
	WhereNotBetween(column string, args interface{}, whereType ...WhereJoinType) QueryBuilder

	// OrWhereNotBetween Add an or where not between statement to the query.
	OrWhereNotBetween(column string, args interface{}) QueryBuilder

	// WhereIsNull Add a "where null" clause to the query.
	WhereIsNull(column string, whereType ...WhereJoinType) QueryBuilder

	// OrWhereIsNull Add an "or where null" clause to the query.
	OrWhereIsNull(column string) QueryBuilder

	// OrWhereNotNull Add an "or where not null" clause to the query.
	OrWhereNotNull(column string) QueryBuilder

	// WhereNotNull Add a "where not null" clause to the query.
	WhereNotNull(column string, whereType ...WhereJoinType) QueryBuilder

	// WhereExists Add an exists clause to the query.
	WhereExists(provider QueryProvider, where ...WhereJoinType) QueryBuilder

	// OrWhereExists Add an or exists clause to the query.
	OrWhereExists(provider QueryProvider) QueryBuilder

	// WhereNotExists Add a where not exists clause to the query.
	WhereNotExists(provider QueryProvider, where ...WhereJoinType) QueryBuilder

	// OrWhereNotExists Add a where not exists clause to the query.
	OrWhereNotExists(provider QueryProvider) QueryBuilder

	// Union Add a union statement to the query.
	Union(builder QueryBuilder, unionType ...UnionJoinType) QueryBuilder

	// UnionAll Add a union all statement to the query.
	UnionAll(builder QueryBuilder) QueryBuilder

	// UnionByProvider Add a union statement to the query and order by.
	UnionByProvider(builder QueryProvider, unionType ...UnionJoinType) QueryBuilder

	// UnionAllByProvider Add a union all statement to the query and order by.
	UnionAllByProvider(builder QueryProvider) QueryBuilder

	// GroupBy Add a "group by" clause to the query.
	GroupBy(columns ...string) QueryBuilder

	// Having Add a "having" clause to the query.
	Having(column string, args ...interface{}) QueryBuilder

	// OrHaving Add an "or having" clause to the query.
	OrHaving(column string, args ...interface{}) QueryBuilder

	// OrderBy Add an "order by" clause to the query.
	OrderBy(column string, columnOrderType ...OrderType) QueryBuilder

	// OrderByDesc Add a descending "order by" clause to the query.
	OrderByDesc(column string) QueryBuilder

	// InRandomOrder Put the query's results in random order.
	InRandomOrder(orderFunc ...OrderType) QueryBuilder

	// When Apply the callback's query changes if the given "value" is true.
	When(condition bool, callback QueryCallback, elseCallback ...QueryCallback) QueryBuilder

	// ToSql get the SQL representation of the query.
	ToSql() string

	// GetBindings get the current query value bindings in a flattened array.
	GetBindings() (results []interface{})

	// Offset Set the "offset" value of the query.
	Offset(offset int64) QueryBuilder

	// Skip Alias to set the "offset" value of the query.
	Skip(offset int64) QueryBuilder

	// Limit Set the "limit" value of the query.
	Limit(num int64) QueryBuilder

	// Take Alias to set the "limit" value of the query.
	Take(num int64) QueryBuilder

	// WithPagination Set the limit and offset for a given page.
	WithPagination(perPage int64, current ...int64) QueryBuilder

	// Chunk chunk the results of the query.
	Chunk(size int, handler func(collection Support.Collection, page int) error) error

	// ChunkById chunk the results of a query by comparing IDs.
	ChunkById(size int, handler func(collection Support.Collection, page int) error) error

	// SelectSql Gets the complete SQL string formed by the current specifications of this query builder.
	SelectSql() (string, []interface{})

	// SelectForUpdateSql Converts this instance into an UPDATE string in SQL.
	SelectForUpdateSql() (string, []interface{})
	CreateSql(value Support.Fields, insertType2 ...InsertType) (sql string, bindings []interface{})
	InsertSql(values []Support.Fields, insertType2 ...InsertType) (sql string, bindings []interface{})
	InsertIgnoreSql(values []Support.Fields) (sql string, bindings []interface{})
	InsertReplaceSql(values []Support.Fields) (sql string, bindings []interface{})
	DeleteSql() (sql string, bindings []interface{})
	UpdateSql(value Support.Fields) (sql string, bindings []interface{})

	// SetExecutor set executor.
	SetExecutor(executor SqlExecutor) QueryBuilder

	// Insert insert new records into the database.
	Insert(values ...Support.Fields) bool

	// InsertGetId insert a new record and get the value of the primary key.
	InsertGetId(values ...Support.Fields) int64

	// InsertOrIgnore insert new records into the database while ignoring errors.
	InsertOrIgnore(values ...Support.Fields) int64

	// InsertOrReplace Insert a new record into the database, and if it exists, delete this row of data first, and then insert new data.
	InsertOrReplace(values ...Support.Fields) int64

	// Create Save a new model and return the instance.
	Create(fields Support.Fields) interface{}

	// FirstOrCreate get the first record matching the attributes or create it.
	FirstOrCreate(values ...Support.Fields) interface{}

	// Update update records in the database.
	Update(fields Support.Fields) int64

	// UpdateOrInsert insert or update a record matching the attributes, and fill it with values.
	UpdateOrInsert(attributes Support.Fields, values ...Support.Fields) bool

	// UpdateOrCreate create or update a record matching the attributes, and fill it with values.
	UpdateOrCreate(attributes Support.Fields, values ...Support.Fields) interface{}

	// Get Execute the query as a "select" statement.
	Get() Support.Collection

	// SelectForUpdate Lock the selected rows in the table for updating.
	SelectForUpdate() Support.Collection

	// Find Execute a query for a single record by ID.
	Find(key interface{}) interface{}

	// First Execute the query and get the first result.
	First() interface{}

	// FirstOr Execute the query and get the first result or call a callback.
	FirstOr(provider IContainer.InstanceProvider) interface{}

	// FirstOrFail Execute the query and get the first result or throw an exception.
	FirstOrFail() interface{}

	// FirstWhere Add a basic where clause to the query, and return the first result.
	FirstWhere(column string, args ...interface{}) interface{}

	// Delete delete records from the database.
	Delete() int64

	// Paginate paginate the given query.
	Paginate(perPage int64, current ...int64) (Support.Collection, int64)

	// SimplePaginate paginate the given query into a simple paginator.
	SimplePaginate(perPage int64, current ...int64) Support.Collection

	// Bind binding Query Builder.
	Bind(QueryBuilder) QueryBuilder
}

type Model interface {

	// GetClass Get the class of the model.
	GetClass() Support.Class

	// GetTable Get the table associated with the model.
	GetTable() string

	// GetConnection Get the database connection for the model.
	GetConnection() string

	// GetPrimaryKey Get the primary key for the model.
	GetPrimaryKey() string
}

// MigrateHandler Database Migration Handler.
type MigrateHandler func(db DBConnection) error

type Migrate struct {
	Name       string
	Connection string
	CreatedAt  time.Time
	Up         MigrateHandler
	Down       MigrateHandler
}

type Migrations []Migrate
