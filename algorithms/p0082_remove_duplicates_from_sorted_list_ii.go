package algorithms

func deleteDuplicatesII(head *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	dummy.Next = head
	current := head
	tail := dummy

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			for current != nil && current.Next != nil && current.Val == current.Next.Val {
				current = current.Next
			}

			if current != nil {
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
