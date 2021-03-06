package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListCycleII2(t *testing.T) {
	a := &ListNode{3, nil}
	b := &ListNode{2, nil}
	c := &ListNode{0, nil}
	d := &ListNode{-4, nil}
	e := &ListNode{5, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = b

	assert.Equal(t, 2, detectCycle2(a).Val)
	assert.Nil(t, detectCycle2(e))
}
