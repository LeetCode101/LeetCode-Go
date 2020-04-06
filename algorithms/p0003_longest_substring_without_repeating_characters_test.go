package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
}
