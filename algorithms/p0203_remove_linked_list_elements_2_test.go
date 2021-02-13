package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveLinkedListElements2(t *testing.T) {
	a := &ListNode{1, nil}
	b := &ListNode{2, nil}
	c := &ListNode{6, nil}
	d := &ListNode{3, nil}
	e := &ListNode{4, nil}
	f := &ListNode{5, nil}
	g := &ListNode{6, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = g

	assert.Equal(t, []int{1, 2, 3, 4, 5}, convertListToArray(removeElements2(a, 6)))
}
