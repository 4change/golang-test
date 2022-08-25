package distributed_lock

import (
	"sync"
	"testing"
)

// Lock try lock
type Lock struct {
	c chan struct{}				// Golang中空struct{}的作用和用法？
}

// NewLock generate a try lock
func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)	// 创建一个容量为1的通道
	l.c <- struct{}{}				// 向通道中放入一个结构体
	return l
}

// Lock try lock, return lock result
func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

// Unlock , Unlock the try lock
func (l Lock) Unlock() {
	l.c <- struct{}{}
}

//var counter int				// 全局变量

func TestTryLock(t *testing.T) {
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				// log error
				println("lock failed")
				return
			}
			counter++
			println("current counter", counter)
			l.Unlock()
		}()
	}

	wg.Wait()
}
