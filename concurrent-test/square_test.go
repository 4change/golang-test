package concurrent_test

import (
	"fmt"
	"testing"
)

func Test_Square_Infinity(t *testing.T) {
	naturals := make(chan int)
	squarer := make(chan int)

	// Counter goroutine
	go func() {
		for i := 0; ; i++ {
			naturals <- i
		}
	}()

	// Squarer goroutine
	go func() {
		for i := range naturals {
			squarer <- i * i
		}
	}()

	// Printer goroutine
	go func() {
		for i := range squarer {
			fmt.Println("square:", i)
		}
	}()

	// main goroutine
	for {

	}
}

func Test_Square_Limit(t *testing.T) {
	naturals := make(chan int)
	squarer := make(chan int)
	exitChan := make(chan bool, 3)

	// Counter goroutine
	go func() {
		for i := 0; i <= 100; i++ {
			naturals <- i
		}
		close(naturals)

		exitChan <- true
	}()

	// Squarer goroutine
	go func() {
		for i := range naturals {
			squarer <- i * i
		}
		close(squarer)

		exitChan <- true
	}()

	// Printer goroutine
	go func() {
		for i := range squarer {
			fmt.Println("square:", i)
		}

		exitChan <- true
	}()

	// main goroutine
	for i := 0; i < 3; i++ {
		<-exitChan
	}
	close(exitChan)
}
