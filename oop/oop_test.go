package oop

import (
	"fmt"
	"testing"
)

// 正方形
type Square struct {
	side float32
}

// 长方形
type Rectangle struct {
	length, width float32
}

// 接口 Shaper
type Shaper interface {
	Area() float32
}

// 计算正方形的面积
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

// 计算长方形的面积
func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

func TestPolymorphic(t *testing.T) {
	r := &Rectangle{10, 2}
	q := &Square{10}

	// 创建一个 Shaper 类型的数组
	shapes := []Shaper{r, q}
	// 迭代数组上的每一个元素并调用 Area() 方法
	for n, _ := range shapes {
		fmt.Println("图形数据: ", shapes[n])
		fmt.Println("它的面积是: ", shapes[n].Area())
	}
}
