package problem0206

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []int
}

func Test_Problem0206(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{[]int{1, 2, 3, 4, 5}},
			ans{[]int{5, 4, 3, 2, 1}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, l2s(reverseList(s2l(p.one))), "输入:%v", p)
	}
}

// convert *ListNode to []int
func l2s(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// convert []int to *ListNode
func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	// 链表头元素
	res := &ListNode{
		Val: nums[0],
	}

	// 尾插法构造链表
	temp := res
	for i := 1; i < len(nums); i++ {
		// 构造链表元素
		temp.Next = &ListNode{
			Val: nums[i],
		}
		// 尾插法插入链表元素
		temp = temp.Next
	}

	return res
}
