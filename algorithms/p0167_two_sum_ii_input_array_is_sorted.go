package algorithms

func twoSumII(numbers []int, target int) []int {
	low := 0
	high := len(numbers) - 1

	for low < high {
		current := numbers[low] + numbers[high]

		if current == target {
			return []int{low + 1, high + 1}
		} else if current > target {
			high -= 1
		} else {
			low += 1
		}
	}

	return []int{-1, -1}
}
