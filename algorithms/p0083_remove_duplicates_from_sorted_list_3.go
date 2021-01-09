package algorithms

func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	head.Next = deleteDuplicates3(head.Next)

	if head.Next != nil && head.Val == head.Next.Val {
		return head.Next
	}

	return head
}
