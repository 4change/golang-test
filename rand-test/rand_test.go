package rand_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	randAgent := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := randAgent.Intn(100)
	if randNum < 0 {
		fmt.Println(true)
	}else {
		fmt.Println(false)
	}
}
