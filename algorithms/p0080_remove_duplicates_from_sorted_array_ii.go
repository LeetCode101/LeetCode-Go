package algorithms

func removeDuplicatesII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tail := 0
	duplicates := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[tail] {
			duplicates += 1
		} else {
			duplicates = 1
		}

		if duplicates <= 2 {
			tail += 1
			nums[tail] = nums[i]
		}
	}

	return tail + 1
}
