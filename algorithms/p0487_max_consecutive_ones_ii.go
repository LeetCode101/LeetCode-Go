package algorithms

import "math"

func findMaxConsecutiveOnesII(nums []int) int {
	maxCount := 0
	countSoFar := 0
	prevCount := 0

	for _, n := range nums {
		if n == 0 {
			prevCount = countSoFar + 1
			countSoFar = 0
		} else {
			countSoFar += 1
		}

		maxCount = int(math.Max(float64(maxCount), float64(countSoFar+prevCount)))
	}

	return maxCount
}
