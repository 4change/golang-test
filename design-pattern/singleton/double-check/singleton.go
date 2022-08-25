package double_check

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var mu sync.Mutex
var initialized int32

type singleton struct{}
var instance *singleton

func GetSingleton() *singleton {

	if atomic.LoadInt32(&initialized) == 1 { // 原子操作
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		fmt.Println("Instance 实例化-----------------------------------------------------------------------------------")
		instance = &singleton{}
		atomic.StoreInt32(&initialized, 1)
	}

	return instance
}
