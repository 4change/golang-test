package goleak_test

import (
	"go.uber.org/goleak"
	"testing"

)

func leak() {
	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
	}()
}

//func TestLeak(t *testing.T) {
//	defer goleak.VerifyNone(t)
//	leak()
//}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}
