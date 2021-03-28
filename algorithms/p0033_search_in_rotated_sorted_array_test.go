package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	assert.Equal(t, -1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}
