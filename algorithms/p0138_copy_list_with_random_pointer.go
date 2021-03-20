package algorithms

func copyRandomList(head *Node) *Node {
	new2Old := make(map[*Node]*Node)
	old2New := make(map[*Node]*Node)
	dummy := &Node{Val: 0}
	current := head
	newCurrent := dummy

	for current != nil {
		newCurrent.Next = &Node{Val: current.Val}
		newCurrent = newCurrent.Next
		new2Old[newCurrent] = current
		old2New[current] = newCurrent
		current = current.Next
	}

	newCurrent = dummy.Next

	for newCurrent != nil {
		newCurrent.Random = old2New[new2Old[newCurrent].Random]
		newCurrent = newCurrent.Next
	}

	return dummy.Next
}
