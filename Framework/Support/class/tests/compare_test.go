package tests

import (
	"github.com/kmsar/laravel-go/Framework/Support/Utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ComparableUser struct {
	Id   int
	Name string
}

type AliasComparableUser ComparableUser

func TestCompareIsEqual(t *testing.T) {
	assert.True(t, Utils.IsEqual(1, 1))
	assert.True(t, Utils.IsEqual(1, "1"))
	assert.True(t, Utils.IsEqual(1, true))
	assert.False(t, Utils.IsEqual(1, ""))
	assert.False(t, Utils.IsEqual(1, "0"))

	// 字符串
	assert.True(t, Utils.IsEqual("goal", "goal"))
	assert.True(t, Utils.IsEqual("1", 1))
	assert.True(t, Utils.IsEqual("true", true))

	// 结构体
	// 所有字段都一样
	assert.True(t, Utils.IsEqual(ComparableUser{Id: 1, Name: "goal"}, ComparableUser{Id: 1, Name: "goal"}))
	assert.True(t, Utils.IsEqual(ComparableUser{Id: 1}, ComparableUser{Id: 1}))

	// 部分字段不一样
	assert.False(t, Utils.IsEqual(ComparableUser{Id: 1, Name: "goal"}, ComparableUser{Id: 1}))

	// 类名不一样 false
	assert.False(t, Utils.IsEqual(ComparableUser{Id: 1, Name: "goal"}, AliasComparableUser{Id: 1, Name: "goal"}))

	// 数组或者切片
	// 完全一致 true
	assert.True(t, Utils.IsEqual([]int{1, 2}, []int{1, 2}))
	// 值不一致 false
	assert.False(t, Utils.IsEqual([]int{1, 2}, []int{1, 2, 44}))
	// 值一致 true
	assert.True(t, Utils.IsEqual([]int{1, 2}, []interface{}{1, 2}))
	assert.True(t, Utils.IsEqual([]int{1, 2}, []float64{1, 2}))
}

func TestOtherCompare(t *testing.T) {
	// 判断存在
	// 值存在 true
	assert.True(t, Utils.IsIn(1, []float32{1, 2.5}))
	// 可转换的值存在 true
	assert.True(t, Utils.IsIn(1, []interface{}{"1"}))
	// 值不存在 false
	assert.False(t, Utils.IsIn(100, []interface{}{"1"}))
	// 第二个参数类型不是 array或者slice false
	assert.False(t, Utils.IsIn(1, 2))

	// 判断不存在
	// 值不存在 true
	assert.True(t, Utils.IsNotIn(100, []interface{}{"1"}))
	// 可转换的值存在 false
	assert.False(t, Utils.IsNotIn(1, []interface{}{"1"}))
	// 值存在 false
	assert.False(t, Utils.IsNotIn(1, []interface{}{1}))
	assert.False(t, Utils.IsNotIn(1, []int{1}))
	// 第二个参数类型不是 array或者slice false
	assert.False(t, Utils.IsNotIn(1, 2))
}
