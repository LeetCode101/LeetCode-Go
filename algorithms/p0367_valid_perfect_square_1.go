package algorithms

func isPerfectSquare1(num int) bool {
	low := 0
	high := num

	for low <= high {
		middle := low + (high-low)/2
		square := middle * middle

		if square == num {
			return true
		} else if square > num {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return false
}
