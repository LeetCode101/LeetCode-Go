package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicatesFromSortedListII(t *testing.T) {
	head := &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}}

	assert.Equal(t, []int{2, 3}, convertListToArray(deleteDuplicatesII(head)))
}
