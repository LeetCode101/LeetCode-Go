package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotateList(t *testing.T) {
	a := &ListNode{1, nil}
	b := &ListNode{2, nil}
	c := &ListNode{3, nil}
	d := &ListNode{4, nil}
	e := &ListNode{5, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e

	assert.Nil(t, rotateRight(nil, 1))
	assert.Equal(t, a, rotateRight(a, 5))
	assert.Equal(t, []int{4, 5, 1, 2, 3}, convertListToArray(rotateRight(a, 2)))
}
