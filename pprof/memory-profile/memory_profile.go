package main

import (
	"github.com/pkg/profile"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func concat(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += randomString(n)
	}
	return s
}

// 内存性能分析
// func main() {
// 	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
// 	concat(100)
// }

func main() {
	defer profile.Start().Stop()
	concat(100)
}