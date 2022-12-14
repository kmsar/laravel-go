package tests

import (
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/class"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Id   int `json:"id" db:"identify"`
	name string
}

var UserClass1 = class.Make(new(User))

func TestDefine(t *testing.T) {
	assert.True(t, class.Container.IsSubClass(class.Application))
	assert.True(t, class.Application.Implements(class.Container))
}

func TestNewByTag(t *testing.T) {
	user := UserClass1.NewByTag(Support.Fields{
		"identify": 2,
		"name":     "goal", // 为导出字段不支持解析
	}, "db").(User)

	fmt.Println("user.Id", user)
	assert.True(t, user.Id == 2 && user.name == "")

	user = UserClass1.NewByTag(Support.Fields{
		"id": 1, // 没有 db 字段没定义，默认就用 json 字段
	}, "db").(User)
	assert.True(t, user.Id == 1)

	var user1 = UserClass1.NewByTag(Support.Fields{
		"id": []byte("1"), // 没有 db 字段没定义，默认就用 json 字段
	}, "db").(User)
	assert.True(t, user1.Id == 1)
	fmt.Println(user1)
}

/*
*
goos: darwin
goarch: amd64
pkg: github.com/goal-web/supports/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkNewByTag
BenchmarkNewByTag-4   	 2916794	       373.4 ns/op
*/
func BenchmarkNewByTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UserClass1.NewByTag(Support.Fields{
			"identify": 2,
		}, "db")
	}
}

/*
*
goos: darwin
goarch: amd64
pkg: github.com/goal-web/supports/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkComplexNewByTag
BenchmarkComplexNewByTag-4   	 2521668	       491.1 ns/op
*/
func BenchmarkComplexNewByTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := UserClass1.NewByTag(Support.Fields{
			"identify": i,
		}, "db").(User)
		user.Id++
	}
}
