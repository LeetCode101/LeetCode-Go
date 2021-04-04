package algorithms

func myPow2(x float64, n int) float64 {
	if n < 0 {
		return pow2(1/x, -n)
	}

	return pow2(x, n)
}

func pow2(x float64, n int) float64 {
	pow := 1.0

	for n >= 0 {
		if n&1 == 1 {
			pow *= x
			n -= 1
		}

		x *= x
		n /= 2
	}

	return pow
}
