package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicatesFromSortedArrayII(t *testing.T) {
	assert.Equal(t, 0, removeDuplicatesII([]int{}))
	assert.Equal(t, 5, removeDuplicatesII([]int{1, 1, 1, 2, 2, 3}))
	assert.Equal(t, 7, removeDuplicatesII([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}
