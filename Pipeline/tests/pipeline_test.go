package tests

import (
	"fmt"

	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IContainer"

	"github.com/pkg/errors"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Container"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IPipeline"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Pipeline"
	"testing"
)

type User struct {
	Id   int
	Name string
}

func TestPipeline(t *testing.T) {
	pipe := Pipeline.New(Container.New())
	result := pipe.Send(User{Id: 1, Name: "goal"}).
		Through(
			func(user User, next IPipeline.Pipe) interface{} {
				fmt.Println("中间件1-start")
				result := next(user)
				fmt.Println("中间件1-end")
				return result
			},
			func(user User, next IPipeline.Pipe) interface{} {
				fmt.Println("中间件2-start")
				result := next(user)
				fmt.Println("中间件2-end")
				return result
			},
		).
		Then(func(user User) interface{} {
			fmt.Println("then", user)
			return user.Id
		})

	fmt.Println("穿梭结果：", result)
	/**
	中间件1-start
	中间件2-start
	then {1 goal}
	中间件2-end
	中间件1-end
	穿梭结果： 1
	*/
}

// TestPipelineException 测试异常情况
func TestPipelineException(t *testing.T) {
	defer func() {
		recover()
	}()
	pipe := Pipeline.New(Container.New())
	pipe.Send(User{Id: 1, Name: "goal"}).
		Through(
			func(user User, next IPipeline.Pipe) interface{} {
				fmt.Println("中间件1-start")
				result := next(user)
				fmt.Println("中间件1-end", result)
				return result
			},
			func(user User, next IPipeline.Pipe) interface{} {
				fmt.Println("中间件2-start")
				result := next(user)
				fmt.Println("中间件2-end", result)
				return result
			},
		).
		Then(func(user User) {
			panic(errors.New("报个错"))
		})
	/**
	中间件1-start
	中间件2-start
	*/
}

// TestStaticPipeline 测试调用magical函数
func TestStaticPipeline(t *testing.T) {
	// 应用启动时就准备好的中间件和控制器函数，在大量并发时用 StaticPipeline 可以提高性能
	middlewares := []IContainer.MagicalFunc{
		Container.NewMagicalFunc(func(user User, next IPipeline.Pipe) interface{} {
			fmt.Println("中间件1-start")
			result := next(user)
			fmt.Println("中间件1-end", result)
			return result
		}),
		Container.NewMagicalFunc(func(user User, next IPipeline.Pipe) interface{} {
			fmt.Println("中间件2-start")
			result := next(user)
			fmt.Println("中间件2-end", result)
			return result
		}),
	}
	controller := Container.NewMagicalFunc(func(user User) int {
		fmt.Println("then", user)
		return user.Id
	})

	pipe := Pipeline.Static(Container.New())
	result := pipe.SendStatic(User{Id: 1, Name: "goal"}).
		ThroughStatic(middlewares...).
		ThenStatic(controller)

	fmt.Println("穿梭结果", result)

	pipe = Pipeline.Static(Container.New())
	result = pipe.SendStatic(User{Id: 1, Name: "goal"}).
		ThroughStatic(middlewares...).
		ThenStatic(controller)
	fmt.Println("穿梭结果", result)
	/**
	中间件1-start
	中间件2-start
	then {1 goal}
	中间件2-end 1
	中间件1-end 1
	穿梭结果 1
	*/
}

// TestPurePipeline 测试纯净的 pipeline
func TestPurePipeline(t *testing.T) {
	// 如果你的应用场景对性能要求极高，不希望反射影响你，那么你可以试试下面这个纯净的管道
	pipe := Pipeline.Pure()
	result := pipe.SendPure(User{Id: 1, Name: "goal"}).
		ThroughPure(
			func(user interface{}, next IPipeline.Pipe) interface{} {
				fmt.Println("中间件1-start")
				result := next(user)
				fmt.Println("中间件1-end", result)
				return result
			},
			func(user interface{}, next IPipeline.Pipe) interface{} {
				fmt.Println("中间件2-start")
				result := next(user)
				fmt.Println("中间件2-end", result)
				return result
			},
		).
		ThenPure(func(user interface{}) interface{} {
			fmt.Println("then", user)
			return user.(User).Id
		})
	fmt.Println("穿梭结果", result)
	/**
	中间件1-start
	中间件2-start
	then {1 goal}
	中间件2-end 1
	中间件1-end 1
	穿梭结果 1
	*/
}
