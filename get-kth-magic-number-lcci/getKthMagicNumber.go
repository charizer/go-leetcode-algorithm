package get_kth_magic_number_lcci

func min(a, b, c int) int {
	if a > b {
		if b > c {
			return c
		} else {
			return b
		}
	} else {
		if a > c {
			return c
		} else {
			return a
		}
	}
}

func GetKthMagicNumber(k int) int {
	if k <= 1 {
		return 1
	}
	dp := make([]int, k)
	dp[0] = 1
	key3, key5, key7 := 0, 0, 0
	for i := 1; i < k; i++ {
		dp[i] = min(dp[key3]*3, dp[key5]*5, dp[key7]*7)
		if dp[i] == dp[key3]*3 {
			key3++
		}
		if dp[i] == dp[key5]*5 {
			key5++
		}
		if dp[i] == dp[key7]*7 {
			key7++
		}
	}
	return dp[k-1]
}
