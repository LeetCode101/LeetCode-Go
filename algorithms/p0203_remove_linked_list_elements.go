package algorithms

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{-1, head}
	prev := dummy
	current := head

	for current != nil {
		if current.Val != val {
			prev.Next = current
			prev = current
		} else {
			prev.Next = current.Next
		}

		current = current.Next
	}

	return dummy.Next
}
