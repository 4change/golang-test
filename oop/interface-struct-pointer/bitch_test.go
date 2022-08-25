package oop

import (
	"fmt"
	"testing"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

//func (stu Stduent) Speak(think string) (talk string) {
func (stu Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func TestPointer(t *testing.T) {
	//var peo People = Stduent{}		// 错误, Student{} 未实现 Speak() 方法，不能算作 People 的子类
	var peo People = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
