package algorithms

func detectCycle2(head *ListNode) *ListNode {
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
			break
		}
	}

	if fast == nil {
		return nil
	}

	p1 := head
	p2 := fast

	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}
