package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSumII(t *testing.T) {
	numbers := []int{2, 7, 11, 15}

	assert.Equal(t, []int{1, 2}, twoSumII(numbers, 9))
	assert.Equal(t, []int{-1, -1}, twoSumII(numbers, 10))
}
