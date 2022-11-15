package gorose

type IOrmExecute interface {
	GetForce() bool

	Insert(data ...interface{}) (int64, error)
	InsertGetId(data ...interface{}) (int64, error)
	Update(data ...interface{}) (int64, error)

	Increment(args ...interface{}) (int64, error)
	Decrement(args ...interface{}) (int64, error)

	Delete() (int64, error)

	Force() IOrm
}
