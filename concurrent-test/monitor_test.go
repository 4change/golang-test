package concurrent

import (
	"fmt"
	"testing"
	"time"
)

func TestMonitor(t *testing.T) {

	// 创建一个可以接受任何值的通道abort
	abort := make(chan struct{})

	// 创建新协程,用于从标准输入读取数据
	go func() {
		for {
			// 时间戳产生通道
			tick := time.Tick(800 * time.Millisecond)

			select {
			case <-tick:
				fmt.Println("tick-------------------------------------------------------------------------------")
			case <-abort:
				fmt.Println("Launch aborted!")
			}
		}
	}()

	select {
	}

}