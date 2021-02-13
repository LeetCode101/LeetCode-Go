package algorithms

func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{-1, head}
	tailOfLeft := dummy
	current := head

	for i := 0; i < m-1; i++ {
		tailOfLeft = current
		current = current.Next
	}

	tailOfReversed := current

	for i := 0; i < n-m; i++ {
		next := tailOfReversed.Next
		tailOfReversed.Next = next.Next
		temp := tailOfLeft.Next
		next.Next = temp
		tailOfLeft.Next = next
	}

	return dummy.Next
}
