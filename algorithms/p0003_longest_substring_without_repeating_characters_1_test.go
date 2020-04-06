package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters1(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring1("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring1("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring1("pwwkew"))
}
