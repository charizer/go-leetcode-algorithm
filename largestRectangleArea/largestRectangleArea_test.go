package largestRectangleArea

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	arr := []int{2,1,5,6,2,3}
	want := 10
	got := LargestRectangleArea(arr)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	arr = []int{1,1}
	want = 2
	got = LargestRectangleArea(arr)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}

func TestLargestRectangleArea2(t *testing.T) {
	arr := []int{2,1,5,6,2,3}
	want := 10
	got := LargestRectangleArea2(arr)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	arr = []int{1,1}
	want = 2
	got = LargestRectangleArea2(arr)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}