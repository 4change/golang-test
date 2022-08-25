package concurrent_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	myMap = make(map[int]int, 10)

	// 这里的uint64可以换成float类型重新计算
	myConcurrentMap = make(map[int]uint64, 10)

	// 声明一个全局的互斥锁
	// lock是一个全局的互斥锁
	// sync是包, synchronized: 同步
	// Mutex: 互斥
	lock sync.Mutex
)

// factorialWithMap: 计算各个数的阶乘，并放入map
func factorialWithMap(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res //concurrent map writes?
}

// 需求： 计算1-200(uint64只能表示60以内有效的数据)的各个数的阶乘, 并且把各个数的阶乘放入到map中, 最后打印出来, 要求使用goroutine完成
// 思路
// 		1. 编写一个函数， 来计算各个数的阶乘， 并放入到 map 中.
// 		2. 我们启动的协程多个， 将计算结果放入到 map 中
// 		3. map 应该做出一个全局的.
func TestFactorialWithMap(t *testing.T) {
	// 启动200个协程, 完成阶乘的计算, 并放入map中
	for i := 1; i <= 50; i++ {
		go factorialWithMap(i)
	}

	//休眠 10 秒钟【 第二个问题 】
	time.Sleep(time.Second * 10)

	// 遍历map, 输出结果
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}

// factorialWithMap: 计算各个数的阶乘，并放入map
func factorialWithConcurrentMap(n int) {
	var res uint64 = 1
	for i := 1; i <= n; i++ {
		res *= uint64(i)
	}

	lock.Lock() // 加锁
	myConcurrentMap[n] = res
	lock.Unlock() // 解锁
}

func TestFactorialWithConcurrentMap(t *testing.T) {
	// 启动100个协程, 完成阶乘的计算, 并放入map中
	for i := 1; i <= 50; i++ {
		go factorialWithConcurrentMap(i)
	}

	//休眠 10 秒钟【 第二个问题 】
	time.Sleep(time.Second * 10)

	// 遍历map, 输出结果
	lock.Lock()
	for i, v := range myConcurrentMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}

// factorialWithMap: 计算各个数的阶乘，并放入map
func factorialWithConcurrentMapAndChannel(n int, exitChan chan bool) {
	var res uint64 = 1
	for i := 1; i <= n; i++ {
		res *= uint64(i)
	}

	// 加锁, 避免map的并发写问题
	lock.Lock() // 加锁
	myConcurrentMap[n] = res
	lock.Unlock() // 解锁

	exitChan <- true
}

func TestFactorialWithConcurrentMapAndChannel(t *testing.T) {
	chanLen := 50 // 计算阶乘的协程数
	factorialChan := make(chan bool, chanLen)

	// 启动50个协程, 完成阶乘的计算, 并放入map中
	for i := 1; i <= chanLen; i++ {
		go factorialWithConcurrentMapAndChannel(i, factorialChan)
	}

	// 遍历factorialChan, 遍历完后进行关闭
	for i := 0; i < chanLen; i++ {
		<-factorialChan
	}
	close(factorialChan)

	// 遍历map, 输出结果
	for i, v := range myConcurrentMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}
