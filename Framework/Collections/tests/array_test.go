package tests

import (
	"errors"
	"fmt"

	"github.com/kmsar/laravel-go/Framework/Collections"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	collect, err := Collections.New(1)
	assert.Nil(t, collect)
	assert.Error(t, err, err)

	collect, err = Collections.New([]int{1})
	assert.NotNil(t, collect)
	assert.Nil(t, err)
}

func TestArray(t *testing.T) {
	intCollection := Collections.MustNew([]interface{}{
		1, 2, 3, true, "fd", "true",
	})
	fmt.Println(intCollection.ToFloat64Array())
	assert.True(t, intCollection.Len() == 6)

	intCollection.Map(func(data, index int) {
		fmt.Println(fmt.Sprintf("index %d ，data：%d", index, data))
	})

	intCollection.Map(func(data, index int, allData []interface{}) {
		if index == 0 {
			fmt.Println("allData", allData)
		}
		fmt.Println(fmt.Sprintf("第 %d 个，值：%d", index, data))
	})

	intCollection.Map(func(data string, index int) {
		fmt.Println(fmt.Sprintf("第 %d 个，值：%s", index, data))
	})
	intCollection.Map(func(data bool, index int) {
		fmt.Println(fmt.Sprintf("第 %d 个，值：%v", index, data))
	})

	intCollection.Map(func(data int) {
		fmt.Println("data: ", data)
	})
	fmt.Println(intCollection.ToIntArray())

	fmt.Println(intCollection.Map(func(data int) int {
		if data > 0 {
			return 1
		}
		return 0
	}).ToIntArray())
}

type User struct {
	id    int
	Name  string `json:"name"`
	Money float64
}

func TestToJson(t *testing.T) {
	users := Collections.MustNew([]User{
		{id: 1, Name: "qbhy", Money: 1},
		{id: 2, Name: "goal", Money: 2},
		{id: 3, Name: "collection", Money: 0},
	})
	any := Collections.MustNew([]interface{}{
		"1", 2, 3.0, true, users.First(), users,
	})

	fmt.Println(users.ToJson())
	fmt.Println(any.ToJson())
}

func TestStructArray(t *testing.T) {

	users := Collections.MustNew([]User{
		{id: 1, Name: "qbhy"},
		{id: 2, Name: "goal"},
	})

	users.Map(func(user User) {
		fmt.Printf("user: id:%d Name:%s \n", user.id, user.Name)
	})

	users.Map(func(user Support.Fields) {
		fmt.Printf("user: id:%v Name:%s \n", user["id"], user["name"])
	})

	assert.True(t, users.Map(func(user User) User {
		if user.id == 1 {
			user.Money = 100
		}
		return user
	}).Where("money", 100).Len() == 1)
}

func TestFilterArray(t *testing.T) {

	users := Collections.MustNew([]User{
		{id: 1, Name: "qbhy", Money: 10000000},
		{id: 2, Name: "goal", Money: 10},
	})

	fmt.Println("第一个数据", users.ToInterfaceArray()[0])

	richUsers := users.Filter(func(user User) bool {
		return user.Money > 100
	})

	assert.True(t, richUsers.Len() == 1)
	fmt.Println(richUsers.ToInterfaceArray())

	poorUsers := users.Skip(func(user User) bool {
		return user.Money > 100
	})

	assert.True(t, poorUsers.Len() == 1)
	fmt.Println(poorUsers.ToInterfaceArray())

	qbhyUsers := users.Where("name", "qbhy")

	assert.True(t, qbhyUsers.Len() == 1)
	fmt.Println(qbhyUsers.ToInterfaceArray())

	assert.True(t, users.WhereLte("money", 50).Len() == 1)
	assert.True(t, users.Where("money", "<=", 50).Len() == 1)

}

// TestAggregateArray 聚合函数测试
func TestAggregateArray(t *testing.T) {

	users := Collections.MustNew([]User{
		{id: 1, Name: "qbhy", Money: 10000000000000000},
		{id: 2, Name: "goal", Money: 10000000000000000},
		{id: 3, Name: "collection", Money: 0.645624123},
	}).(*Collections.Collection)

	// SafeSum、SafeAvg、SafeMax、SafeMin 等方法需要 *collection.Collection 类型
	fmt.Println("Sum", users.SafeSum("money"))
	fmt.Println("Avg", users.SafeAvg("money"))
	fmt.Println("Max", users.SafeMax("money"))
	fmt.Println("Min", users.SafeMin("money"))
	sum, _ := decimal.NewFromString("20000000000000000.645624123")
	avg, _ := decimal.NewFromString("6666666666666666.8818747076666667")
	max, _ := decimal.NewFromString("10000000000000000")
	min, _ := decimal.NewFromString("0.645624123")

	assert.True(t, users.SafeSum("money").Equal(sum))
	assert.True(t, users.SafeAvg("money").Equal(avg))
	assert.True(t, users.SafeMax("money").Equal(max))
	assert.True(t, users.SafeMin("money").Equal(min))

	users = Collections.MustNew([]User{
		{id: 1, Name: "qbhy", Money: 1},
		{id: 2, Name: "goal", Money: 2},
		{id: 3, Name: "collection", Money: 0},
	}).(*Collections.Collection)

	assert.True(t, users.Sum("money") == 3)
	assert.True(t, users.Avg("money") == 1)
	assert.True(t, users.Max("money") == 2)
	assert.True(t, users.Min("money") == 0)
	assert.True(t, users.Count() == 3)
}

