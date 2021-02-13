package algorithms

func reverseList2(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func reverse(prev *ListNode, current *ListNode) *ListNode {
	if current == nil {
		return prev
	}

	next := current.Next
	current.Next = prev

	return reverse(current, next)
}
