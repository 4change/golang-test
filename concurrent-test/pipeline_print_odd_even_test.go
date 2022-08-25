package concurrent_test

import (
	"fmt"
	"sync"
	"testing"
)

// 好好理解一下这个程序的执行流程
func Test_Print_Odd_Even_By_Channel(t *testing.T) {
	startChan := make(chan int)
	notifyEven := make(chan int, 1)
	notifyOdd := make(chan int, 1)
	done := make(chan int, 2)

	// 打印奇数
	go func() {
		<-startChan
		for i := 1; i < 1000; i += 2 {
			fmt.Println("print odd-------------------------", i, "------------------------------", GetGoroutineID())
			notifyEven <- 1 // 奇数打印完，通知开始打印偶数
			<-notifyOdd
		}
		notifyEven <- 1
		//close(notifyEven)
		done <- 1
	}()

	// 打印偶数
	go func() {
		<-notifyEven
		for i := 2; i <= 1000; i += 2 {
			fmt.Println("print even-------------------------", i, "-------------", GetGoroutineID())
			notifyOdd <- 1 // 偶数打印完，通知开始打印奇数
			<-notifyEven
		}
		//close(notifyOdd)
		done <- 1
	}()

	startChan <- 1
	for i := 0; i < 2; i++ {
		<-done
	}
	fmt.Println("print odd even done !**********************************************************", GetGoroutineID())
}

func Test_Print_Odd_Even_By_WaitGroup(t *testing.T) {
	startChan := make(chan int)
	notifyEven := make(chan int, 1)
	notifyOdd := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		<-startChan
		for i := 1; i < 1000; i += 2 {
			fmt.Println("print odd-------------------------", i, "------------------------------", GetGoroutineID())
			notifyEven <- 1
			<-notifyOdd
		}
		notifyEven <- 1
		wg.Done()
	}()

	go func() {
		<-notifyEven
		for i := 2; i <= 1000; i += 2 {
			fmt.Println("print even-------------------------", i, "-------------", GetGoroutineID())
			notifyOdd <- 1
			<-notifyEven
		}
		wg.Done()
	}()

	startChan <- 1

	wg.Wait()
	fmt.Println("print odd even done !**********************************************************", GetGoroutineID())
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func Test_Print_Odd_Even_Interview_Exercise_1(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	startChan := make(chan int, 1)
	oddChan := make(chan int, 1)
	evenChan := make(chan int, 1)

	go func() {
		<-startChan
		oddChan <- 1
		for i := 1; i <= 1000; i += 2 {
			<-oddChan
			fmt.Println("print odd----------------------------", i)
			evenChan <- 1
		}
		<-oddChan
		close(oddChan)
		wg.Done()
	}()

	go func() {
		for i := 2; i <= 1000; i += 2 {
			<-evenChan
			fmt.Println("print even----------------------------------------------------------------------------", i)
			oddChan <- 1
		}

		close(evenChan)
		wg.Done()
	}()

	startChan <- 1
	wg.Wait()
	close(startChan)

	fmt.Println("print odd even done !**********************************************************", GetGoroutineID())
}
