package algorithms

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{-1, head}
	tailOfLeft := dummy
	current := head

	for i := 0; i < m-1; i++ {
		tailOfLeft = current
		current = current.Next
	}

	tailOfReversed := current
	var prev *ListNode
	var next *ListNode

	for i := 0; i < n-m+1; i++ {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	tailOfLeft.Next = prev
	tailOfReversed.Next = current

	return dummy.Next
}
