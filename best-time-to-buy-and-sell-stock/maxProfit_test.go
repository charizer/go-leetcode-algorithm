package best_time_to_buy_and_sell_stock

import "testing"

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	want := 5
	got := MaxProfit(prices)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	prices = []int{7, 6, 4, 3, 1}
	want = 0
	got = MaxProfit(prices)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}

func TestMaxProfit2(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	want := 5
	got := MaxProfit2(prices)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	prices = []int{7, 6, 4, 3, 1}
	want = 0
	got = MaxProfit2(prices)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
