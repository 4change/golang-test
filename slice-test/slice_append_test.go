package slice_test

import (
	"fmt"
	"testing"
)

func TestSliceAppend(t *testing.T) {
	s := make([]int, 5)
	// 向[]int slice中追加1, 2, 3这些元素
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	sublist := s[0:0]
	fmt.Println(sublist)
}

func TestNil(t *testing.T) {
	var slice []map[string]interface{}

	slice = append(slice, nil,	nil)

	fmt.Println(len(slice), "-----------", cap(slice))
}