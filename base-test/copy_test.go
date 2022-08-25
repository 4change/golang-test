package base_test

import (
	"fmt"
	"testing"
)

// 复制Slice_方案一: 复制所有数据
func TestCopySlice_CopyAll(t *testing.T) {
	a := []int{1, 2}
	b := []int{3, 4}
	check := a
	copy(a, b)
	fmt.Println(a, b, check)
	// Output: [3 4] [3 4] [3 4]
}

func TestCopySlice_CopyValue(t *testing.T) {
	a := []int{1, 2}
	b := []int{3, 4}
	check := a
	a = b
	fmt.Println(a, b, check)
	// Output: [3 4] [3 4] [1 2]
}

func TestCopyMap_CopyAll(t *testing.T) {
	a := map[string]bool{"A": true, "B": true}
	check := a
	b := make(map[string]bool)
	for key, value := range a {
		b[key] = value
	}

	fmt.Println(a, b, check)
	// Output: map[A:true B:true] map[A:true B:true] map[A:true B:true]
}

func TestCopyMap_CopyDesc(t *testing.T) {
	a := map[string]bool{"A": true, "B": true}
	b := map[string]bool{"C": true, "D": true}
	check := a
	a = b
	fmt.Println(a, b, check)
	// Output: map[C:true D:true] map[C:true D:true] map[A:true B:true]
}