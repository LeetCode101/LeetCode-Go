package algorithms

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	current := dummy
	current1 := l1
	current2 := l2

	for current1 != nil && current2 != nil {
		if current1.Val < current2.Val {
			current.Next = current1
			current1 = current1.Next
		} else {
			current.Next = current2
			current2 = current2.Next
		}

		current = current.Next
	}

	if current1 != nil {
		current.Next = current1
	}

	if current2 != nil {
		current.Next = current2
	}

	return dummy.Next
}
