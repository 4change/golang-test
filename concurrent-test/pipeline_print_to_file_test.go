package concurrent_test

import (
	"log"
	"os"
	"testing"
)

func Test_Pipeline_Print_To_File(t *testing.T) {
	a, _ := os.OpenFile("./a.txt", os.O_WRONLY|os.O_APPEND, 0666)
	b, _ := os.OpenFile("./b.txt", os.O_WRONLY|os.O_APPEND, 0666)
	c, _ := os.OpenFile("./c.txt", os.O_WRONLY|os.O_APPEND, 0666)
	d, _ := os.OpenFile("./d.txt", os.O_WRONLY|os.O_APPEND, 0666)

	files := []*os.File{a, b, c, d} // 文件写入顺序
	sign := make(chan int, 1)       // pipeline通道, 控制pipeline顺序
	for i := 1; i < 100; i++ {
		sign <- 1
		go out1(files[0], sign)
		sign <- 1
		go out2(files[1], sign)
		sign <- 1
		go out3(files[2], sign)
		sign <- 1
		go out4(files[3], sign)

		files = append(files[len(files)-1:], files[:len(files)-1]...) // 文件写入顺序的修改
	}
	a.Close()
	b.Close()
	c.Close()
	d.Close()
}

func out1(f *os.File, c chan int) int {
	f.Write([]byte("1 "))
	// f.Close()
	log.Println(f.Name() + " write finish...")
	<-c
	return 1
}

func out2(f *os.File, c chan int) int {
	f.Write([]byte("2 "))
	// f.Close()
	log.Println(f.Name() + " write finish...")
	<-c
	return 2
}

func out3(f *os.File, c chan int) int {
	f.Write([]byte("3 "))
	// f.Close()
	log.Println(f.Name() + " write finish...")
	<-c
	return 3
}

func out4(f *os.File, c chan int) int {
	f.Write([]byte("4 "))
	// f.Close()
	log.Println(f.Name() + " write finish...")
	<-c
	return 4
}
