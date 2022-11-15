package Logs

import (
	"errors"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"

	"sync"
	"testing"
)

type TestException struct {
	error
	fields Support.Fields
}

func (this TestException) Error() string {
	return this.error.Error()
}

func (this TestException) Fields() Support.Fields {
	return this.fields
}

func TestLogger(t *testing.T) {
	WithError(errors.New("error")).Info("info")

	WithFields(Support.Fields{"id": "1"}).Warn("info")

	WithException(TestException{
		error:  errors.New("error"),
		fields: Support.Fields{"id": 1, "name": "qbhy"},
	}).Info("info")
}

func TestWithField(t *testing.T) {
	wg := sync.WaitGroup{}
	WithError(errors.New("error")).WithField("field1", "1").Info("info")

	wg.Add(1)
	go (func() interface{} {
		WithError(errors.New("error")).WithField("field1", "1").Info("info")
		wg.Done()
		return nil
	})()

	wg.Wait()
}
