package sliding_window_maximum

import "testing"

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1,3,-1,-3,5,3,6,7}
	k := 3
	want := []int{3,3,5,5,6,7}
	got := MaxSlidingWindow(nums, k)
	for idx, _ := range want {
		if got[idx] != want[idx]{
			t.Errorf("want:%d got:%d", want[idx],got[idx])
			break
		}
	}
	nums = []int{1}
	k = 1
	want = []int{1}
	got = MaxSlidingWindow(nums, k)
	for idx, _ := range want {
		if got[idx] != want[idx]{
			t.Errorf("want:%d got:%d", want[idx],got[idx])
			break
		}
	}
	nums = []int{1,-1}
	k = 1
	want = []int{1,-1}
	got = MaxSlidingWindow(nums, k)
	for idx, _ := range want {
		if got[idx] != want[idx]{
			t.Errorf("want:%d got:%d", want[idx],got[idx])
			break
		}
	}
	nums = []int{9,11}
	k = 2
	want = []int{11}
	got = MaxSlidingWindow(nums, k)
	for idx, _ := range want {
		if got[idx] != want[idx]{
			t.Errorf("want:%d got:%d", want[idx],got[idx])
			break
		}
	}
	nums = []int{4,-2}
	k = 2
	want = []int{4}
	got = MaxSlidingWindow(nums, k)
	for idx, _ := range want {
		if got[idx] != want[idx]{
			t.Errorf("want:%d got:%d", want[idx],got[idx])
			break
		}
	}
}
