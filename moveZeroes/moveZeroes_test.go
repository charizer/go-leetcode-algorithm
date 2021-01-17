package moveZeroes

import "testing"

func TestMoveZeroes(t *testing.T) {
	arr := []int{0,1,0,3,12}
	want := []int{1,3,12,0,0}
	MoveZeroes(arr)
	for idx, item := range arr {
		if item != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], item)
		}
	}
	arr = []int{0,1,0,3,12}
	want = []int{1,3,12,0,0}
	MoveZeroes2(arr)
	for idx, item := range arr {
		if item != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], item)
		}
	}
}
