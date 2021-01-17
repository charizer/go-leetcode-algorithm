package singleNumber

import "testing"

func TestSingleNumber(t *testing.T) {
	arr := []int{2,2,1}
	want := 1
	got := SingleNumber(arr)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	arr = []int{4,1,2,1,2}
	want = 4
	got = SingleNumber(arr)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}