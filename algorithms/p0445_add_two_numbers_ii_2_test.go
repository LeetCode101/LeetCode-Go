package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbersII2(t *testing.T) {
	l1 := &ListNode{9, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}
	l2 := &ListNode{9, &ListNode{6, &ListNode{4, nil}}}

	assert.Equal(t, []int{1, 0, 2, 0, 7}, convertListToArray(addTwoNumbersII2(l1, l2)))
}
