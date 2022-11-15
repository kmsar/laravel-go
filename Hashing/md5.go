package Hashing

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

type Md5 struct {
	salt string
}

func (m *Md5) mixWithSalt(value string) string {
	return value + m.salt
}

func (m *Md5) Info(_ string) Support.Fields {
	return nil
}

func (m *Md5) Make(value string, _ Support.Fields) string {
	d := []byte(m.mixWithSalt(value))
	hash := md5.New()
	hash.Write(d)
	return hex.EncodeToString(hash.Sum(nil))

}

func (m *Md5) Check(value, hashedValue string, _ Support.Fields) bool {
	return m.Make(value, nil) == hashedValue
}

func Md5Hash(str string) string {
	m := Md5{}
	return m.Make(str, nil)
}
