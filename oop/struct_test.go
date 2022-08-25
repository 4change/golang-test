package oop

import (
	"fmt"
	"io/ioutil"
	"sort"
	"testing"
	"time"
	"unsafe"
)

func TestEmptyStruct(t *testing.T) {
	a := struct {}{}
	println(unsafe.Sizeof(a))
}

func TestEmptyStructWithMap(t *testing.T) {
	set := make(map[string]struct{})
	for _, value := range []string{"apple", "orange", "apple"} {
		set[value] = struct{}{}
	}
	fmt.Println(set)
}

func Test_Compare_Struct(t *testing.T) {
	// 比较两个结构相同且不包含slice, map, func的结构体: 可以正常进行比较
	type Foo struct {
		A int
		B string
		C interface{}
	}
	a := Foo{A: 1, B: "one", C: "two"}
	b := Foo{A: 1, B: "one", C: "two"}

	println(a == b)
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// 比较两个结构相同但包含slice, map, func的结构体: 不可以正常进行比较, 代码无法通过编译
	//type Bar struct {
	//	M []int
	//}
	//x := Bar{M: []int{1}}
	//y := Bar{M: []int{1}}
	//
	//println(x == y)
}

type info struct {
	Name string
	Time time.Time
}
type newlist []*info

func TestSort(t *testing.T) {
	l, e := getFilelist("./")
	if e != nil {
		fmt.Println(e)
	}
	sort.Sort(newlist(l))  //调用标准库的sort.Sort必须要先实现Len(),Less(),Swap() 三个方法.
	for _, v := range l {
		fmt.Println("文件名：", v.Name, "修改时间：", v.Time.Unix())
	}
}

func getFilelist(path string) ([]*info, error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return []*info{}, err
	}
	var list []*info
	for _, v := range l {
		list = append(list, &info{v.Name(), v.ModTime()})
	}
	return list, nil
}

func (I newlist) Len() int {
	return len(I)
}
func (I newlist) Less(i, j int) bool {
	return I[i].Time.Unix() < I[j].Time.Unix()
}
func (I newlist) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Person struct {
	Name string
	Age int
}
type PersonSlice []*Person

func TestPersonSort(t *testing.T) {
	p1 := &Person{
		Name:"C",
		Age:10,
	}
	p2 := &Person{
		Name:"D",
		Age:11,
	}
	p3 := &Person{
		Name:"B",
		Age:12,
	}

	var ps []*Person

	ps = append(ps, p1)
	ps = append(ps, p2)
	ps = append(ps, p3)

	sort.Sort(PersonSlice(ps))  //调用标准库的sort.Sort必须要先实现Len(),Less(),Swap() 三个方法.

	for _, p := range ps {
		fmt.Println(p.Name)
	}
}

func (ps PersonSlice) Len() int {
	return len(ps)
}
func (ps PersonSlice) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}
func (ps PersonSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}