package algorithms

func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	current := root
	values := []int{}
	stack := []*TreeNode{}

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		values = append(values, current.Val)
		current = current.Right
	}

	return values
}
