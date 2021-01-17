package best_time_to_buy_and_sell_stock_ii

func MaxProfit(prices []int) int {
	ans := 0
	if len(prices) == 0 {
		return ans
	}
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > buy {
			ans += prices[i] - buy
		}
		buy = prices[i]
	}
	return ans
}
