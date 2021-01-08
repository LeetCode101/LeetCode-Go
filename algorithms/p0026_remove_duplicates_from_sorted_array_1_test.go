package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicatesFromSortedArray1(t *testing.T) {
	assert.Equal(t, 0, removeDuplicates1([]int{}))
	assert.Equal(t, 2, removeDuplicates1([]int{1, 1, 2}))
}
