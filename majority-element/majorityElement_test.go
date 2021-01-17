package majority_element

import "testing"

func TestMajorityElement(t *testing.T) {
	nums := []int{3,2,3}
	want := 3
	got := MajorityElement(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	nums = []int{2,2,1,1,1,2,2}
	want = 2
	got = MajorityElement(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}

func TestMajorityElement2(t *testing.T) {
	nums := []int{3,2,3}
	want := 3
	got := MajorityElement2(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	nums = []int{2,2,1,1,1,2,2}
	want = 2
	got = MajorityElement2(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
