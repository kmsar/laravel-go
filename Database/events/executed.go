package events

import "time"

type QueryExecuted struct {
	Sql        string
	Bindings   []interface{}
	Connection string
	Time       time.Duration
	Error      error
}

func (e *QueryExecuted) Event() string {
	return "QUERY_EXECUTED"
}

func (e *QueryExecuted) Sync() bool {
	return true
}