// TestSortArray 测试排序功能
func TestSortArray(t *testing.T) {
	users := Collections.MustNew([]User{
		{id: 1, Name: "qbhy", Money: 12},
		{id: 2, Name: "goal", Money: 1},
		{id: 2, Name: "goal", Money: 15},
		{id: 2, Name: "goal99", Money: 99},
		{id: 3, Name: "collection", Money: -5},
		{id: 3, Name: "移动", Money: 10086},
	})

	fmt.Println(users.ToInterfaceArray())

	// 暂不支持转成 contracts.Fields
	usersOrderByMoneyDesc := users.Sort(func(user User, next User) bool {
		return user.Money > next.Money
	})
	fmt.Println(usersOrderByMoneyDesc.ToInterfaceArray())
	assert.True(t, usersOrderByMoneyDesc.ToInterfaceArray()[0].(User).Money == 10086)

	usersOrderByMoneyAsc := users.Sort(func(user User, next User) bool {
		return user.Money < next.Money
	})
	fmt.Println(usersOrderByMoneyAsc.ToInterfaceArray())
	assert.True(t, usersOrderByMoneyAsc.ToInterfaceArray()[0].(User).Money == -5)

	numbers := Collections.MustNew([]interface{}{
		8, 0, 1, 2, 0.6, 4, 5, 6, -0.2, 7, 9, 3, "10086",
	})

	sortedNumbers := numbers.Sort(func(i, j float64) bool {
		return i > j
	}).ToFloat64Array()

	fmt.Println(sortedNumbers)
	assert.True(t, sortedNumbers[0] == 10086)
}

// TestCombine 测试组合集合功能
func TestCombine(t *testing.T) {
	users := Collections.MustNew([]User{
		{id: 1, Name: "qbhy", Money: 12},
	})

	users = users.Push(User{id: 2, Name: "goal", Money: 1000})
	//users = users.Prepend(User{id: 2, Name: "goal", Money: 1000}) // 插入到开头

	assert.True(t, users.Len() == 2)
	fmt.Println(users.ToInterfaceArray())

	others := Collections.MustNew([]User{
		{id: 3, Name: "马云", Money: 100000000},
	})

	all := others.Merge(users).Sort(func(pre User, next User) bool {
		return pre.Money > next.Money
	})

	assert.True(t, all.Len() == 3)
	fmt.Println(all.ToInterfaceArray())
	fmt.Println(all.Only("money", "name").ToArrayFields())

	assert.True(t, all.First("name") == "马云") // 最有钱还是马云

	normalUsers := all.Where("money", ">", 100)
	assert.True(t, normalUsers.Len() == 2)                       // 两个普通人
	assert.True(t, normalUsers.Last("name") == "goal")           // 筛选不影响排序，跟马云比还差了点
	assert.False(t, normalUsers.IsEmpty())                       // 有普通人
	assert.True(t, normalUsers.Where("money", "<", 0).IsEmpty()) // 普通人都没有负债

	randomUsers := all.Random(2)
	// 随机获取两个数据
	assert.True(t, randomUsers.Len() == 2)
	fmt.Println(randomUsers.ToInterfaceArray())

	assert.True(t, all.Pull().(User).Name == "qbhy") // 从末尾取走一个
	assert.True(t, all.Len() == 2)                   // 判断取走后的长度
	assert.True(t, all.Shift().(User).Name == "马云")  // 从开头取走一个
	assert.True(t, all.Len() == 1)                   // 判断取走后的长度

}

func TestChunk(t *testing.T) {
	collect := Collections.MustNew([]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	})

	err := collect.Chunk(5, func(collection Support.Collection, page int) error {
		fmt.Printf("页码：%d，数量：%d %v\n", page, collection.Len(), collection.ToInterfaceArray())
		switch page {
		case 4:
			assert.True(t, collection.Len() == 4)
		default:
			assert.True(t, collection.Len() == 5)
		}
		return nil
	})
	assert.Nil(t, err)

	err = Collections.MustNew([]User{
		{id: 1, Name: "qbhy", Money: 12},
		{id: 2, Name: "goal", Money: 1},
		{id: 2, Name: "goal", Money: 15},
		{id: 2, Name: "goal99", Money: 99},
		{id: 3, Name: "collection", Money: -5},
		{id: 3, Name: "移动", Money: 10086},
	}).Chunk(3, func(collection Support.Collection, page int) error {
		assert.True(t, page == 1)
		assert.True(t, collection.First("name") == "qbhy")
		assert.True(t, collection.Last("name") == "goal")
		return errors.New("第一页退出")
	})

	assert.Error(t, err)

}
