package best_time_to_buy_and_sell_stock

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MaxProfit(prices []int) int {
	ans := 0
	if len(prices) <= 1 {
		return ans
	}
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			ans = max(prices[i]-min, ans)
		}
	}
	return ans
}

func MaxProfit2(prices []int) int {
	ans := 0
	if len(prices) == 0 {
		return ans
	}
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		ans = max(ans, prices[i]-buy)
		if prices[i] < buy {
			buy = prices[i]
		}
	}
	return ans
}
