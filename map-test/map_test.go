package map_test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

// Go Map的遍历
//		1. Go中map的遍历是无序的
//		2. 同时, Go运行时也不会维护map元素的遍历顺序
func TestMapLoop(t *testing.T) {
	myMap := make(map[string]string)

	myMap["A"] = "a"
	myMap["B"] = "b"
	myMap["C"] = "c"
	myMap["D"] = "d"
	myMap["E"] = "e"
	myMap["F"] = "f"
	myMap["G"] = "g"
	myMap["H"] = "h"
	myMap["I"] = "i"
	myMap["J"] = "j"

	// map的遍历是无序的
	fmt.Println("myMap Loop: print key----------------------------------------------------------------------------")
	for i := range myMap {
		fmt.Println("key:", i)
	}

	fmt.Println("myMap Loop: print key & value====================================================================")
	for i := range myMap {
		fmt.Println("key:", i, ", value:",  myMap[i])
	}

	fmt.Println("myMap Loop: print key & value********************************************************************")
	for i, v := range myMap {
		fmt.Println("key:", i, ", value:",  v)
	}
}

// map的并发问题: 并发写问题
// 		如果map由多协程同时写会出现fatal error: concurrent map writes的错误
func TestMap_Exception_1_Concurrent_Write(t *testing.T) {
	c := make(map[string]int)
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 1000000; j++ {
				c[fmt.Sprintf("%d", j)] = j
			}
		}()
	}
	time.Sleep(time.Second * 20)
}

// map的并发问题: 并发读-写问题
// 		如果map由多协程同时读和写会出现 fatal error:concurrent map read and map write的错误
func TestMap_Exception_2_Concurrent_Read_Write(t *testing.T) {
	c := make(map[string]int)
	go func() {//开一个协程写map
		for j := 0; j < 1000000; j++ {
			c[fmt.Sprintf("%d", j)] = j
		}
	}()
	go func() {    //开一个协程读map
		for j := 0; j < 1000000; j++ {
			fmt.Println(c[fmt.Sprintf("%d",j)])
		}
	}()

	time.Sleep(time.Second*20)
}

// 解决方案一: 基于Mutex(互斥锁)的并发安全的map
type ConcurrentSafeMapByMutex struct {
	Data map[string]string
	Lock sync.Mutex
}

func (d *ConcurrentSafeMapByMutex) Get(k string) string {
	d.Lock.Lock()
	defer d.Lock.Unlock()

	return d.Data[k]
}


func (d *ConcurrentSafeMapByMutex) Set(k,v string) {
	d.Lock.Lock()
	defer d.Lock.Unlock()

	d.Data[k]=v
}

func TestMap_ConcurrentSafeMapByMutex_Concurrent_Write(t *testing.T) {
	csmbm := &ConcurrentSafeMapByMutex{
		Data: make(map[string]string),
	}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				temp := strconv.Itoa(j)
				csmbm.Set(temp, temp)
			}
		}()
	}

	time.Sleep(time.Second * 5)

	fmt.Println("ConcurrentSafeMapByMutex Concurrent Write Success------------------------------------------------")
}

func TestMap_ConcurrentSafeMapByMutex_Concurrent_Read_Write(t *testing.T)  {
	csmbm := &ConcurrentSafeMapByMutex{
		Data: make(map[string]string),
	}

	go func() {
		for j := 0; j < 100; j++ {
			temp := strconv.Itoa(j)
			csmbm.Set(temp, temp)
		}
	}()

	go func() {
		for j := 0; j < 100; j++ {
			temp := strconv.Itoa(j)
			csmbm.Get(temp)
		}
	}()

	time.Sleep(time.Second * 5)

	fmt.Println("ConcurrentSafeMapByMutex Concurrent Read Write Success-------------------------------------------")
}

// 解决方案二: 基于RWMutex(读写锁)的并发安全的map
type ConcurrentSafeMapByRWMutex struct {
	Data map[string]string
	Lock sync.RWMutex
}

func NewConcurrentSafeMapByRWMutex() *ConcurrentSafeMapByRWMutex {
	csmbrwm := new(ConcurrentSafeMapByRWMutex)
	csmbrwm.Data = make(map[string]string)
	return csmbrwm
}

func (csmbrwm *ConcurrentSafeMapByRWMutex) Get(k string) string {
	csmbrwm.Lock.RLock()
	defer csmbrwm.Lock.RUnlock()

	return csmbrwm.Data[k]
}

func (csmbrwm *ConcurrentSafeMapByRWMutex) Set(k, v string) {
	csmbrwm.Lock.Lock()
	defer csmbrwm.Lock.Unlock()
	csmbrwm.Data[k] = v
}

func TestMap_ConcurrentSafeMapByRWMutex_Concurrent_Write(t *testing.T) {
	csmbrwm := NewConcurrentSafeMapByRWMutex()

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				temp := strconv.Itoa(j)
				csmbrwm.Set(temp, temp)
			}
		}()
	}

	time.Sleep(time.Second * 5)

	fmt.Println("ConcurrentSafeMapByRWMutex Concurrent Write Success----------------------------------------------")
}

func TestMap_ConcurrentSafeMapByRWMutex_Concurrent_Read_Write(t *testing.T)  {
	csmbrwm := NewConcurrentSafeMapByRWMutex()

	go func() {
		for j := 0; j < 100; j++ {
			temp := strconv.Itoa(j)
			csmbrwm.Set(temp, temp)
		}
	}()

	go func() {
		for j := 0; j < 100; j++ {
			temp := strconv.Itoa(j)
			csmbrwm.Get(temp)
		}
	}()

	time.Sleep(time.Second * 5)

	fmt.Println("ConcurrentSafeMapByRWMutex Concurrent Read Write Success-----------------------------------------")
}

type Info struct {
	age int
}
type AccountMap struct {
	accounts map[string]*Info
	ch chan func()							// 这里的定义是什么高级用法？
}
func NewAccountMap() *AccountMap {
	p := &AccountMap{
		accounts: make(map[string]*Info),
		ch: make(chan func()),
	}
	go func() {
		for {(<-p.ch)()}
	}()
	return p
}
func (p *AccountMap) add(name string, age int) {
	p.ch <- func() {
		p.accounts[name] = &Info{age}
	}
}
func (p *AccountMap) del(name string) {
	p.ch <- func() {
		delete(p.accounts, name)
	}
}
func (p *AccountMap) find(name string) *Info {
	// 每次查询都要创建一个信道
	c := make(chan *Info)
	p.ch <- func() {
		res, ok := p.accounts[name]
		if !ok {
			c <- nil
		} else {
			inf := *res
			c <- &inf
		}
	}
	return <-c
}

func TestAccountMap(t *testing.T)  {
	am := NewAccountMap()

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				temp := strconv.Itoa(j)
				am.add(temp, j)
			}
		}()
	}

	time.Sleep(time.Second * 5)

	fmt.Println("ConcurrentSafeMapByRWMutex Concurrent Write Success----------------------------------------------")
}