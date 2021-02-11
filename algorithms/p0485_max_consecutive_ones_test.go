package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxConsecutiveOnes(t *testing.T) {
	assert.Equal(t, 3, findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
