package algorithms

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	tail := head
	current := head.Next

	for current != nil {
		if tail.Val != current.Val {
			tail.Next = current
			tail = current
		} else {
			tail.Next = nil
		}

		current = current.Next
	}

	return head
}
