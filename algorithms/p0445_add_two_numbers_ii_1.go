package algorithms

func addTwoNumbersII1(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)

	dummy := &ListNode{-1, nil}
	carry := 0

	for l1 != nil || l2 != nil {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current := &ListNode{sum % 10, nil}
		current.Next = dummy.Next
		dummy.Next = current
	}

	if carry != 0 {
		current := &ListNode{carry, nil}
		current.Next = dummy.Next
		dummy.Next = current
	}

	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
