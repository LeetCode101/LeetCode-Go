package algorithms

import "math"

func findMaxConsecutiveOnes2(nums []int) int {
	maxCount := 0
	countSoFar := 0

	for _, n := range nums {
		countSoFar *= n
		countSoFar += n
		maxCount = int(math.Max(float64(maxCount), float64(countSoFar)))
	}

	return maxCount
}
