package func_test

import (
	"fmt"
	"testing"
)

// 声明 query 类型, 底层为 func(string) string 类型
type query func(string) string

// name string
// vs ...query: query 类型的数组 vs, 实际就是 func(string) string 构成的一个数组
func exec(name string, vs ...query) string {
	ch := make(chan string)

	// 声明一个匿名函数fn, 参数为 i int, 无返回值
	fn := func(i int) {
		fmt.Println(fmt.Sprintf("执行-----------------------------------------------------------------vs[%d]", i))
		// 调用 vs[i] 对应的 func(string) string 函数, 传入参数 name, 执行结果放入通道 ch
		ch <- vs[i](name)
	}

	for i, _ := range vs {
		go fn(i)
	}

	return <-ch
}

// return <-ch 只执行一次，所以不管传入多少 query 函数, 都只是读取最先执行完的 query 的结果
func TestFunc(t *testing.T) {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})

	fmt.Println(ret)
}
