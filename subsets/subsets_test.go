package subsets

import "testing"

func TestSubsets(t *testing.T) {
	nums := []int{1,2,3}
	got := Subsets(nums)
	t.Logf("got:%v", got)
}

func TestSubsets2(t *testing.T) {
	nums := []int{1,2,3}
	got := Subsets2(nums)
	t.Logf("got:%v", got)
}

func TestSubsets3(t *testing.T) {
	nums := []int{1,2,3}
	got := Subsets3(nums)
	t.Logf("got:%v", got)
}
