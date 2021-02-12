package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquaresOfASortedArray(t *testing.T) {
	assert.Equal(t, []int{0, 1, 9, 16, 100}, sortedSquares([]int{-4, -1, 0, 3, 10}))
}
