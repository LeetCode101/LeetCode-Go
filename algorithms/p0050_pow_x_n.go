package algorithms

func myPow(x float64, n int) float64 {
	if n < 0 {
		return pow(1/x, -n)
	}

	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n&1 == 1 {
		return x * pow(x*x, (n-1)/2)
	}

	return pow(x*x, n/2)
}
