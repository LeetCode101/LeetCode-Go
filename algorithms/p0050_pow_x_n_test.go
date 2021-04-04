package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPow(t *testing.T) {
	assert.Equal(t, 1024.0, myPow(2, 10))
}
