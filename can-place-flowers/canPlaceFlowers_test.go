package can_place_flowers

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	arr := []int{1,0,0,0,1}
	want := true
	n := 1
	got := CanPlaceFlowers(arr, n)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	arr = []int{0,0,0,1,0}
	want = false
	n = 2
	got = CanPlaceFlowers(arr, n)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
}
