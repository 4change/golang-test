package concurrent

import (
	"fmt"
	"sync"
	"testing"
)

var once sync.Once

func TestSyncOnce(t *testing.T) {
	var wg sync.WaitGroup

	// 10次输出, once.Do只执行一次
	for i := 0; i < 10; i++ {
		once.Do(onces)
		fmt.Println("count:", "---", i)
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onced)
			fmt.Println("213")

			wg.Done()
		}()
	}

	wg.Wait()
}

func onces() {
	fmt.Println("onces")
}

func onced() {
	fmt.Println("onced--------------------------")
}
