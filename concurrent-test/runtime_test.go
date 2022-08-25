package concurrent_test

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"testing"
)

func TestNumCPU(t *testing.T) {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum:", cpuNum)

	// 可以自己设置程序运行过程中使用多少个CPU
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}

func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
