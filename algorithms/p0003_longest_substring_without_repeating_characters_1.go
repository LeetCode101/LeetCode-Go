package algorithms

func lengthOfLongestSubstring1(s string) int {
	i := 0
	maxLength := 0
	mapping := make(map[int32]int)

	for j, c := range s {
		if index, ok := mapping[c]; ok {
			for k := i; k <= index; k++ {
				delete(mapping, int32(s[k]))
			}

			i = index + 1
		}

		if j-i+1 > maxLength {
			maxLength = j - i + 1
		}

		mapping[c] = j
	}

	return maxLength
}
