package main

import "fmt"

//Observer　观察者
type Observer interface {
	Response(msg string)
}

// =====================================================================================================================

//ConcreteObserver　实例化的观察者
type ConcreteObserver struct {
	username string
}

func (t *ConcreteObserver) Response(msg string) {
	fmt.Println(t.username, ":", msg)
}
