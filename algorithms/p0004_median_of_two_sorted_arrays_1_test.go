package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedianOfTwoSortedArrays1(t *testing.T) {
	assert.Equal(t, 2.5, findMedianSortedArrays1([]int{}, []int{2, 3}))
	assert.Equal(t, 2.0, findMedianSortedArrays1([]int{1, 3}, []int{2}))
}
