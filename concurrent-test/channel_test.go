package concurrent

import (
	"fmt"
	"testing"
)

func Test_Channel(t *testing.T) {
	var ch chan int = nil   //make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			ch <- 5
		}
		close(ch)
	}()
	for i := 0; i < 15; i++ {
		fmt.Println(<-ch)
	}
}
