package algorithms

func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)

	for i, number := range nums {
		if index, ok := mapping[target-number]; ok {
			return []int{index, i}
		} else {
			mapping[number] = i
		}
	}

	return []int{-1, -1}
}
