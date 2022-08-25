package concurrent_test

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var bufChan = make(chan int, 1000)
var msgChan = make(chan int)

func SingleProducer(max int) {
	for i := 0; i < max; i++ {
		bufChan <- i
		fmt.Println("producer----------------------------------------------------------------------------------", i)
	}

	close(bufChan)
}

func SingleConsumer() {
	for {
		data, ok := <-bufChan
		if !ok {
			break
		}
		fmt.Println("consumer------------------------------------", data)
	}

	msgChan <- 1
}

func Test_Single_Producer_Consumer(t *testing.T) {
	max := 100

	go SingleProducer(max)
	go SingleConsumer()

	<-msgChan
	close(msgChan)

	fmt.Println("Single Producer Consumer Task Done !-------------------------------------------------------------")
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var strChan = make(chan string)

// write data to channel
func writer(max int, producerName string) {
	for {
		for i := 0; i < max; i++ {
			bufChan <- i
			fmt.Fprintf(os.Stderr, "%v write: %d\n", os.Getpid(), i)
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// read data fro m channel
func reader(customerName string) {
	for {
		r := <-bufChan
		fmt.Printf("%s read value: %d\n", customerName, r)
	}
	strChan <- customerName
}

func TestWriterAndReader(t *testing.T) {
	max := 100

	// 开启多个writer的goroutine，不断地向channel中写入数据
	go writer(max, "writer1")
	go writer(max, "writer2")

	// 开启多个reader的goroutine，不断的从channel中读取数据，并处理数据
	go reader("read1")
	go reader("read2")
	go reader("read3")

	// 获取三个reader的任务完成状态
	name1 := <-strChan
	name2 := <-strChan
	name3 := <-strChan

	fmt.Printf("%s,%s,%s: All is done!!", name1, name2, name3)
}
