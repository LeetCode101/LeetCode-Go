package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPow2(t *testing.T) {
	assert.Equal(t, 1.0, myPow2(1, 0))
	assert.Equal(t, 0.25, myPow2(2, -2))
	assert.Equal(t, 1024.0, myPow2(2, 10))
}
