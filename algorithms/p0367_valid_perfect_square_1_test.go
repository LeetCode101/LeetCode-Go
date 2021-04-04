package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidPerfectSquare(t *testing.T) {
	assert.True(t, isPerfectSquare(4))
}
