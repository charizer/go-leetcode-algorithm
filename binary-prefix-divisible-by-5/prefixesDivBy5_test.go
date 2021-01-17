package binary_prefix_divisible_by_5

import "testing"

func TestPrefixesDivBy5(t *testing.T) {
	nums := []int{0,1,1}
	want := []bool{true, false, false}
	got := PrefixesDivBy5(nums)
	for idx := range want {
		if got[idx] != want[idx] {
			t.Errorf("want:%t got:%t", want[idx],got[idx])
		}
	}
}
