package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPow1(t *testing.T) {
	assert.Equal(t, 1.0, myPow1(1, 0))
	assert.Equal(t, 0.25, myPow1(2, -2))
	assert.Equal(t, 1024.0, myPow1(2, 10))
}
