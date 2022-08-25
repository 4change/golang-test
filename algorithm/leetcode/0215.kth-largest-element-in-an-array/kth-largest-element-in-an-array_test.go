package problem0215

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int			// 待测试数组
	k    int			// 数组中第K大的元素
	ans  int			// 期望测试答案
}{

	{
		[]int{3, 3, 3, 3, 3, 3, 3, 3, 3},
		1,
		3,
	},

	{
		[]int{3, 2, 1, 5, 6, 4},
		2,
		5,
	},

	// 可以有多个 testcase
}

func Test_findKthLargest(t *testing.T) {
	ast := assert.New(t)			// 新建一个测试断言对象

	// 对tcs中的测试案例进行测试
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findKthLargest(tc.nums, tc.k), "输入:%v", tc)
	}
}

func Test_heap(t *testing.T) {
	ast := assert.New(t)			// 新建一个测试断言对象

	h := new(highHeap)				// 新建一个heap

	i := 5
	h.Push(i)
	ast.Equal(i, h.Pop(), "Pop() after Push(%d)", i)
}

func Benchmark_findKthLargest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findKthLargest(tc.nums, tc.k)
		}
	}
}
