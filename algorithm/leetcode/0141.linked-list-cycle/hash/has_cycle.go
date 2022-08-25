package hash

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	myMap := make(map[*ListNode]int)

	for head != nil {
		if _, ok := myMap[head]; ok {
			return true
		} else {
			myMap[head] = 1
		}

		head = head.Next
	}

	return false
}
