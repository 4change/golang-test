package redis_v4_test

import (
	_redis "gopkg.in/redis.v4"
	"log"
	"testing"
	"time"
)

type Redis struct {
	Connector   *_redis.Client
	PubSub      *_redis.PubSub
}

var redis *Redis = nil

func NewRedis() bool {
	if redis == nil {
		redis = new(Redis)
		redis.Connector = _redis.NewClient(&_redis.Options{
			Addr: "127.0.0.1:6379",
			Password: "123456",
			DB: 0,
		})
		log.Println(nil, "Connected to Redis")
		err := redis.Init()
		if err != nil {
			log.Println(nil, "Cannot setup Redis:", err.Error())
			return false
		}
		return true
	}
	return false
}

func (this *Redis) Init() error {
	pubsub, err := this.Connector.Subscribe("__keyevent@0__:expired")
	if err != nil {
		return err
	}
	//defer pubsub.Close()
	this.PubSub = pubsub

	go func() {
		for {
			msgi, err := this.PubSub.Receive()
			if err != nil {
				log.Println(nil, "PubSub error:", err.Error())
			}

			switch msg := msgi.(type) {
			case *_redis.Message:
				log.Println(nil, "Received", msg.Payload, "on channel", msg.Channel)
			default:
				log.Println(nil, "Got control message", msg)
			}
		}
	}()

	return nil
}

func ExpiredSubscribe(client *_redis.Client) error {
	pubSub, err := client.PSubscribe("__keyevent@0__:expired")
	if err != nil {
		return err
	}

	go func() {
		for {
			message, err := pubSub.Receive()
			if err != nil {
				log.Println(nil, "PubSub error:", err.Error())
			}

			switch msg := message.(type) {
			case *_redis.Message:
				log.Println("channel==" + msg.Channel + "---pattern==" + msg.Pattern + "---payload==" + msg.Payload + "---string==" + msg.String())
			default:
				log.Println(nil, "Got control message", msg)
			}
		}
	}()

	return nil
}


func TestPub(t *testing.T) {
	client := createClient()
	_ = ExpiredSubscribe(client)

	for {
		time.Sleep(1 * time.Second)
	}
}
