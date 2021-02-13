package algorithms

func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	head.Next = removeElements2(head.Next, val)

	if head.Val == val {
		return head.Next
	}

	return head
}
