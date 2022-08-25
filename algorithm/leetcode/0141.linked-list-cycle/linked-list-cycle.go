package problem0141

func hasCycle(head *ListNode) bool {
	// 链表是否为空的校验
	if head == nil {
		return false
	}

	// 快慢指针的分配
	slow, fast := head, head.Next

	// 链表有环时, slow = fast 会导致循环退出
	// 链表无环时, fast = nil 或 fast.Next = nil 会导致循环退出
	for fast != nil && fast.Next != nil && slow != fast {
		slow, fast = slow.Next, fast.Next.Next
	}

	// 通过快慢指针是否相等判定链表是否有环
	return slow == fast
}
