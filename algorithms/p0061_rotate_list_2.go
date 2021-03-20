package algorithms

func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	current := head
	i := 1
	index2Node := make(map[int]*ListNode)

	for current.Next != nil {
		i += 1
		index2Node[i] = current
		current = current.Next
	}

	step := i - k%i - 1
	current.Next = head
	current = head

	for i := step; i > 0; i-- {
		current = current.Next
	}

	newHead := current.Next
	current.Next = nil

	return newHead
}
