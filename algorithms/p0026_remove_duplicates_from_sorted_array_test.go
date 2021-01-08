package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	assert.Equal(t, 0, removeDuplicates([]int{}))
	assert.Equal(t, 2, removeDuplicates([]int{1, 1, 2}))
}
