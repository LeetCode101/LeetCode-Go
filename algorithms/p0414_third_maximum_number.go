package algorithms

import "math"

func thirdMax(nums []int) int {
	max1 := math.MinInt64
	max2 := math.MinInt64
	max3 := math.MinInt64

	for _, n := range nums {
		if n > max1 {
			max3 = max2
			max2 = max1
			max1 = n
		} else if n > max2 && n < max1 {
			max3 = max2
			max2 = n
		} else if n > max3 && n < max2 {
			max3 = n
		}
	}

	if max3 != math.MinInt64 {
		return max3
	}

	return max1
}
