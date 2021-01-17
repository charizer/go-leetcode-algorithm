package summary_ranges

import "testing"

func TestSummaryRanges(t *testing.T) {
	nums := []int{0,1,2,4,5,7}
	want := []string{"0->2","4->5","7"}
	got := SummaryRanges(nums)
	for idx := range got {
		if got[idx] != want[idx] {
			t.Errorf("want:%s got:%s", want, got)
		}
	}
	nums = []int{0,2,3,4,6,8,9}
	want = []string{"0","2->4","6","8->9"}
	got = SummaryRanges(nums)
	for idx := range got {
		if got[idx] != want[idx] {
			t.Errorf("want:%s got:%s", want, got)
		}
	}
}
