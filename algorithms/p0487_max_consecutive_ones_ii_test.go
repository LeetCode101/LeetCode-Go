package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxConsecutiveOnesII(t *testing.T) {
	assert.Equal(t, 4, findMaxConsecutiveOnesII([]int{1, 0, 1, 1, 0}))
}
