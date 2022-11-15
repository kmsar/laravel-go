package gorose

type IOrm interface {
	IOrmApi
	IOrmQuery
	IOrmExecute
	IOrmSession

	Close()
	BuildSql(operType ...string) (string, []interface{}, error)
	Table(tab interface{}) IOrm

	Fields(fields ...string) IOrm
	AddFields(fields ...string) IOrm

	Distinct() IOrm
	Data(data interface{}) IOrm

	Group(group string) IOrm
	GroupBy(group string) IOrm
	Having(having string) IOrm
	Order(order string) IOrm
	OrderBy(order string) IOrm
	Limit(limit int) IOrm
	Offset(offset int) IOrm
	Page(page int) IOrm

	Join(args ...interface{}) IOrm
	LeftJoin(args ...interface{}) IOrm
	RightJoin(args ...interface{}) IOrm
	CrossJoin(args ...interface{}) IOrm

	Where(args ...interface{}) IOrm
	OrWhere(args ...interface{}) IOrm
	WhereNull(arg string) IOrm
	OrWhereNull(arg string) IOrm
	WhereNotNull(arg string) IOrm
	OrWhereNotNull(arg string) IOrm
	WhereRegexp(arg string, expstr string) IOrm
	OrWhereRegexp(arg string, expstr string) IOrm
	WhereNotRegexp(arg string, expstr string) IOrm
	OrWhereNotRegexp(arg string, expstr string) IOrm
	WhereIn(needle string, hystack []interface{}) IOrm
	OrWhereIn(needle string, hystack []interface{}) IOrm
	WhereNotIn(needle string, hystack []interface{}) IOrm
	OrWhereNotIn(needle string, hystack []interface{}) IOrm
	WhereBetween(needle string, hystack []interface{}) IOrm
	OrWhereBetween(needle string, hystack []interface{}) IOrm
	WhereNotBetween(needle string, hystack []interface{}) IOrm
	OrWhereNotBetween(needle string, hystack []interface{}) IOrm

	GetDriver() string

	SetBindValues(v interface{})
	GetBindValues() []interface{}
	ClearBindValues()
	Transaction(closers ...func(db IOrm) error) (err error)
	Reset() IOrm
	ResetTable() IOrm
	ResetWhere() IOrm
	GetISession() ISession
	GetOrmApi() *OrmApi

	SharedLock() *Orm

	LockForUpdate() *Orm
}
