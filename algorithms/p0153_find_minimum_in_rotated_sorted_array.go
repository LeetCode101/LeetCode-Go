package algorithms

func findMin(nums []int) int {
	low := 0
	high := len(nums) - 1

	for low < high {
		middle := low + (high-low)/2

		if nums[middle] > nums[high] {
			low = middle + 1
		} else {
			high = middle
		}
	}

	return nums[high]
}
