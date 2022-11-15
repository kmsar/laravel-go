package gorose

type IOrmSession interface {
	Begin() (err error)
	Rollback() (err error)
	Commit() (err error)

	Query(sqlstring string, args ...interface{}) ([]Data, error)
	Execute(sqlstring string, args ...interface{}) (int64, error)

	LastInsertId() int64
	LastSql() string

	GetIBinder() IBinder
	SetUnion(u interface{})
	GetUnion() interface{}
}
