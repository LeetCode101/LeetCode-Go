package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicatesFromSortedList1(t *testing.T) {
	head := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}

	assert.Equal(t, []int{}, convertListToArray(deleteDuplicates1(nil)))
	assert.Equal(t, []int{1, 2, 3}, convertListToArray(deleteDuplicates1(head)))
}
