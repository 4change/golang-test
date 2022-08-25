package oop

import (
	"fmt"
	"testing"
)

type Set map[string]struct{}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}

func (s Set) Delete(key string) {
	delete(s, key)
}

func TestNilStructSet(t *testing.T) {
	s := make(Set)
	s.Add("Tom")
	s.Add("Sam")
	fmt.Println(s.Has("Tom"))
	fmt.Println(s.Has("Jack"))
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func worker(ch chan struct{}) {
	<-ch
	fmt.Println("do something")
	close(ch)
}

func TestNilStructChannel(t *testing.T) {
	ch := make(chan struct{})
	go worker(ch)
	ch <- struct{}{}		// 向 ch 传递一个空结构体，用作控制信号
}

type Door struct{}

func (d Door) Open() {
	fmt.Println("Open the door")
}

func (d Door) Close() {
	fmt.Println("Close the door")
}