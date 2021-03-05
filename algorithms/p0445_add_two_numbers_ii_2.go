package algorithms

func addTwoNumbersII2(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := convertToStack(l1)
	stack2 := convertToStack(l2)
	size1 := len(stack1)
	size2 := len(stack2)

	dummy := &ListNode{-1, nil}
	carry := 0

	for size1 > 0 || size2 > 0 {
		sum := carry

		if size1 > 0 {
			sum += stack1[size1-1]
			size1--
		}

		if size2 > 0 {
			sum += stack2[size2-1]
			size2--
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

func convertToStack(head *ListNode) []int {
	var stack []int
	current := head

	for current != nil {
		stack = append(stack, current.Val)
		current = current.Next
	}

	return stack
}
