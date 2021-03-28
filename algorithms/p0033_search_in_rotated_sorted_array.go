package algorithms

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		middle := low + (high-low)/2
		value := nums[middle]

		if value > target {
			if value < nums[high] || nums[high] < target {
				high = middle - 1
			} else {
				low = middle + 1
			}
		} else if value < target {
			if value >= nums[high] || nums[high] >= target {
				low = middle + 1
			} else {
				high = middle - 1
			}
		} else {
			return middle
		}
	}

	return -1
}
