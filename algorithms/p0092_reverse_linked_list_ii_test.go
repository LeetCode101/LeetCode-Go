package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseLinkedListII(t *testing.T) {
	a := &ListNode{1, nil}
	b := &ListNode{2, nil}
	c := &ListNode{3, nil}
	d := &ListNode{4, nil}
	e := &ListNode{5, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e

	assert.Equal(t, []int{1, 4, 3, 2, 5}, convertListToArray(reverseBetween(a, 2, 4)))
}
