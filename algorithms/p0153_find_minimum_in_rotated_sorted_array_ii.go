package algorithms

func findMinII(nums []int) int {
	low := 0
	high := len(nums) - 1

	for low < high {
		middle := low + (high-low)/2

		if nums[middle] > nums[high] {
			low = middle + 1
		} else if nums[middle] < nums[high] {
			high = middle
		} else {
			high -= 1
		}
	}

	return nums[high]
}
