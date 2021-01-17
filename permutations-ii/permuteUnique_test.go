package permutations_ii

import "testing"

func TestPermuteUnique(t *testing.T) {
	nums := []int{1,1,2}
	want := [][]int{
		{1,1,2},
		{1,2,1},
		{2,1,1},
	}
	got := PermuteUnique(nums)
	for idx  := range got {
		for i := range got[idx] {
			if want[idx][i] != got[idx][i] {
				t.Errorf("want:%d got:%d", want[idx][i], got[idx][i])
			}
		}
	}
}
