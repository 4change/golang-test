package distributed_lock

import (
	"github.com/zieckey/etcdsync"
	"log"
	"testing"
	"time"
)

func TestEtcdLock(t *testing.T) {
	m, err := etcdsync.New("/lock", 10, []string{"http://127.0.0.1:2379"})
	if m == nil || err != nil {
		log.Printf("etcdsync.New failed")
		return
	}

	err = m.Lock()
	if err != nil {
		log.Printf("etcdsync.Lock failed")
		return
	}

	log.Printf("etcdsync.Lock OK")
	log.Printf("Get the lock. Do something here.")

	time.Sleep(10 * time.Second)

	err = m.Unlock()
	if err != nil {
		log.Printf("etcdsync.Unlock failed")
	} else {
		log.Printf("etcdsync.Unlock OK")
	}
}
