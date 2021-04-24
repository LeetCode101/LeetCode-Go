package algorithms

import (
	"github.com/stretchr/testify/assert"
	"leetcode/algorithms/graph"
	"testing"
)

func TestCloneGraph2(t *testing.T) {
	a := &graph.Node{Val: 1}
	b := &graph.Node{Val: 2}
	c := &graph.Node{Val: 3}
	d := &graph.Node{Val: 4}
	a.Neighbors = []*graph.Node{b, d}
	b.Neighbors = []*graph.Node{a, c}
	c.Neighbors = []*graph.Node{b, d}
	d.Neighbors = []*graph.Node{a, c}

	assert.Nil(t, cloneGraph2(nil))
	assert.Equal(t, 1, cloneGraph2(a).Val)
}
