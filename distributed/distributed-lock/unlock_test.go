package distributed_lock

import (
	"sync"
	"testing"
)

var counter int				// 全局变量

func TestUnlock(t *testing.T) {
	println("init counter------------------------------------------------------------------------------", counter)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	println("end counter-------------------------------------------------------------------------------", counter)
}
