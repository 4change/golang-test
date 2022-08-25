package concurrent_test

import "testing"

func TestReadWriteChannel(t *testing.T) {
	// 创建两个带缓冲区的管道,加快数据处理
	var intChan = make(chan int, 100) // intChan: 用于读写操作的channel
	var exitChan = make(chan bool, 1) // exitChan: 用于控制main goroutine结束时机的channel

	// 创建两个goroutine, 同时读写intChan, 并在intChan读取完成后设置exitChan
	go writeChannel(intChan)
	go readChannel(intChan, exitChan)

	// 遍历一个channel, 当channel为空&未关闭时, 该channel返回什么数据？
	// 遍历一个channel, 当channel已关闭&遍历到channel末尾时, 该channel返回什么数据？
	for {
		i, ok := <-exitChan
		println("main goroutine--exitChan--i:", i)
		println("main goroutine--exitChan--ok:", ok)
		if !ok {
			break
		}
	}
}

// 向intChan写入数据, 写入完成后关闭intChan
// 若未关闭intChan, 则在遍历时会报deadlock错误
// 原因: 同一个channel, 在写入数据完成后, 该channel未显式关闭;
// 在遍历读取该channel时, 当读取完channel中已有的数据时，继续读取会因为读取数据为空而包deadlock错误

// 正常的状态是, 新创建的channel为空&未关闭, channel可读的标志被设置为false, 并等待goroutine向其中写入数据;
// 当channel中有数据时, channel可读的标志被设置为true, 等待goroutine从其中取出数据
func writeChannel(intChan chan int) {
	for i := 1; i <= 100; i++ {
		intChan <- i
		//time.Sleep(time.Second)
		println("write channel:", i)
	}

	println("close intChan")
	close(intChan)
}

// 从intChannel读取数据, 读取完成后设置exitChan告诉main goroutine可以退出, 然后关闭exitChan
func readChannel(intChan chan int, exitChan chan bool) {
	for {
		i, ok := <-intChan
		//time.Sleep(time.Second)
		println("read channel:", i)

		if !ok {
			break
		}
	}

	exitChan <- true
	close(exitChan)
}
