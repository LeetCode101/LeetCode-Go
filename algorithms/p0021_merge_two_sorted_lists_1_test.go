package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoSortedLists1(t *testing.T) {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	assert.Equal(t, []int{1, 1, 2, 3, 4, 4}, convertListToArray(mergeTwoLists1(l1, l2)))
	assert.Equal(t, []int{1, 2, 3}, convertListToArray(mergeTwoLists1(&ListNode{1, nil}, &ListNode{2, &ListNode{3, nil}})))
}
