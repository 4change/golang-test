package concurrent

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 在主线程(可以理解成进程)中， 开启一个 goroutine, 该协程每隔 1 秒输出 "hello,world"
// 在主线程中也每隔一秒输出"hello,golang", 输出 10 次后， 退出程序
// 要求主线程和 goroutine 同时执行.

// 编写一个函数，每个1s输出"hello world"
func GoroutinePrint() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test(), hello world, " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}

func TestGoroutinePrint(t *testing.T) {
	// 使用go关键字开启一个协程
	go GoroutinePrint()

	for i := 1; i <= 10; i++ {
		fmt.Println("main(), hello golang, " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
