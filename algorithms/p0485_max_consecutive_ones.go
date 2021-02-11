package algorithms

import "math"

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	countSoFar := 0

	for _, n := range nums {
		if n == 0 {
			maxCount = int(math.Max(float64(maxCount), float64(countSoFar)))
			countSoFar = 0
		} else {
			countSoFar += 1
		}
	}

	maxCount = int(math.Max(float64(maxCount), float64(countSoFar)))

	return maxCount
}
