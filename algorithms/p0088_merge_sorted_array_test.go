package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	actualList := []int{1, 2, 3, 0, 0, 0}
	merge(actualList, 3, []int{2, 5, 6}, 3)

	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, actualList)
}
