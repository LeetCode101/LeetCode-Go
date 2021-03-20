package algorithms

func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	current := head
	i := 0
	index2Node := make(map[int]*ListNode)

	for current != nil {
		i += 1
		index2Node[i] = current
		current = current.Next
	}

	step := k % i

	if step == 0 {
		return head
	}

	index2Cut := i - step
	index2Node[i].Next = head
	index2Node[index2Cut].Next = nil

	return index2Node[index2Cut+1]
}
