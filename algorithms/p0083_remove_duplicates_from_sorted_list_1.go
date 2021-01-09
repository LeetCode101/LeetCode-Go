package algorithms

func deleteDuplicates1(head *ListNode) *ListNode {
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
			tail.Next = current.Next
		}

		current = current.Next
	}

	return head
}
