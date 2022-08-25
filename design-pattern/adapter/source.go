package adapter

import (
	"fmt"
)

type Player interface {
	attack()
	defense()
}

type Forwards struct {
	name string
}

// Forwards.attack()方法，Forwards类实现了Player接口，attack()方法实现自Player接口
func (f *Forwards) attack() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在进攻")
}

// Forwards.defense()方法，Forwards类实现了Player接口，defense()方法实现自Player接口
func (f *Forwards) defense() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在防守")
}

// Forwards类构造器
func NewForwards(name string) Player {
	return &Forwards{name}
}

type Centers struct {
	name string
}

// Centers.attack()方法，Centers类实现了Player接口，attack()方法实现自Player接口
func (f *Centers) attack() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在进攻")
}

// Centers.defense()方法，Centers类实现了Player接口，defense()方法实现自Player接口
func (f *Centers) defense() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在防守")
}

// NewCenter类构造器
func NewCenter(name string) Player {
	return &Centers{name}
}
