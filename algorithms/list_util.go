package algorithms

func convertListToArray(head *ListNode) []int {
	var values []int

	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	return values
}
