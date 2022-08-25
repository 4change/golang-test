package oop

import (
	"fmt"
	"testing"
)

type Human struct{}

func (p *Human) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *Human) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	Human
}

//func (t *Teacher) ShowA() {
//	fmt.Println("teacher showA")
//}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func TestStructExtends(t *testing.T) {
	teacher := Teacher{}
	// 当类型A组合进类型B时, 通过类型B调用方法method():
	//		优先类型B的方法method()
	//		若类型B中不包含方法method(), 则通过组合调用类型A的方法method()
	teacher.ShowA()
	fmt.Println("-------------------------------------------------------------------------------------------------")
	teacher.Human.ShowA()
}
