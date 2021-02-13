package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveNthNodeFromEndOfList(t *testing.T) {
	a := &ListNode{1, nil}
	b := &ListNode{2, nil}
	c := &ListNode{3, nil}
	d := &ListNode{4, nil}
	e := &ListNode{5, nil}
	f := &ListNode{1, nil}
	g := &ListNode{2, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	f.Next = g

	assert.Equal(t, []int{1, 2, 3, 5}, convertListToArray(removeNthFromEnd(a, 2)))
	assert.Equal(t, []int{2}, convertListToArray(removeNthFromEnd(f, 2)))
}
