package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	lineSum int               // 总代码行数, 全局变量, 需控制并发修改
	mutex   = new(sync.Mutex) // 互斥锁
)

var (
	rootPath = "/media/fly/D/go/src/yuncross/crossrouter"		// 要进行代码统计的根路径
	// 被排除(不进行代码统计)的子路径
	noDirs     = [...]string{"/vendor/golang.org", "/vendor/github.com", "/vendor/google.golang.org", "/vendor/gopkg.in", "/vendor/yuncross"}
	suffixName = ".go" // 要进行代码统计的文件后缀
)

func main() {
	// 命令行参数定制
	argsLen := len(os.Args)
	if argsLen == 2 {
		rootPath = os.Args[1]
	} else if argsLen == 3 {
		rootPath = os.Args[1]
		suffixName = os.Args[2]
	}

	done := make(chan bool)						// done channel用于控制main goroutine的结束时机
	go codeLineSum(rootPath, done)
	<-done

	fmt.Println("total line-----------------------------------------------------------------------------", lineSum)
}

func codeLineSum(root string, done chan bool) {
	var goes int              // 子goroutine的数目，对每一个文件单独开一个goroutine进行代码行数统计

	goDone := make(chan bool) // goDone channel用于控制codeLineSum goroutine的结束时机

	// 校验当前目录是否为要进行代码统计的子目录
	isDstDir := checkDir(root)

	defer func() {
		// 捕获异常并处理
		if pan := recover(); pan != nil {
			fmt.Printf("root: %s, panic:%#v\n", root, pan)
		}

		// 结束所有子协程
		// waiting for his children done
		for i := 0; i < goes; i++ {
			<-goDone
		}

		// 结束父协程
		// this goroutine done, notify his parent
		done <- true
	}()

	if !isDstDir {
		return
	}

	// 获取根文件信息
	rootFile, err := os.Lstat(root)
	checkErr(err)

	// 打开根路径
	rootDir, err := os.Open(root)
	checkErr(err)
	defer rootDir.Close()

	if rootFile.IsDir() {
		// 获取目录剩余的FileInfo构成的Slice
		fis, err := rootDir.Readdir(0)
		checkErr(err)

		for _, fi := range fis {
			// 跳过目录中的隐藏文件和文件夹　.fileName--在linux中表示隐藏文件或文件夹
			if strings.HasPrefix(fi.Name(), ".") {
				continue
			}

			// 对每一个文件, 开启一个goroutine进行代码行数统计
			goes++

			// 递归处理目录
			if fi.IsDir() {
				go codeLineSum(root+"/"+fi.Name(), goDone)
			} else {
				go readFile(root+"/"+fi.Name(), goDone)
			}
		}
	} else {
		// 根路径指向一个文件，则只开一个协程去统计该文件的行数
		goes = 1 // if rootFile is a file, current goroutine has only one child
		go readFile(root, goDone)
	}
}

// 统计单个文件的代码行数
//		filename string: 文件名, 绝对路径
//		done chan bool: 控制main goroutine结束的channel, 每个统计单个文件代码行数的goroutine结束时, 会向该channel写入true, 表示当前goroutine的结束
func readFile(filename string, done chan bool) {
	var line int	// 当前文件的代码行数

	// 根据文件后缀名判定当前文件是否为要进行代码行数统计的目标文件
	isDstFile := strings.HasSuffix(filename, suffixName)

	defer func() {
		// 异常处理
		if pan := recover(); pan != nil {
			fmt.Printf("filename: %s, panic:%#v\n", filename, pan)
		}

		// 如果当前文件是目标文件，则修改全局变量linesum
		if isDstFile {
			addLineNum(line)
			fmt.Printf("file %s complete, line = %d\n", filename, line)
		}

		// 当前goroutine结束, 向done channel写入true, 通知父goroutine当前goroutine的结束
		// this goroutine done, notify his parent
		done <- true
	}()

	// 如果当前文件不是要进行代码行数统计的文件, 则直接退出, 在函数退出前执行defer相关代码
	if !isDstFile {
		return
	}

	// 打开当前文件，用于后续读取，在退出前关闭文件
	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	// 循环读取文件中的每一行，isPrefix表示在返回数据时是否对该行进行了切割
	reader := bufio.NewReader(file)
	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}

		// 未对该行进行切割, 则line自增
		if !isPrefix {
			line++
		}
	}
}

// 校验当前目录是否为要进行代码统计的目录
// check whether this dir is the dest dir
func checkDir(dirPath string) bool {
	for _, dir := range noDirs {
		//fmt.Printf("rootPath + dir = %v, dirPath = %v \n", rootPath + dir, dirPath)
		if rootPath + dir == dirPath {
			return false
		}
	}
	return true
}

func addLineNum(num int) {
	mutex.Lock()
	defer mutex.Unlock()

	lineSum += num
}

// if error happened, throw a panic, and the panic will be recover in defer function
func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
