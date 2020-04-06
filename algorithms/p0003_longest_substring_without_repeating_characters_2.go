package algorithms

func lengthOfLongestSubstring2(s string) int {
	i := 0
	maxLength := 0
	mapping := make(map[int32]int)

	for j, c := range s {
		if index, ok := mapping[c]; ok && index+1 > i {
			i = index + 1
		}

		if j-i+1 > maxLength {
			maxLength = j - i + 1
		}

		mapping[c] = j
	}

	return maxLength
}
