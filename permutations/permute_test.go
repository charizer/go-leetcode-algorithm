package permutations

import "testing"

func TestPermute(t *testing.T) {
	nums := []int{1,2,3}
	want := [][]int{
		{1,2,3},
		{1,3,2},
		{2,1,3},
		{2,3,1},
		{3,1,2},
		{3,2,1},
	}
	got := Permute(nums)
	for idx  := range got {
		for i := range got[idx] {
			if want[idx][i] != got[idx][i] {
				t.Errorf("want:%d got:%d", want[idx][i], got[idx][i])
			}
		}
	}
}

func TestPermute2(t *testing.T) {
	nums := []int{1,2,3}
	want := [][]int{
		{1,2,3},
		{1,3,2},
		{2,1,3},
		{2,3,1},
		{3,1,2},
		{3,2,1},
	}
	got := Permute2(nums)
	for idx  := range got {
		for i := range got[idx] {
			if want[idx][i] != got[idx][i] {
				t.Errorf("want:%d got:%d", want[idx][i], got[idx][i])
			}
		}
	}
}

func TestPermute3(t *testing.T) {
	nums := []int{1,2,3}
	want := [][]int{
		{1,2,3},
		{1,3,2},
		{2,1,3},
		{2,3,1},
		{3,1,2},
		{3,2,1},
	}
	got := Permute3(nums)
	for idx  := range got {
		for i := range got[idx] {
			if want[idx][i] != got[idx][i] {
				t.Errorf("want:%d got:%d", want[idx][i], got[idx][i])
			}
		}
	}
}
