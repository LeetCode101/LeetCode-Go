package algorithms

func inorderTraversal1(root *TreeNode) []int {
	return inorderTraversalInternal(root, []int{})
}

func inorderTraversalInternal(root *TreeNode, values []int) []int {
	if root == nil {
		return values
	}

	values = inorderTraversalInternal(root.Left, values)
	values = append(values, root.Val)
	values = inorderTraversalInternal(root.Right, values)

	return values
}
