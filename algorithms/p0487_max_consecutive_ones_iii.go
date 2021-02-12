package algorithms

import "math"

func longestOnes(A []int, K int) int {
	zeros := 0
	maxCount := 0
	windowStart := 0

	for j, n := range A {
		if n == 0 {
			zeros += 1
		}

		for zeros > K {
			if A[windowStart] == 0 {
				zeros -= 1
			}

			windowStart += 1
		}

		maxCount = int(math.Max(float64(maxCount), float64(j-windowStart+1)))
	}

	return maxCount
}
