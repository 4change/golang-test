package oop

import (
	"fmt"
	"testing"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func TestNilInterfaceInterview(t *testing.T) {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type EmptyInterface interface {}
type NotEmptyInterface interface {
	Show()
}

type EmptyInterfaceStruct struct {}
type NotEmptyInterfaceStruct struct {}
func (neis *NotEmptyInterfaceStruct) Show() {}

func ReturnEmptyInterface() EmptyInterface {
	var eis *EmptyInterfaceStruct
	return eis
}

func ReturnNotEmptyInterface() NotEmptyInterface {
	var neis *NotEmptyInterfaceStruct
	return neis
}

func TestNilInterface(t *testing.T) {
	var emptyInterface interface{}
	fmt.Println("emptyInterface == nil:", emptyInterface == nil)

	var notEmptyInterface interface{
		Show()
	}
	fmt.Println("notEmptyInterface == nil:", notEmptyInterface == nil)

	fmt.Println("ReturnEmptyInterface() == nil:", ReturnEmptyInterface() == nil)
	fmt.Println("ReturnNotEmptyInterface() == nil:", ReturnNotEmptyInterface() == nil)
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////