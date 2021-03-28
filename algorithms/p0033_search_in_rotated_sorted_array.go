package algorithms

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		middle := low + (high-low)/2

		if nums[middle] > target {
			if nums[middle] < nums[high] || nums[high] < target {
				high = middle - 1
			} else {
				low = middle + 1
			}
		} else if nums[middle] < target {
			if nums[middle] >= nums[high] || nums[high] >= target {
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
