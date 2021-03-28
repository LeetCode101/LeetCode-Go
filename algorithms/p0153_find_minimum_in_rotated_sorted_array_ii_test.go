package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMinimumInRotatedSortedArrayII(t *testing.T) {
	assert.Equal(t, 1, findMinII([]int{1, 3, 5}))
	assert.Equal(t, 0, findMinII([]int{2, 2, 2, 0, 1}))
}
