package once

import (
	"fmt"
	"sync"
)

type singleton struct {
	data int
}

// 单例对象设为私有, 禁止外部访问
var sin *singleton

// 同步Once, 保证单例对象创建函数只执行一次
var once sync.Once

// 获取实例对象函数
func GetSingleton() *singleton {
	once.Do(func() {
		sin = &singleton{12}
			fmt.Println("实例对象的信息和地址", sin, &sin)
	})
	return sin
}
