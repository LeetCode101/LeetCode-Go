package algorithms

func removeDuplicates1(nums []int) int {
	length := len(nums)
	size := 0

	for i := 0; i < length; {
		j := i + 1

		for j < length && nums[j] == nums[i] {
			j += 1
		}

		nums[size] = nums[i]
		size += 1
		i = j
	}

	return size
}
