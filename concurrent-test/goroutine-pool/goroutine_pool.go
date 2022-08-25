package main

import (
	"fmt"
	"time"
)

type Pool struct {
    work chan func()   // 任务
    sem  chan struct{} // 数量
}

func New(size int) *Pool {
    return &Pool{
        work: make(chan func()),
        sem:  make(chan struct{}, size),
    }
}

func (p *Pool) NewTask(task func()) {
    select {
        case p.work <- task:
            fmt.Println("匹配work-----------")
        case p.sem <- struct{}{}:
            fmt.Println("匹配sem-----------")
            go p.worker(task)
    }
}

func (p *Pool) worker(task func()) {
    defer func() { 
        <-p.sem 
    }()
    
    for {
        task()
        task = <-p.work
    }
}

// func main() {
//     pool := New(128)
//     pool.NewTask(func(){
//         fmt.Println("run task")
//     })
// }

func main()  {
    pool := New(2)

    for i := 1; i <5; i++{
        pool.NewTask(func(){
            time.Sleep(2 * time.Second)
            fmt.Println(time.Now())
        })
    }

    // 保证所有的协程都执行完毕
    time.Sleep(5 * time.Second)
}