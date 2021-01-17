package last_stone_weight

import "testing"

func TestLastStoneWeight(t *testing.T) {
	arr := []int{2,7,4,1,8,1}
	want := 1
	got := LastStoneWeight(arr)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}

func TestLastStoneWeight2(t *testing.T) {
	arr := []int{2,7,4,1,8,1}
	want := 1
	got := LastStoneWeight2(arr)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}

func TestLastStoneWeight3(t *testing.T) {
	arr := []int{2,7,4,1,8,1}
	want := 1
	got := LastStoneWeight3(arr)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	arr = []int{1, 3}
	want = 2
	got = LastStoneWeight3(arr)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	arr = []int{3, 7, 8}
	want = 2
	got = LastStoneWeight3(arr)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
