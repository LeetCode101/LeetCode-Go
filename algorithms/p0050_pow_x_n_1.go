package algorithms

func myPow1(x float64, n int) float64 {
	if n < 0 {
		return pow1(1/x, -n)
	}

	return pow1(x, n)
}

func pow1(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n&1 == 1 {
		return x * pow1(x*x, (n-1)/2)
	}

	return pow1(x*x, n/2)
}
