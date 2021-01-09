package algorithms

func deleteDuplicatesII(head *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	dummy.Next = head
	current := head
	tail := dummy

	for current != nil && current.Next != nil {
		currentValue := current.Val

		if current.Next.Val == currentValue {
			for current != nil && current.Val == currentValue {
				current = current.Next
			}

			tail.Next = current
		} else {
			tail.Next = current
			tail = current
			current = current.Next
		}
	}

	return dummy.Next
}
