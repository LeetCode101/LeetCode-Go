package algorithms

func detectCycle1(head *ListNode) *ListNode {
	visited := map[*ListNode]bool{}
	current := head

	for current != nil && current.Next != nil {
		if visited[current.Next] {
			return current.Next
		} else {
			visited[current] = true
			current = current.Next
		}
	}

	return nil
}
