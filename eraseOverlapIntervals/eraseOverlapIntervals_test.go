package eraseOverlapIntervals

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	arr := [][]int{{1,2}, {2,3}, {3,4}, {1,3} }
	want := 1
	got := EraseOverlapIntervals(arr)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	arr = [][]int{{1,2},{1,2},{1,2}}
	want = 2
	got = EraseOverlapIntervals(arr)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	arr = [][]int{{1,2},{2,3}}
	want = 0
	got = EraseOverlapIntervals(arr)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}