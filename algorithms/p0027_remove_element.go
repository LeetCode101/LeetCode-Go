package algorithms

func removeElement(nums []int, val int) int {
	tail := 0

	for _, n := range nums {
		if n != val {
			nums[tail] = n
			tail += 1
		}
	}

	return tail
}
