package practice

import "sync"

type singleton struct {
	data int
}

var instance *singleton

var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{data: 1}
	})
	return instance
}
