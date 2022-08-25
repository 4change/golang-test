package concurrent_test

import "testing"

func TestFuncInputParameter(t *testing.T) {
	exitChan := make(chan bool, 2)

	para1 := "abcdefghi"
	para2 := 123456789

	go func(input string) {
		for j := 0; j < 10; j++ {
			println("Goroutine Id:", GetGoroutineID(), "--------------------input:", input, "--------------j:", j)
		}

		exitChan <- true
	}(para1)

	go func(input int) {
		for j := 0; j < 10; j++ {
			println("Goroutine Id:", GetGoroutineID(), "--------------------input:", input, "--------------j:", j)
		}

		exitChan <- true
	}(para2)

	for i := 0; i < 2; i++ {
		<-exitChan
	}
	close(exitChan)

	println("Goroutine Id:", GetGoroutineID(), "---------------------------------------------Exit Main Goroutine")
}
