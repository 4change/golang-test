package slice_test

import (
	"fmt"
	"testing"
)

type student struct {
	Name string
	Age  int
}

func GetSliceElementWrong() map[string]*student {
	m := make(map[string]*student)

	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	// 遍历slice时，不可以使用&stu的方式获取slice中的子元素
	// &stu表示获取stu这个变量的地址，而该变量的地址表示指向stus[i]的地址，故而该变量的地址会随着遍历过程的进行而改变
	// 该变量的地址在初始化时指向stus[0]的地址； 随着遍历过程的进行，该变量的地址会被修改为stus[i]的地址；最终该变量的地址会修改为stus[len - 1]的地址；
	// 所以通过&stu获取的将会是stus中最后一个子元素的地址
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	return m
}

func GetSliceElementWrongValidation() map[string]*student {
	m := make(map[string]*student)

	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	// &stu地址不变，始终指向stus[len - 1]元素的地址
	for _, stu := range stus {
		fmt.Printf("%v\t%p\n",stu,&stu)
		m[stu.Name] = &stu
	}

	return m
}

func GetSliceElementCorrect() map[string]*student {
	m := make(map[string]*student)

	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	for i, _ := range stus {
		stu := stus[i]
		m[stu.Name] = &stu
	}

	return m
}

func Test_Get_Slice_Element_Wrong(t *testing.T) {
	students := GetSliceElementWrong()
	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}

func Test_Get_Slice_Element_Wrong_Validation(t *testing.T) {
	students := GetSliceElementWrongValidation()
	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}

func Test_Get_Slice_Element_Correct(t *testing.T) {
	students := GetSliceElementCorrect()
	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}

type Stu struct {
	name string
	index int
}

func Test(t *testing.T)  {
	stus := []Stu{
		{"a", 1},
		{"b", 2},
	}
	stuMap := map[string]*Stu{}
	for _, stu := range stus {
		tmp := stu
		stuMap[stu.name] = &tmp

		stuMap[stu.name] = &stu
	}
	fmt.Println(stuMap)
}
