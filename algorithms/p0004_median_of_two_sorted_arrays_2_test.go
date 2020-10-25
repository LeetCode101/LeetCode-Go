package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedianOfTwoSortedArrays2(t *testing.T) {
	assert.Equal(t, 2.5, findMedianSortedArrays2([]int{}, []int{2, 3}))
	assert.Equal(t, 2.0, findMedianSortedArrays2([]int{1, 3}, []int{2}))
	assert.Equal(t, 2.0, findMedianSortedArrays2([]int{2, 3}, []int{1}))
}
