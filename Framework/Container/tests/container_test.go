package tests

import (
	"fmt"

	"github.com/kmsar/laravel-go/Framework/Container"
	"github.com/kmsar/laravel-go/Framework/Contracts/IContainer"
	"github.com/kmsar/laravel-go/Framework/Support/supports-master/utils"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type DemoParam struct {
	Id string
}

type DemoComponent struct {
	Param DemoParam
}

func TestArgumentsTypeMap(t *testing.T) {
	args := Container.NewArgumentsTypeMap([]interface{}{"啦啦啦", DemoParam{Id: "111"}})
	str := args.Pull("string")
	fmt.Println(str)
	assert.True(t, str == "啦啦啦")

	args = Container.NewArgumentsTypeMap([]interface{}{})
	assert.True(t, args.Pull("string") == nil)
}

func TestBaseContainer(t *testing.T) {
	app := Container.New()

	app.Instance("a", "a")
	assert.True(t, app.HasBound("a"))
	assert.True(t, app.Get("a") == "a")

	app.Alias("a", "A")

	assert.True(t, app.Get("A") == "a")
	assert.True(t, app.HasBound("A"))

	app.NamedBind("DemoParam", func() DemoParam {
		return DemoParam{Id: "测试一下"}
	})

	assert.True(t, app.Get(utils.GetTypeKey(reflect.TypeOf(DemoParam{}))).(DemoParam).Id == "测试一下")

	app.Call(Container.NewMagicalFunc(func(param DemoParam) {
		assert.True(t, param.Id == "测试一下")
	}))

}

func TestContainer(t *testing.T) {
	app := Container.New()

	app.NamedBind("DemoParam", func() DemoParam {
		return DemoParam{Id: "没有外部参数的话，从容器中获取"}
	})

	fn := Container.NewMagicalFunc(func(param DemoParam) string {
		return param.Id
	})

	// 自己传参
	assert.True(t, app.Call(fn, DemoParam{Id: "优先使用外部参数"})[0] == "优先使用外部参数")

	// 不传参，使用容器中的实例
	assert.True(t, app.Call(fn)[0] == "没有外部参数的话，从容器中获取")

}

type DemoStruct struct {
	Param  DemoParam `di:""`       // 注入对应类型的实例
	Config string    `di:"config"` // 注入指定 key 的实例
}

func TestContainerMake(t *testing.T) {
	app := Container.New()

	app.Instance("config", "通过容器设置的配置")

	app.NamedBind("DemoParam", func() DemoParam {
		return DemoParam{Id: "没有外部参数的话，从容器中获取"}
	})

	demo := &DemoStruct{}

	app.DI(demo)

	fmt.Println(demo)
}

func TestAliasType(t *testing.T) {
	app := Container.New()

	app.NamedSingleton("param", func() DemoParam {
		return DemoParam{
			Id: "a",
		}
	})

	type AliasParam DemoParam

	app.Call(Container.NewMagicalFunc(func(param AliasParam) {
		fmt.Println(param)
	}), app.Get("param"))
}

type DemoStruct2 struct {
	DemoStruct
}

func (d *DemoStruct2) Construct(Container2 IContainer.IContainer) {
	d.DemoStruct = Container2.Get("struct").(DemoStruct)
}

// 调用方法支持注入自定义类
func TestAutoContainer(t *testing.T) {
	app := Container.New()

	app.NamedSingleton("struct", func() DemoStruct {
		return DemoStruct{
			Param:  DemoParam{Id: "id"},
			Config: "config",
		}
	})

	//struct2Type := reflect.TypeOf(DemoStruct2{})
	//struct2Value := reflect.New(struct2Type).Interface()
	struct2Value := &DemoStruct2{}

	app.DI(struct2Value)

	app.Call(Container.NewMagicalFunc(func(struct2 DemoStruct2) {
		assert.True(t, struct2.Config == "config" && struct2.Param.Id == "id")
	}))

	app.Call(Container.NewMagicalFunc(func(struct2 DemoStruct2, struct1 DemoStruct) { // 因为 DemoStruct2 实现了 contracts.Component 所以不会使用自定义参数
		assert.True(t, struct2.Config == "config" && struct2.Param.Id == "id")
		assert.True(t, struct1.Config == "config22" && struct1.Param.Id == "custom")
	}), DemoStruct{
		Param:  DemoParam{Id: "custom"},
		Config: "config22",
	})
}

// 测试控制器执行
type DemoDependent struct {
	Id string
}

type DemoController struct {
	Dep DemoDependent `di:""` // 表示需要注入
}

func (this *DemoController) PrintDep() {
	fmt.Println(this.Dep)
}

func TestControllerCall(t *testing.T) {
	app := Container.New()
	app.NamedSingleton("DemoDependent", func() DemoDependent {
		return DemoDependent{
			Id: "id ddd",
		}
	})

	controller := &DemoController{}

	app.DI(controller)

	app.Call(Container.NewMagicalFunc(controller.PrintDep))
}

func TestCallAndDIContainer(t *testing.T) {
	app := Container.New()

	app.Call(Container.NewMagicalFunc(func(Container2 IContainer.IContainer) {
		fmt.Println(Container2)
	}))
}
