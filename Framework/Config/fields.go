package Config

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
)

type FieldsProvider struct {
	Data Support.Fields
}

func (provider FieldsProvider) Fields() Support.Fields {
	return provider.Data
}
