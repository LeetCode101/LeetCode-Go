package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyListWithRandomPointer(t *testing.T) {
	a := &Node{Val: 7}
	b := &Node{Val: 13}
	c := &Node{Val: 11}
	d := &Node{Val: 10}
	e := &Node{Val: 1}
	a.Next = b
	b.Next = c
	b.Random = a
	c.Next = d
	c.Random = e
	d.Random = c
	d.Next = e
	e.Random = a

	assert.Equal(t, []int{7, 13, 11, 10, 1}, convertListToArray2(copyRandomList(a)))
}
