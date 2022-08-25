package oop

import (
	"fmt"
	"testing"
)

type student struct{
    name string
    age int
}
func (stu student) speak1(){
    fmt.Println("I am a student1, I am ", stu.age)
}
func (stu *student) speak2(){
    fmt.Println("I am a student2, I am ", stu.age)
}

func TestStructPointer(t *testing.T){
    var s1 student
    var s2 = s1         // 两种声明方式结果一样
    // var s2 = &s1     // 两种声明方式结果一样
    s1.speak1()
    s1.speak2()
    s2.speak1()
    s2.speak2()
    //var p people
    //p = student{name:"RyuGou", age:12} //应该改为 p = &student{name:"RyuGou", age:12}
    //p.speak()
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// type people interface {
//     speak1()
//     speak2()
// }

// type student struct{
//     name string
//     age int
// }
// func (stu student) speak1(){
//     fmt.Println("I am a student1, I am ", stu.age)
// }
// func (stu *student) speak2(){
//     fmt.Println("I am a student2, I am ", stu.age)
// }

// func TestInterfacePointer(t *testing.T){
//     var p people
//     p = &student{name:"RyuGou", age:12} // 如果不是用指针会报错
//     // p = student{name:"RyuGou", age:12} // 如果不是用指针会报错
//     p.speak1()
//     p.speak2()
// }