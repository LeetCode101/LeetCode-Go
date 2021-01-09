package algorithms

func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head.Next = deleteDuplicates3(head.Next)

	if head.Val != head.Next.Val {
		return head
	}

	return head.Next
}
