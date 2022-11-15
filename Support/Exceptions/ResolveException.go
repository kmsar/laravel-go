package Exceptions

import (
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/pkg/errors"
)

func ResolveException(v interface{}) IExeption.Exception {
	if v == nil {
		return nil
	}

	switch e := v.(type) {
	case IExeption.Exception:
		return e
	case error:
		return WithError(e, Support.Fields{})
	case string:
		return WithError(errors.New(e), Support.Fields{})
	case Support.Fields:
		if e["error"] != nil {
			return WithError(fmt.Errorf("%v", e["error"]), e)
		}
		if e["msg"] != nil {
			return WithError(fmt.Errorf("%v", e["msg"]), e)
		}
		if e["message"] != nil {
			return WithError(fmt.Errorf("%v", e["message"]), e)
		}
		if e["err"] != nil {
			return WithError(fmt.Errorf("%v", e["err"]), e)
		}
		return WithError(errors.New("Server exception"), e)
	default:
		return New("Server exception", Support.Fields{"err": v})
	}
}
