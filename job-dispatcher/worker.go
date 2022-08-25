package main

import (
	"log"
)

// Job represents the job to be run
type Job struct {
	Payload Payload
}

type Payload struct {
	Version            string                 `json:"version"`
	RequestID          string                 `json:"request_id"`
	BusinessID         string                 `json:"business_id"`
	FlowID             string                 `json:"flow_id"`
	Nodes              map[string]Node        `json:"nodes"`								// 决策流中的节点
	SrcCreateTimestamp CustomerTime           `json:"create_timestamp"`
	Output             map[string]interface{} `json:"output"`
	IsTestEvaluation   bool                   `json:"is_test_evaluation"`
	ErrorCode          *int                   `json:"error_code"`
	ErrorMsg           *string                `json:"error_msg"`
	GzipRawData        []byte
}

type Node struct {
	Type      string                 `json:"type"`
	ModelID   string                 `json:"model_id"`
	Output    map[string]interface{} `json:"output"`
	ErrorCode int                    `json:"error_code"`
	ErrorMsg  string                 `json:"error_msg"`
}

// 任务执行器
// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

// 新建一个任务执行器
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start() 方法启动 worker 的循环等待, 在收到退出信号后便退出
// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// 这里是什么用法？
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			// 从任务执行器的任务队列中取出一个任务并执行
			select {
			case job := <-w.JobChannel:
				log.Println(job)
				//job.Payload.MaintainMetrics()
				//if err := job.Payload.SendToDataLog(); err != nil {
				//	log.Println(job.Payload.RequestID, "error sending to datalog: %s", err.Error())
				//}

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
