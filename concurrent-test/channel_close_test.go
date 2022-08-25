package concurrent_test

import (
	"fmt"
	"testing"
)

func TestChannelClose(t *testing.T) {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	// channel被关闭后, 可以进行读取, 但不可以进行写入
	//intChan <- 300
	fmt.Println("okook~")
	n1 := <-intChan
	fmt.Println("n1 =", n1)
}
