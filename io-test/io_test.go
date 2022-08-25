package io

import (
	"os"
	"strings"
	"testing"
	"time"
)

// 打印内容到文件中
func TestFileAppend(t *testing.T) {
	fd, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fdTime := time.Now().Format("2006-01-02 15:04:05")
	fdContent := strings.Join([]string{"======", fdTime, "=====", " This is file content ! ", "\n"}, "")
	buf := []byte(fdContent)
	fd.Write(buf)
	fd.Close()
}
