package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThirdMaximumNumber(t *testing.T) {
	assert.Equal(t, 1, thirdMax([]int{3, 2, 1}))
}
