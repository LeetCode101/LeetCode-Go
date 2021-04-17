package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfIslands2(t *testing.T) {
	grid := [][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}

	assert.Equal(t, 0, numIslands2(nil))
	assert.Equal(t, 3, numIslands2(grid))
}
