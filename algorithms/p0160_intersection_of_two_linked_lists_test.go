package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersectionOfTwoLinkedLists(t *testing.T) {
	a := &ListNode{4, nil}
	b := &ListNode{1, nil}
	c := &ListNode{8, nil}
	d := &ListNode{4, nil}
	e := &ListNode{5, nil}
	f := &ListNode{5, nil}
	g := &ListNode{6, nil}
	h := &ListNode{1, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	f.Next = g
	g.Next = h
	h.Next = c

	assert.Equal(t, 8, getIntersectionNode(a, f).Val)
}
