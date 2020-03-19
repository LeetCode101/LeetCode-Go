package algorithms

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	prev := dummy
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
		prev.Next = &ListNode{sum % 10, nil}
		prev = prev.Next
	}

	if carry != 0 {
		prev.Next = &ListNode{carry, nil}
	}

	return dummy.Next
}
