package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListCycle(t *testing.T) {
	a := &ListNode{3, nil}
	b := &ListNode{2, nil}
	c := &ListNode{0, nil}
	d := &ListNode{-4, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = b

	assert.True(t, hasCycle(a))
}
