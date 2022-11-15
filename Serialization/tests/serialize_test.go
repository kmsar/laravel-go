package tests

import (
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Container"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Serialization"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Serialization/serializers"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/class"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Id   string `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

type Detail struct {
	Title     string `json:"title"`
	User      User   `json:"user"`
	Followers []User `json:"followers"`
}

var (
	UserClass = class.Make(new(User))
)

func TestJsonSerialize(t *testing.T) {

	json := serializers.Json{}

	jsonStr := json.Serialize(User{
		Id:   "10086",
		Name: "goal",
	})
	fmt.Println(jsonStr)
	assert.True(t, "{\"id\":\"10086\",\"name\":\"goal\"}" == jsonStr)

	var user User
	assert.Nil(t, json.UnSerialize(jsonStr, &user))
	assert.True(t, user.Id == "10086")
	fmt.Println(user)
}

func TestXmlSerialize(t *testing.T) {
	serializer := serializers.Xml{}
	serialized := serializer.Serialize(User{
		Id:   "10086",
		Name: "goal",
	})
	fmt.Println(serialized)
	var user User
	assert.True(t, serialized == "<User><id>10086</id><name>goal</name></User>")
	assert.Nil(t, serializer.UnSerialize(serialized, &user))
	assert.True(t, user.Id == "10086")
	fmt.Println(user)
}

func TestComplexJsonSerialize(t *testing.T) {

	json := serializers.Json{}

	jsonStr := json.Serialize(Detail{
		Title: "goal",
		User: User{
			Id:   "10086",
			Name: "goal",
		},
		Followers: []User{
			{
				Id:   "1",
				Name: "goal",
			},
			{
				Id:   "2",
				Name: "goal",
			},
		},
	})
	fmt.Println(jsonStr)
	assert.True(t, "{\"title\":\"goal\",\"user\":{\"id\":\"10086\",\"name\":\"goal\"},\"followers\":[{\"id\":\"1\",\"name\":\"goal\"},{\"id\":\"2\",\"name\":\"goal\"}]}" == jsonStr)

	var detail Detail
	assert.Nil(t, json.UnSerialize(jsonStr, &detail))
	assert.True(t, detail.User.Id == "10086")
	fmt.Println(detail)
}

func TestGobSerialize(t *testing.T) {

	gob := serializers.Gob{}

	jsonStr := gob.Serialize(User{
		Id:   "10086",
		Name: "goal",
	})
	fmt.Println(jsonStr)

	var user User
	assert.Nil(t, gob.UnSerialize(jsonStr, &user))
	fmt.Println(user)
}

func TestClassSerializer(t *testing.T) {
	serializer := Serialization.NewClassSerializer(Container.New())

	serializer.Register(UserClass)
	serializer.Register(class.Make(new(Detail)))

	serialized := serializer.Serialize(Detail{
		Title: "goal",
		User: User{
			Id:   "10086",
			Name: "goal",
		},
		Followers: []User{
			{
				Id:   "1",
				Name: "goal",
			},
			{
				Id:   "2",
				Name: "goal",
			},
		},
	})

	fmt.Println("serialized", serialized)
	detail, err := serializer.Parse(serialized)
	assert.Nil(t, err)
	assert.True(t, detail.(*Detail).Title == "goal")
	fmt.Println(detail)
}

/*
*
goos: darwin
goarch: amd64
pkg: github.com/goal-web/serialization/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkClassUnserialize
BenchmarkClassUnserialize-4   	  187540	      6283 ns/op
*/
func BenchmarkClassUnserialize(b *testing.B) {
	serializer := Serialization.NewClassSerializer(Container.New())

	serializer.Register(UserClass)
	serializer.Register(class.Make(new(Detail)))

	serialized := serializer.Serialize(Detail{
		Title: "goal",
		User: User{
			Id:   "10086",
			Name: "goal",
		},
		Followers: []User{
			{
				Id:   "1",
				Name: "goal",
			},
			{
				Id:   "2",
				Name: "goal",
			},
		},
	})

	for i := 0; i < b.N; i++ {
		serializer.Parse(serialized)
	}
}

/*
*
goos: darwin
goarch: amd64
pkg: github.com/goal-web/serialization/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkClassSerialize
BenchmarkClassSerialize-4   	  285645	      4218 ns/op
*/
func BenchmarkClassSerialize(b *testing.B) {
	serializer := Serialization.NewClassSerializer(Container.New())

	serializer.Register(UserClass)
	serializer.Register(class.Make(new(Detail)))

	for i := 0; i < b.N; i++ {
		serializer.Serialize(Detail{
			Title: "goal",
			User: User{
				Id:   "10086",
				Name: "goal",
			},
			Followers: []User{
				{
					Id:   "1",
					Name: "goal",
				},
				{
					Id:   "2",
					Name: "goal",
				},
			},
		})

	}
}

/*
*
goos: darwin
goarch: amd64
pkg: github.com/goal-web/serialization/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkClassSerializer
BenchmarkClassSerializer-4   	   95265	     11996 ns/op
*/
func BenchmarkClassSerializer(b *testing.B) {
	serializer := Serialization.NewClassSerializer(Container.New())

	serializer.Register(class.Make(new(Detail)))
	detail := Detail{
		Title: "goal",
		User: User{
			Id:   "10086",
			Name: "goal",
		},
		Followers: []User{
			{
				Id:   "1",
				Name: "goal",
			},
			{
				Id:   "2",
				Name: "goal",
			},
		},
	}

	for i := 0; i < b.N; i++ {
		serialized := serializer.Serialize(detail)
		_, _ = serializer.Parse(serialized)
	}
}

/*
*
goos: darwin
goarch: amd64
pkg: github.com/goal-web/serialization/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkJsonSerialize
BenchmarkJsonSerialize-4   	  894805	      1673 ns/op
*/
func BenchmarkJsonSerialize(b *testing.B) {
	serializer := serializers.Json{}
	for i := 0; i < b.N; i++ {

		jsonStr := serializer.Serialize(User{
			Id:   "10086",
			Name: "goal",
		})

		var user User
		_ = serializer.UnSerialize(jsonStr, &user)
	}
}

/*
*
goos: darwin
goarch: amd64
pkg: github.com/goal-web/serialization/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkGobSerialize
BenchmarkGobSerialize-4   	   46285	     22629 ns/op
*/
func BenchmarkGobSerialize(b *testing.B) {
	serializer := serializers.Gob{}
	for i := 0; i < b.N; i++ {

		jsonStr := serializer.Serialize(User{
			Id:   "10086",
			Name: "goal",
		})

		var user User
		_ = serializer.UnSerialize(jsonStr, &user)
	}
}

/*
*
goos: darwin
goarch: amd64
pkg: github.com/goal-web/serialization/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkXmlSerialize
BenchmarkXmlSerialize-4   	  188887	      5706 ns/op
*/
func BenchmarkXmlSerialize(b *testing.B) {
	serializer := serializers.Xml{}
	for i := 0; i < b.N; i++ {

		jsonStr := serializer.Serialize(User{
			Id:   "10086",
			Name: "goal",
		})

		var user User
		_ = serializer.UnSerialize(jsonStr, &user)
	}
}
