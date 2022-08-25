package concurrent_test

import (
	"fmt"
	"testing"
)

// channel的遍历方式二: 在明确知道channel长度时, 可以先使用普通的for循环进行channel的遍历, 在遍历后在进行channel的关闭
func TestChannelLoop_2_BeforeClose(t *testing.T) {
	goroutineNum := 10
	exitChan := make(chan bool, goroutineNum)

	for i := 1; i <= goroutineNum; i++ {
		go SubGoroutine(exitChan, i)
	}

	for i := 1; i <= goroutineNum; i++ {
		<-exitChan
	}
	close(exitChan)

	fmt.Println("Goroutine ID:", GetGoroutineID(), "*******************************************Exit Main Goroutine")
}

func SubGoroutine(exitChan chan bool, i int) {
	println("Goroutine Id:", GetGoroutineID(), "-----------------------------------------------------------i:", i)
	exitChan <- true
}

// channel的遍历方式一: 在不明确知道channel长度时, 需要先进行channel的关闭, 然后才能进行channel的遍历
func TestChannelLoop_1_AfterClose(t *testing.T) {
	exitChan := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Goroutine ID:", GetGoroutineID(), "---------------------------------------------------", i)
		}

		exitChan <- true
		close(exitChan)
	}()

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

	fmt.Println("Goroutine ID:", GetGoroutineID(), "*******************************************Exit Main Goroutine")
}
