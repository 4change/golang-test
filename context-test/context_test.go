package context_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_Control_Goroutine_WaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	// 向 WaitGroup 中添加要执行的 Goroutine 的数量
	wg.Add(2)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1号完成---------------------------------------------------------------------------------------")
		// 设置 WaitGroup 中一个 Goroutine 执行完成的标志
		wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2号完成---------------------------------------------------------------------------------------")
		// 设置 WaitGroup 中一个 Goroutine 执行完成的标志
		wg.Done()
	}()

	// 等待 WaitGroup 中所有 Goroutine 执行完毕
	wg.Wait()

	fmt.Println("好了，大家都干完了，放工------------------------------------------------------------------------------")
}

func Test_Control_Goroutine_Select_Chan(t *testing.T) {
	// 用于接收信号的chan
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

	time.Sleep(5 * time.Second)
	fmt.Println("可以了，通知监控停止---------------------------------------------------------------------------------")
	stop<- true
}

func Test_Context_Control_One_Goroutine(t *testing.T) {
	// 依据 Background Context 创建一个 Context, 取消函数为 cancel
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():					// 接收到 context cancel 函数的结束信号, 结束当前context的进行
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止--------------------------------------------------------------------------------")

	// cancel()函数, 通知结束 context
	cancel()

	time.Sleep(5 * time.Second)
}

func Test_Context_Control_Multi_Goroutine(t *testing.T) {
	// 以 Background Context 为父节点创建 Context, 取消函数为 cancel
	ctx, cancel := context.WithCancel(context.Background())

	go Watch(ctx,"【监控1】")			// 以 context 为上下文, 创建一个 goroutine 并运行
	go Watch(ctx,"【监控2】________")
	go Watch(ctx,"【监控3】________________")

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止---------------------------------------------------------------------------------")

	// 结束 ctx 的运行, 以该 context 为父节点的 goroutine 都会结束
	cancel()
}

func Watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name,"监控退出，停止了...")
			return
		default:
			fmt.Println(name,"goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

var key string = "name"

func Test_Context_WithValue(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	//附加值
	valueCtx := context.WithValue(ctx, key, "【监控1】")
	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
