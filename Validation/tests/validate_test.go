package tests

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Validation"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Param struct {
	fields Support.Fields
}

func (p Param) Fields() Support.Fields {
	return p.fields
}

func (p Param) Rules() Support.Fields {
	return Support.Fields{
		"id":   "required,gte=1",
		"name": "required",
	}
}

func TestValidate(t *testing.T) {

	assert.True(t, len(Validation.Valid(Support.Fields{}, Support.Fields{
		"name": "required",
	})) == 1)

	assert.True(t, len(Validation.Valid(Support.Fields{
		"name": "啦啦啦",
	}, Support.Fields{
		"name": "required",
	})) == 0)

	assert.True(t, len(Validation.Form(Param{Support.Fields{
		"id":   1,
		"name": "goal",
	}})) == 0)

}
