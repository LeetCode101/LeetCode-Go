package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqrtX(t *testing.T) {
	assert.Equal(t, 2, mySqrt(4))
	assert.Equal(t, 2, mySqrt(8))
}
