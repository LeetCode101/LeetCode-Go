package algorithms

func removeDuplicates(nums []int) int {
	tail := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[tail] {
			tail += 1
			nums[tail] = nums[i]
		}
	}

	if len(nums) == 0 {
		return 0
	} else {
		return tail + 1
	}
}
