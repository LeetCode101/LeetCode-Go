package algorithms

import "testing"
import "github.com/stretchr/testify/assert"

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	assert.Equal(t, []int{7, 0, 8}, convertListToArray(addTwoNumbers(l1, l2)))
}
