package defer_test

import (
	"fmt"
	"testing"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////
// defer调用并在后序的defer中调用recover进行异常处理
func Test_Defer_And_Recover_Post_Order(t *testing.T) {

	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()

	// 必须要先声明defer，否则recover()不能捕获到panic异常
	defer func() {
		// 捕获异常并进行处理
		if err := recover(); err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印后")
	}()

	panic("触发异常")

}
////////////////////////////////////////////////////////////////////////////////////////////////////
// defer调用并在中序的defer中调用recover进行异常处理
func Test_Defer_And_Recover_In_Order(t *testing.T) {

	defer func() { fmt.Println("打印前") }()

	// 必须要先声明defer，否则recover()不能捕获到panic异常
	defer func() {
		if err := recover();err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印中")
	}()

	defer func() { fmt.Println("打印后") }()

	panic("触发异常")

}
////////////////////////////////////////////////////////////////////////////////////////////////////
// defer调用并在每个defer中调用recover进行异常处理
func Test_Defer_And_Recover_Every_Order(t *testing.T) {

	defer func() {
		if err := recover();err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印前")
	}()

	// 必须要先声明defer，否则recover()不能捕获到panic异常
	defer func() {
		if err := recover();err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印中")
	}()

	defer func() {
		if err := recover();err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印后")
	}()

	panic("触发异常")

}
////////////////////////////////////////////////////////////////////////////////////////////////
func Test_Defer_Call_Order(t *testing.T) {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}