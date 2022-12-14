package distributed_lock

import (
	"sync"
	"testing"
)

//var counter int					// ćšć±ćé

func TestWGLock(t *testing.T) {
	println("init counter------------------------------------------------------------------------------", counter)
	var wg sync.WaitGroup
	var l sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter++
			l.Unlock()
		}()
	}

	wg.Wait()
	println(counter)
	println("end counter-------------------------------------------------------------------------------", counter)
}