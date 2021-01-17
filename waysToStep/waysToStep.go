package waysToStep

func WaysToStep(n int) int {
	if n < 2 {
		return 1
	}
	if n == 2 {
		return 2
	}
	f0, f1, f2, f3 := 1, 1, 2, 4
	for i := 3; i <= n; i++ {
		f3 = (f0 + f1 + f2) % 1000000007
		f0 = f1
		f1 = f2
		f2 = f3
	}
	return f3
}

func WaysToStep2(n int) int {
	if n < 2 {
		return 1
	}
	if n == 2 {
		return 2
	}
	f0, f1, f2, f3 := 1, 1, 2, 4
	for i := 3; i <= n; i++ {
		f3 = (f0 + f1 + f2) % 1000000007
		f0 = f1
		f1 = f2
		f2 = f3
	}
	return f3
}
