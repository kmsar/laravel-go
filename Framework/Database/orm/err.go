package gorose

import (
	"errors"
	"fmt"
)

type Error uint

type Lang uint

const (
	CHINESE Lang = iota

	ENGLISH

	CHINESE_TRADITIONAL
)

const (
	ERR_PARAMS_COUNTS Error = iota

	ERR_PARAMS_MISSING

	ERR_PARAMS_FORMAT
)

func (e *Err) Default() map[Error]string {
	return map[Error]string{
		ERR_PARAMS_COUNTS:  "参数数量有误",
		ERR_PARAMS_MISSING: "参数缺失",
		ERR_PARAMS_FORMAT:  "参数格式错误",
	}
}

var langString = map[Lang]string{
	CHINESE:             "chinese",
	ENGLISH:             "english",
	CHINESE_TRADITIONAL: "chinese_traditional",
}

func (l Lang) String() string {
	return langString[l]
}

type Err struct {
	lang Lang
	err  map[Lang]map[Error]string
}

var gErr *Err

func init() {
	var tmpLang = make(map[Lang]map[Error]string)
	gErr = &Err{err: tmpLang}
	gErr.lang = CHINESE
	gErr.Register(gErr.Default())
}

func NewErr() *Err {
	return gErr
}

func (e *Err) SetLang(l Lang) {
	e.lang = l
}

func (e *Err) GetLang() Lang {
	return e.lang
}

func (e *Err) Register(err map[Error]string) {
	e.err[e.GetLang()] = err
}

func (e *Err) Get(err Error) string {
	return e.err[e.GetLang()][err]
}

func GetErr(err Error, args ...interface{}) error {
	var argreal string
	if len(args) > 0 {
		argreal = fmt.Sprint(":", args)
	}
	return errors.New(fmt.Sprint(
		NewErr().
			Get(err),
		argreal))
}
