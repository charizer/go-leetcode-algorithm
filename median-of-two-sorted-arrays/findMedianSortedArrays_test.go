package median_of_two_sorted_arrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1,3}
	nums2 := []int{2}
	want := 2.0
	got := FindMedianSortedArrays(nums1, nums2)
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	want = 2.5
	got = FindMedianSortedArrays(nums1, nums2)
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
}

func TestFindMedianSortedArrays2(t *testing.T) {
	nums1 := []int{1,3}
	nums2 := []int{2}
	want := 2.0
	got := FindMedianSortedArrays2(nums1, nums2)
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	want = 2.5
	got = FindMedianSortedArrays2(nums1, nums2)
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
}
