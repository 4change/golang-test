package concurrent_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 方案三: main goroutine延时等待其他goroutine的退出
func TestExitMain_3_Sleep(t *testing.T) {
	go func() {
		fmt.Println("Goroutine 1")
	}()

	go func() {
		fmt.Println("Goroutine 2")
	}()

	time.Sleep(time.Second * 1) // 睡眠1秒，等待上面两个协程结束

	fmt.Println("Exit Main Goroutine")
}

// 方案二: 利用管道实现同步，子协程结束后关闭管道，然后主协程再执行
func TestExitMain_2_Channel(t *testing.T) {
	ch := make(chan struct{})
	count := 2 // count 表示活动的协程个数

	go func() {
		fmt.Println("Goroutine 1")
		ch <- struct{}{} // 协程结束，发出信号
	}()

	go func() {
		fmt.Println("Goroutine 2")
		ch <- struct{}{} // 协程结束，发出信号
	}()

	for range ch {
		// 每次从ch中接收数据，表明一个活动的协程结束
		count--
		// 当所有活动的协程都结束时，关闭管道
		if count == 0 {
			close(ch)
		}
	}

	fmt.Println("Exit Main Goroutine")
}

// 方案一: 利用sync.WaitGroup实现协程同步
//		sync.WaitGroup内部实现了一个计数器，用来记录未完成的操作个数，WaitGroup 中 goroutine 的执行是无序的，它提供了三个方法:
//			> Add()用来添加计数。
//			> Done()用来在操作结束时调用，使计数减一。
//			> Wait()用来等待所有的操作结束，即计数变为0，该函数会在计数不为0时等待，在计数为0时立即返回。
func TestExitMain_1_WaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(2) // 因为有两个动作，所以增加2个计数
	go func() {
		fmt.Println("Goroutine 1")
		wg.Done() // 操作完成，减少一个计数
	}()

	go func() {
		fmt.Println("Goroutine 2")
		wg.Done() // 操作完成，减少一个计数
	}()

	wg.Wait() // 等待，直到计数为0

	// ...			继续主协程的各种其他操作

	fmt.Println("exit main goroutine")
}

// 协程同步面临的问题: main goroutine在子goroutine还未执行完成时便退出, 导致子goroutine无法顺利执行
//		例如如下测试案例: Goroutine 1和Goroutine 2很可能无法得到执行, 因为main goroutine退出会导致子goroutine没有时间执行
func TestExitMain_Exception(t *testing.T) {
	go func() {
		fmt.Println("Goroutine 1")
	}()

	go func() {
		fmt.Println("Goroutine 2")
	}()

	fmt.Println("Exit Main Goroutine")
}
