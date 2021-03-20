package algorithms

func convertListToArray(head *ListNode) []int {
	values := []int{}

	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	return values
}

func convertListToArray2(head *Node) []int {
	values := []int{}

	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	return values
}
