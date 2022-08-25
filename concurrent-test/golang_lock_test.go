package concurrent_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSyncMap(t *testing.T) {
	var syncMap sync.Map

	syncMap.Store(1, "a")             // store方法, 添加元素
	if v, ok := syncMap.Load(1); ok { // Load方法, 获得value
		fmt.Println("syncMap[1]=", v)
	}

	// LoadOrStore方法，获取或者保存
	// 参数是一对key: value，如果该key存在且没有被标记删除则返回原先的value（不更新）和true；不存在则store，返回该value和false
	if vv, ok := syncMap.LoadOrStore(1, "c"); ok { // 测试获取
		fmt.Println("syncMap[1]=", vv)
	}
	if vv, ok := syncMap.LoadOrStore(2, "c"); !ok { // 测试保存
		fmt.Println("syncMap[2]=", vv)
	}

	//遍历该map，参数是个函数，该函数的两个参数是遍历获得的key和value，返回一个bool值，当返回false时，遍历立刻结束。
	fmt.Println("syncMap Loop:")
	syncMap.Range(func(k, v interface{}) bool {
		fmt.Println(k, ":", v)
		return true
	})
}

func TestRWMutex(t *testing.T) {
	var rwm sync.RWMutex
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println("try to lock read ", i)
			rwm.RLock()
			fmt.Println("get locked ", i)
			time.Sleep(time.Second * 2)
			fmt.Println("try to unlock for reading ", i)
			rwm.RUnlock()
			fmt.Println("unlocked for reading ", i)
		}(i)
	}
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("try to lock for writing")
	rwm.Lock()
	fmt.Println("locked for writing")
}
