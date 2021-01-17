package intersection_of_two_arrays_ii

import "testing"

func TestIntersect(t *testing.T) {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	want := []int{2,2}
	got := Intersect(nums1, nums2)
	for idx, g := range got {
		if g != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], g)
		}
	}
	nums1 = []int{4,9,5}
	nums2 = []int{9,4,9,8,4}
	want = []int{9,4}
	got = Intersect(nums1, nums2)
	for idx, g := range got {
		if g != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], g)
		}
	}
}
