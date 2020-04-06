package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters2(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring2("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring2("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring2("pwwkew"))
}
