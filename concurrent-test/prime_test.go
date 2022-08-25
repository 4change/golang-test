package concurrent_test

import (
	"fmt"
	"testing"
)

// putNum: 向 intChan 放入 100000 个数
func putNum(intChan chan int) {
	for i := 1; i <= 100000; i++ {
		intChan <- i
	}
	// 关闭 intChan
	close(intChan)
}

// 从 intChan 取出数据， 并判断是否为素数,如果是， 就放入到 primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool // 是否为素数的flag

	for {
		num, ok := <-intChan
		// intChan 取不到..
		if !ok {
			break
		}

		flag = true //假设是素数
		// 判断 num 是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { // 说明该 num 不是素数
				flag = false
				break
			}
		}

		if flag {
			// 将这个数就放入到 primeChan
			primeChan <- num
		}
	}

	fmt.Println("有一个 primeNum 协程因为取不到数据， 退出")
	//这里我们还不能关闭 primeChan
	//向 exitChan 写入 true
	exitChan <- true
}

func TestPrime(t *testing.T) {
	intChan := make(chan int, 2000)   // 存放输入数据的管道
	primeChan := make(chan int, 2000) // 存放输出结果的管道
	exitChan := make(chan bool, 4)    // 标识primeNum协程退出的管道, 4 个

	// 开启一个协程putNum, 向 intChan 放入 100000 个数
	go putNum(intChan)

	// 开启 4 个协程， 从 intChan 取出数据， 并判断是否为素数, 如果是, 就放入到 primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		// 等待四个统计素数的协程退出
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		// 当我们从 exitChan 取出了 4 个结果， 就可以放心的关闭 primeChan
		close(primeChan)
	}()

	//遍历我们的 primeChan ,把结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		// 将结果输出
		fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("main 线程退出")
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func TestPrimeExercise_1(t *testing.T) {
	exitChan := make(chan bool, 4)
	numChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)

	go PutNum(numChan)

	for i := 0; i < 4; i++ {
		go JudgePrime(numChan, primeChan, exitChan)
	}

	// 这里为什么要单独放在一个goroutine中执行?
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		close(primeChan)
	}()

	for {
		if prime, ok := <-primeChan; ok {
			fmt.Println(prime)
		} else {
			break
		}
	}
}

func PutNum(numChan chan int) {
	for i := 0; i <= 10000; i++ {
		numChan <- i
	}

	close(numChan)
}

func JudgePrime(numChan, primeChan chan int, exitChan chan bool) {
	for {
		num, ok := <-numChan
		if !ok {
			break
		}

		flag := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			primeChan <- num
		}
	}

	exitChan <- true
}
