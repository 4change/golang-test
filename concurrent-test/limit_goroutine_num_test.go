package concurrent_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// 方案一: 通过阻塞通道限制Goroutine并发执行的数量
func Test_Limit_Goroutine_Num_By_Channel(t *testing.T) {
	ch := make(chan int, 5) // 通过阻塞通道限制Goroutine并发执行的数量为5
	for i := 0; i < 10; i++ {
		ch <- 1
		fmt.Println("the ch value send", ch)
		go func() {
			// 通过阻塞通道限制Goroutine的并发数量为5后, 再往通道中添加数据时会出现错误
			//<-ch
			fmt.Println("the ch value receive", ch)
		}()
		fmt.Println("the result i", i)
	}
}

// Pool Goroutine Pool
type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

// New 新建一个协程池
func NewPool(size int) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

// Add 新增一个执行
func (p *Pool) Add(delta int) {
	// delta为正数就添加
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	// delta为负数就减少
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

// Done 执行完成减一
func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}

// Wait 等待Goroutine执行完毕
func (p *Pool) Wait() {
	p.wg.Wait()
}

func Test_Limit_Goroutine_Num_By_WaitGroup(t *testing.T) {
	// 这里限制5个并发
	pool := NewPool(5)
	fmt.Println("the NumGoroutine begin is:", runtime.NumGoroutine())
	for i := 0; i < 20; i++ {
		pool.Add(1)
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println("the NumGoroutine continue is:", runtime.NumGoroutine())
			pool.Done()
		}(i)
	}
	pool.Wait()
	fmt.Println("the NumGoroutine done is:", runtime.NumGoroutine())
}
