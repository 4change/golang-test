package redigo_test_test

import (
	"fmt"
	log "github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"reflect"
	"testing"
	"time"
)

//通过 go 向 redis 写入数据和读取数据
func Test_Redis(t *testing.T) {

	//1. 连接到 redis
	conn, err := redis.Dial("tcp", "0.0.0.0:6699")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() //关闭..

	// 连接带密码的redis, 不带密码则去掉如下代码
	if _, err := conn.Do("AUTH", "123456"); err != nil {
		conn.Close()
		fmt.Println(err)
		return
	}

	//2. 通过 go 向 redis 写入数据 string [key-val]
	_, err = conn.Do("Set", "name", "tomjerry 猫猫")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	// 3. 通过 go 向 redis 读取数据 string [key-val]
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	// 因为返回 r 是 interface{}
	//因为 name 对应的值是 string ,因此我们需要转换
	//nameString := r.(string)
	fmt.Println("操作 ok ", r)
}

type PSubscribeCallback func (pattern, channel, message string)

type PSubscriber struct {
	client redis.PubSubConn
	cbMap map[string]PSubscribeCallback
}

func (c *PSubscriber) PConnect() {
	// 链接到 redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Critical("redis dial failed.")
	}

	// 连接带密码的redis, 不带密码则去掉如下代码
	if _, err := conn.Do("AUTH", "123456"); err != nil {
		_ = conn.Close()
		fmt.Println(err)
		return
	}

	c.client = redis.PubSubConn{conn}
	c.cbMap = make(map[string]PSubscribeCallback)

	go func() {
		for {
			log.Debug("waiting for a keyEvent......")
			switch res := c.client.Receive().(type) {
			case redis.Subscription:
				fmt.Println("type of res: redis.Subscription----------------------------------------------------")
				fmt.Printf("res.Channel: %s------res.Kind: %s-----res.Count:%d\n", res.Channel, res.Kind, res.Count)
			case redis.PMessage:
				fmt.Println("type of res: redis.PMessage--------------------------------------------------------")
				pattern := res.Pattern
				channel := res.Channel
				message := string(res.Data)
				fmt.Printf("design-pattern:%s-----channel:%s-----message:%v\n", pattern, channel, message)
				c.cbMap[channel](pattern, channel, message)
			case error:
				log.Error("error handle...")
				continue
			default:
				fmt.Printf("type of res: %s---------------------------------------------\n", reflect.TypeOf(res))
			}
		}
	}()

}
func (c *PSubscriber)Psubscribe(channel interface{}, cb PSubscribeCallback) {
	err := c.client.PSubscribe(channel)
	if err != nil{
		log.Critical("redis Subscribe error.")
	}

	c.cbMap[channel.(string)] = cb
}

// redis key 超时监控回调事件
func SubCallback(patter , chann, msg string){
	fmt.Println("redis 订阅回调-----------------------------------------------------------------------------------")
	log.Debug( "SubCallback patter : " + patter + " channel : ", chann, " message : ", msg)
}

func TestRedis(t *testing.T) {
	var pSub PSubscriber
	// 连接到 redis
	pSub.PConnect()
	// redis key 超时监控, SubCallback 为回调事件
	pSub.Psubscribe("__keyevent@0__:expired", SubCallback)
	// 还可以是： `__keyspace@0__:cool`
	for{
		time.Sleep(1 * time.Second)
	}
}
