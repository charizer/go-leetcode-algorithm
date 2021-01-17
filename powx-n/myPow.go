package powx_n

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	fast := fastPow(x, n/2)
	if n % 2 == 1 {
		return fast * fast * x
	}else{
		return fast * fast
	}
}

func MyPow(x float64, n int) float64 {
	if n < 0 {
		x = 1/x
		n = -n
	}
	return fastPow(x, n)
}

