package main

import "fmt"

func main() {
	ob1 := &ConcreteObserver{username: "小明"}
	ob2 := &ConcreteObserver{username: "小汪"}
	sub := &ConcreteSubject{}
	sub.Notify() //do nothing
	fmt.Println("===============================================================================================")

	sub.Add(ob1)
	sub.Add(ob2)
	sub.Notify()
	fmt.Println("***********************************************************************************************")

	sub.Remove(ob1)
	sub.Notify()
}
