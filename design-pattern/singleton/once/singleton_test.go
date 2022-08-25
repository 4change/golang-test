package once

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println("协程 ", strconv.Itoa(i), " 输出-------------------------", GetSingleton() == GetSingleton())
			wg.Done()
		}(i)
	}

	wg.Wait()
}

