package different_ways_to_add_parentheses

import "testing"

func TestDiffWaysToCompute(t *testing.T) {
	s := "2-1-1"
	want := []int{0,2}
	got := DiffWaysToCompute(s)
	for idx := range want {
		if got[idx] != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], got[idx])
		}
	}
}
