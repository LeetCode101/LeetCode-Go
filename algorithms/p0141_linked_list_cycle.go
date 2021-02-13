package algorithms

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for slow != nil && fast != nil {
		slow = slow.Next

		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
		}

		if slow == fast && slow != nil {
			return true
		}
	}

	return false
}
