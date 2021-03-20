package algorithms

func mySqrt(x int) int {
	low := 0
	high := x

	for low <= high {
		middle := low + (high - low) / 2
		value := middle * middle

		if value > x {
			high = middle - 1
		} else if value < x {
			low  = middle + 1
		} else {
			return middle
		}
	}

	return high
}