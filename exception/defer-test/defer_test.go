package defer_test

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	// defer表达式中的i是对for循环中i的引用, 到最后, i加到5, 故最后全部打印5
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func TestDeferWithParameter(t *testing.T) {
	for i := 0; i < 5; i++ {
		// 传入的 i，会立即被求值保存为 idx
		defer func(idx int) {
			fmt.Println(idx)
		}(i)
	}
}

func TestDeferExec(t *testing.T) {
	a := 1
	defer fmt.Println("the value of a1:",a)
	a++

	defer func() {
		fmt.Println("the value of a2:",a)
	}()
}