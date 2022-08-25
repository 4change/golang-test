package io

import (
	"bufio"
	"fmt"
	"os"
)

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	var str = "曾---------n"
	var filename = "./test.txt"
	var f *os.File
	var err1 error
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	defer f.Close()
	if err1 != nil {
		panic(err1)
	}
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n, _ := w.WriteString(str)
	str = "ces000000000000000hi"
	n, _ = w.WriteString(str)
	n, _ = w.WriteString("asdfasfasdfasf")
	n, _ = w.WriteString("ASFASDFASDFASDF\n")
	fmt.Printf("写入 %d 个字节n", n)
	w.Flush()
}
