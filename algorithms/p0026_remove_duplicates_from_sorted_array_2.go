package algorithms

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tail := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[tail] {
			tail += 1
			nums[tail] = nums[i]
		}
	}

	return tail + 1
}
