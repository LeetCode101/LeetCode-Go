package algorithms

import "testing"
import "github.com/stretchr/testify/assert"

func TestTwoSum(t *testing.T) {
	numbers := []int{2, 7, 11, 15}

	assert.Equal(t, []int{0, 1}, twoSum(numbers, 9))
	assert.Equal(t, []int{-1, -1}, twoSum(numbers, 10))
}
