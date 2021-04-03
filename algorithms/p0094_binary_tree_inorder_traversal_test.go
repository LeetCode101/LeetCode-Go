package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeInorderTraversal(t *testing.T) {
	a := &TreeNode{Val: 4}
	b := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	c := &TreeNode{Val: 5}
	a.Left = b
	a.Right = c

	assert.Equal(t, []int{1, 2, 3, 4, 5}, inorderTraversal(a))
}
