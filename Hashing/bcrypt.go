package Hashing

import (
	ISupport "github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Field"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	cost int
	salt string
}

func (b Bcrypt) mixWithSalt(value string) string {
	return value + b.salt
}

func (b *Bcrypt) Info(hashedValue string) ISupport.Fields {
	cost, _ := bcrypt.Cost([]byte(hashedValue))
	return ISupport.Fields{
		"cost": cost,
	}
}

func (b *Bcrypt) getCost(options ISupport.Fields) int {
	return Field.GetIntField(options, "cost", b.cost)
}

func (b *Bcrypt) Make(value string, options ISupport.Fields) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(b.mixWithSalt(value)), b.getCost(options))
	return string(bytes)
}

func (b *Bcrypt) Check(value, hashedValue string, _ ISupport.Fields) bool {
	hashedBytes := []byte(hashedValue)
	err := bcrypt.CompareHashAndPassword(hashedBytes, []byte(b.mixWithSalt(value)))
	return err == nil
}
