package adapter

import (
	"fmt"
)

// 适配器将要适配到的目标
// ForeignCenter 类并没有继承自Palyer接口，因为其attack()方法与Player中attack()方法签名不一致
type ForeignCenter struct {
	name string
}

func (f *ForeignCenter) attack(what string) {
	if f == nil {
		return
	}
	fmt.Println(f.name, what)
}
func (f *ForeignCenter) defense() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在防守")
}
