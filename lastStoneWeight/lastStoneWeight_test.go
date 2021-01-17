package lastStoneWeight

import "testing"

func TestLastStoneWeight(t *testing.T) {
	arr := []int{2,7,4,1,8,1}
	want := 1
	got := LastStoneWeight(arr)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}