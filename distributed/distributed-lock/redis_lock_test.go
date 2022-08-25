package distributed_lock

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"testing"
	"time"
)

func incr() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	var lockKey = "counter_lock"
	var counterKey = "counter"
	// lock
	resp := client.SetNX(lockKey, 1, time.Second*10)
	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}
	// counter ++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			// log err
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)
	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}

}

func TestRedisLock(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}

	wg.Wait()
}
