package algorithms

import "math"

func sortedSquares(nums []int) []int {
	var squares = make([]int, len(nums))
	low := 0
	high := len(nums) - 1
	end := len(nums) - 1

	for low <= high {
		left := math.Abs(float64(nums[low]))
		right := math.Abs(float64(nums[high]))

		if left > right {
			squares[end] = int(left * left)
			low += 1
		} else {
			squares[end] = int(right * right)
			high -= 1
		}

		end -= 1
	}

	return squares
}
