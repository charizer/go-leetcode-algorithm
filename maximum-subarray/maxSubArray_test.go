package maximum_subarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	want := 6
	got := MaxSubArray(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
