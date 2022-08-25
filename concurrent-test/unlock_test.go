package concurrent

import (
	"sync"
	"testing"
)

// 全局变量
var counter int
func TestUnlock(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	println(counter)
}
