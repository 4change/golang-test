package cron_test

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/robfig/cron"
	"log"
	"strconv"
	"testing"
)

func text() {
	log.Print("text")
}

// 测试每隔1秒执行一次的定时任务
func TestCronEverySecond(t *testing.T) {
	// 新建一个定时任务
	c := cron.New()

	// 定时任务执行的函数的定制
	// spec: 秒 分 时 日 月 星期,
	// "* * * * * *": 每秒执行一次该定时任务
	if err := c.AddFunc("* * * * * *", func() { text() }); err != nil {
		log.Println("执行 text() 定时任务出错: ", err.Error())
	}

	// 开始定时任务的调度
	c.Start()

	// 主线程持续循环, 等待子线程的执行, 相当于 for {}
	select {}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// 测试每隔5s执行一次的定时任务
func TestCronEvery5Second(t *testing.T) {
	// 定时任务的执行次数
	i := 0

	// 新建一个定时任务
	c := cron.New()

	// 定时任务执行的函数的定制
	// "*/5 * * * * ?": 每5s执行一次该定时任务
	err := c.AddFunc("*/5 * * * * ?", func() {
		i++
		log.Println("cron running:", i)
	})
	if err != nil {
		log.Println("定时任务函数定制出错----------------------------------------------------------------------------")
	}

	// 开始执行定时任务, 定时任务会被调度器调度执行
	c.Start()

	// 主线程进入等待状态
	select{}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type TestJob struct {
}

func (testJob TestJob)Run() {
	log.Println("testJob1...")
}

type TestJob2 struct {
}

func (testJob2 TestJob2)Run() {
	log.Println("testJob2...")
}

// 测试多个定时任务的执行
func TestMultiTask(t *testing.T) {
	i := 0
	// 新建cron调度器
	c := cron.New()

	// AddFunc
	spec := "*/5 * * * * ?"

	// AddFunc()和AddJob()方法的差别是什么?
	err := c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	if err != nil {
		log.Println("定时任务函数定制出错----------------------------------------------------------------------------")
	}

	// 添加定时任务
	err = c.AddJob(spec, TestJob{})
	if err != nil {
		log.Println("定时任务函数定制出错----------------------------------------------------------------------------")
	}

	err = c.AddJob(spec, TestJob2{})
	if err != nil {
		log.Println("定时任务函数定制出错----------------------------------------------------------------------------")
	}

	// 启动计划任务
	c.Start()

	// 关闭计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	select{}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type RedisJob struct {}

var JOB_RUN_NO = 1
var JOB_RUN_NO_KEY = "JOB_RUN_KEY_"

// 实现Run()方法, 即表示实现Job接口, 定义一个定时任务
func (redisJob RedisJob) Run() {
	//1. 连接到 redis
	conn, err := redis.Dial("tcp", "0.0.0.0:6699")
	if err != nil {
		log.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() //关闭..

	// 连接带密码的redis, 不带密码则去掉如下代码
	if _, err := conn.Do("AUTH", "123456"); err != nil {
		conn.Close()
		log.Println(err)
		return
	}

	//2. 通过 go 向 redis 写入数据 string [key-val]
	_, err = conn.Do("Set", JOB_RUN_NO_KEY + strconv.Itoa(JOB_RUN_NO), JOB_RUN_NO)
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	// 3. 通过 go 向 redis 读取数据 string [key-val]
	r, err := redis.String(conn.Do("Get", JOB_RUN_NO_KEY + strconv.Itoa(JOB_RUN_NO)))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}

	JOB_RUN_NO++

	// 因为返回 r 是 interface{}
	//因为 name 对应的值是 string ,因此我们需要转换
	//nameString := r.(string)
	fmt.Println("操作 ok ", r)
}

func TestCronRedis(t *testing.T) {
	// 新建cron调度器
	c := cron.New()

	// 添加定时任务
	err := c.AddJob("* * * * * ?", RedisJob{})
	if err != nil {
		log.Println("定时任务函数定制出错----------------------------------------------------------------------------")
	}

	// 启动计划任务
	c.Start()

	// 关闭计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	select{}
}

