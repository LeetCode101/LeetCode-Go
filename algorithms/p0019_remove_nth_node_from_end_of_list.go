package algorithms

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	step := 0
	slow := head
	fast := head

	for fast.Next != nil {
		if step >= n {
			slow = slow.Next
		}

		step += 1
		fast = fast.Next
	}

	if step == n-1 {
		return head.Next
	}

	slow.Next = slow.Next.Next

	return head
}
