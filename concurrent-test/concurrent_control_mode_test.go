package concurrent_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_Concurrent_Control_Mode_By_Channel(t *testing.T) {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop<- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func Test_Concurrent_Control_Mode_By_WaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1号完成")
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2号完成")
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("好了，大家都干完了，放工")
}

func Test_Concurrent_Control_Mode_By_Context(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
