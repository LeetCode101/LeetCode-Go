package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidPerfectSquare1(t *testing.T) {
	assert.True(t, isPerfectSquare1(4))
	assert.False(t, isPerfectSquare1(999))
}
