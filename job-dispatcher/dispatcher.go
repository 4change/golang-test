package main

import "fmt"

// 待执行的任务队列
// A buffered channel that we can send work requests on.
var JobQueue chan Job

// 任务分发器, 一共有4个任务队列, 每个任务队列最多包含20个任务
type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	maxWorkers int
	WorkerPool chan chan Job
}

// 创建新的任务分发器, 一个任务分发器可以包含多个任务执行器, 每个任务执行器有一个任务队列, 该队列中包含多个待执行的任务
func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	// 创建并启动多个任务执行器
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	fmt.Println("Worker que dispatcher started...")
	for {
		select {
		// 从全局任务队列中取出一个任务, 并将该任务交给任务分发器进行分发
		case job := <-JobQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}
