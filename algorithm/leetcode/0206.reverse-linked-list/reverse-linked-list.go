package problem0206

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}

	return prev
}

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}