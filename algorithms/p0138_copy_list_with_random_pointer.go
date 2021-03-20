package algorithms

func copyRandomList(head *Node) *Node {
	old2New := make(map[*Node]*Node)
	dummy := &Node{Val: 0}
	current := head
	newCurrent := dummy

	for current != nil {
		newCurrent.Next = &Node{Val: current.Val}
		newCurrent = newCurrent.Next
		old2New[current] = newCurrent
		current = current.Next
	}

	current = head
	newCurrent = dummy.Next

	for current != nil {
		if current.Random != nil {
			newCurrent.Random = old2New[current.Random]
		}

		current = current.Next
		newCurrent = newCurrent.Next
	}

	return dummy.Next
}
