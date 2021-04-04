package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPow1(t *testing.T) {
	assert.Equal(t, 1024.0, myPow1(2, 10))
}
